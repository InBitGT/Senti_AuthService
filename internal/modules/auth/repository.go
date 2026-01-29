package auth

import (
	"time"

	"AuthService/internal/modules/otp"
	"AuthService/internal/modules/permission"
	"AuthService/internal/modules/refreshtoken"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/tenant"
	"AuthService/internal/modules/user"

	"gorm.io/gorm"
)

type AuthRepository interface {
	// TENANT
	GetActiveTenantByCode(code string) (*tenant.Tenant, error)
	CreateTenant(t *tenant.Tenant) error
	HardDeleteTenant(id uint) error
	UpdateTenantSuscription(tenantID uint, suscriptionID uint) error

	// ROLES (globales)
	GetRoleByName(name string) (*role.Role, error)
	CreateRole(r *role.Role) error

	// USER (auth db)
	CreateUser(u *user.User) error
	UpdateUser(u *user.User) error
	GetUserByID(id uint) (*user.User, error)
	GetUserByEmail(tenantID uint, email string) (*user.User, error)

	// PERMS legacy (role_permissions viejo)
	GetPermissionsByRole(roleID uint) ([]permission.Permission, error)

	// RBAC nuevo (por módulo)
	GetModulePermissionsByRole(roleID uint) (map[uint][]string, error)

	// REFRESH TOKENS
	SaveRefreshToken(token *refreshtoken.RefreshToken) error
	GetRefreshToken(token string) (*refreshtoken.RefreshToken, error)
	RevokeRefreshToken(id uint) error

	// OTP
	SaveOTP(o *otp.UserOTP) error
	FindValidOTP(userID, tenantID uint, code string, now time.Time) (*otp.UserOTP, error)
	MarkOTPUsed(id uint) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

// ------------------------------ TENANT ------------------------------

func (r *authRepository) GetActiveTenantByCode(code string) (*tenant.Tenant, error) {
	var t tenant.Tenant
	err := r.db.Where("code = ? AND status = true", code).First(&t).Error
	return &t, err
}

func (r *authRepository) CreateTenant(t *tenant.Tenant) error {
	return r.db.Create(t).Error
}

func (r *authRepository) HardDeleteTenant(id uint) error {
	return r.db.Unscoped().Where("id_tenant = ?", id).Delete(&tenant.Tenant{}).Error
}

func (r *authRepository) UpdateTenantSuscription(tenantID uint, suscriptionID uint) error {
	return r.db.Model(&tenant.Tenant{}).
		Where("id_tenant = ?", tenantID).
		Update("suscription_id", suscriptionID).Error
}

// ------------------------------ ROLES ------------------------------

func (r *authRepository) GetRoleByName(name string) (*role.Role, error) {
	var rl role.Role
	err := r.db.Where("name = ? AND status = true", name).First(&rl).Error
	return &rl, err
}

func (r *authRepository) CreateRole(rle *role.Role) error {
	return r.db.Create(rle).Error
}

// ------------------------------ USERS (auth db) ------------------------------

// Cuando migremos definitivamente a UserService y saquemos user del AuthService, esto se ajusta.
func (r *authRepository) CreateUser(u *user.User) error { return r.db.Create(u).Error }
func (r *authRepository) UpdateUser(u *user.User) error { return r.db.Save(u).Error }

func (r *authRepository) GetUserByID(id uint) (*user.User, error) {
	var u user.User
	err := r.db.First(&u, id).Error
	return &u, err
}

// Cuando alinees user a status bool, se cambia a "status = true".
func (r *authRepository) GetUserByEmail(tenantID uint, email string) (*user.User, error) {
	var u user.User
	err := r.db.Where("tenant_id = ? AND email = ? AND is_active = true", tenantID, email).First(&u).Error
	return &u, err
}

// ------------------------------ PERMISSIONS (legacy role_permissions) -----------------------------

func (r *authRepository) GetPermissionsByRole(roleID uint) ([]permission.Permission, error) {
	var perms []permission.Permission

	err := r.db.Table("permissions p").
		Joins("JOIN role_permissions rp ON rp.id_permission = p.id_permission").
		Where("rp.id_role = ?", roleID).
		Find(&perms).Error

	return perms, err
}

// ------------------------------ RBAC NUEVO (permission_module / role_permission_module) ------------------------------

// Retorna: map[module_id][]permission_key
func (r *authRepository) GetModulePermissionsByRole(roleID uint) (map[uint][]string, error) {
	type row struct {
		ModuleID uint   `gorm:"column:module_id"`
		Key      string `gorm:"column:key"`
	}

	var rows []row

	err := r.db.Raw(`
		SELECT pm.module_id AS module_id, p.key AS key
		FROM role_permission_module rpm
		INNER JOIN permission_module pm ON pm.id_permission_module = rpm.permission_module_id
		INNER JOIN permissions p ON p.id_permission = pm.permission_id
		WHERE rpm.role_id = ?
		  AND rpm.status = true
		  AND pm.status = true
		  AND p.status = true
	`, roleID).Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	out := map[uint][]string{}
	for _, it := range rows {
		out[it.ModuleID] = append(out[it.ModuleID], it.Key)
	}
	return out, nil
}

// ------------------------------ REFRESH TOKENS ------------------------------

func (r *authRepository) SaveRefreshToken(rt *refreshtoken.RefreshToken) error {
	return r.db.Create(rt).Error
}

func (r *authRepository) GetRefreshToken(token string) (*refreshtoken.RefreshToken, error) {
	var rt refreshtoken.RefreshToken
	err := r.db.Where("token = ? AND revoked = false", token).First(&rt).Error
	return &rt, err
}

func (r *authRepository) RevokeRefreshToken(id uint) error {
	return r.db.Model(&refreshtoken.RefreshToken{}).
		Where("id_refresh_token = ?", id).
		Update("revoked", true).Error
}

// ------------------------------ OTP ------------------------------

func (r *authRepository) SaveOTP(o *otp.UserOTP) error {
	return r.db.Create(o).Error
}

func (r *authRepository) FindValidOTP(userID, tenantID uint, code string, now time.Time) (*otp.UserOTP, error) {
	var out otp.UserOTP
	err := r.db.Where(`
		user_id = ? AND tenant_id = ? AND code = ? AND used = false AND expires_at >= ?`,
		userID, tenantID, code, now,
	).First(&out).Error
	return &out, err
}

func (r *authRepository) MarkOTPUsed(id uint) error {
	return r.db.Model(&otp.UserOTP{}).
		Where("id_otp = ?", id).
		Update("used", true).Error
}
