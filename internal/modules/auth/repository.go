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
	GetActiveTenantByCode(code string) (*tenant.Tenant, error)
	GetRoleByName(tenantID uint, name string) (*role.Role, error)

	CreateUser(u *user.User) error
	UpdateUser(u *user.User) error
	GetUserByID(id uint) (*user.User, error)
	GetUserByEmail(tenantID uint, email string) (*user.User, error)

	GetPermissionsByRole(roleID uint) ([]permission.Permission, error)

	SaveRefreshToken(token *refreshtoken.RefreshToken) error
	GetRefreshToken(token string) (*refreshtoken.RefreshToken, error)
	RevokeRefreshToken(id uint) error

	SaveOTP(o *otp.UserOTP) error
	FindValidOTP(userID, tenantID uint, code string, now time.Time) (*otp.UserOTP, error)
	MarkOTPUsed(id uint) error

	CreateTenant(t *tenant.Tenant) error
	HardDeleteTenant(id uint) error

	CreateRole(r *role.Role) error
}

type authRepository struct {
	db *gorm.DB
}

type ModulePermissionRow struct {
	ModuleID      uint   `gorm:"column:module_id"`
	PermissionKey string `gorm:"column:permission_key"`
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

// ✅ antes: code + is_active=true
// ✅ ahora: code + status=true (nuevo schema)
func (r *authRepository) GetActiveTenantByCode(code string) (*tenant.Tenant, error) {
	var t tenant.Tenant
	err := r.db.Where("code = ? AND status = true", code).First(&t).Error
	return &t, err
}

func (r *authRepository) GetRoleByName(tenantID uint, name string) (*role.Role, error) {
	var rl role.Role
	err := r.db.Where("tenant_id = ? AND name = ?", tenantID, name).First(&rl).Error
	return &rl, err
}

func (r *authRepository) CreateUser(u *user.User) error {
	return r.db.Create(u).Error
}

func (r *authRepository) UpdateUser(u *user.User) error {
	return r.db.Save(u).Error
}

func (r *authRepository) GetUserByID(id uint) (*user.User, error) {
	var u user.User
	err := r.db.First(&u, id).Error
	return &u, err
}

func (r *authRepository) GetUserByEmail(tenantID uint, email string) (*user.User, error) {
	var u user.User
	// NOTA: aquí todavía filtra is_active=true (tu user model aún lo tiene).
	// Cuando alinees user al nuevo schema (status bool), lo cambiamos.
	err := r.db.Where("tenant_id = ? AND email = ? AND is_active = true", tenantID, email).First(&u).Error
	return &u, err
}

func (r *authRepository) GetPermissionsByRole(roleID uint) ([]permission.Permission, error) {
	var perms []permission.Permission

	err := r.db.Table("permissions p").
		Joins("JOIN role_permissions rp ON rp.id_permission = p.id_permission").
		Where("rp.id_role = ?", roleID).
		Find(&perms).Error

	return perms, err
}

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

func (r *authRepository) CreateTenant(t *tenant.Tenant) error {
	return r.db.Create(t).Error
}

func (r *authRepository) HardDeleteTenant(id uint) error {
	return r.db.Unscoped().Where("id_tenant = ?", id).Delete(&tenant.Tenant{}).Error
}

func (r *authRepository) CreateRole(rle *role.Role) error {
	return r.db.Create(rle).Error
}
