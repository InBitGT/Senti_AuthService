package auth

import (
	"time"

	"gorm.io/gorm"
)

type AuthRepository interface {
	GetActiveTenantByCode(code string) (*Tenant, error)
	GetRoleByName(tenantID uint, name string) (*Role, error)
	CreateUser(user *User) error
	GetUserByEmail(tenantID uint, email string) (*User, error)
	GetPermissionsByRole(roleID uint) ([]Permission, error)
	SaveRefreshToken(token *RefreshToken) error
	GetRefreshToken(token string) (*RefreshToken, error)
	RevokeRefreshToken(id uint) error
	SaveOTP(otp *UserOTP) error
	FindValidOTP(userID, tenantID uint, code string, now time.Time) (*UserOTP, error)
	MarkOTPUsed(id uint) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) GetActiveTenantByCode(code string) (*Tenant, error) {
	var t Tenant
	err := r.db.Where("code = ? AND is_active = ?", code, true).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *authRepository) GetRoleByName(tenantID uint, name string) (*Role, error) {
	var role Role
	err := r.db.Where("id_tenant = ? AND name = ?", tenantID, name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *authRepository) CreateUser(u *User) error {
	return r.db.Create(u).Error
}

func (r *authRepository) GetUserByEmail(tenantID uint, email string) (*User, error) {
	var u User
	err := r.db.Where("id_tenant = ? AND email = ? AND is_active = ?", tenantID, email, true).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *authRepository) GetPermissionsByRole(roleID uint) ([]Permission, error) {
	var perms []Permission
	err := r.db.
		Table("permissions p").
		Joins("INNER JOIN role_permissions rp ON rp.id_permission = p.id_permission").
		Where("rp.id_role = ?", roleID).
		Find(&perms).Error
	return perms, err
}

func (r *authRepository) SaveRefreshToken(t *RefreshToken) error {
	return r.db.Create(t).Error
}

func (r *authRepository) GetRefreshToken(token string) (*RefreshToken, error) {
	var rt RefreshToken
	err := r.db.Where("token = ? AND revoked = ?", token, false).First(&rt).Error
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *authRepository) RevokeRefreshToken(id uint) error {
	return r.db.Model(&RefreshToken{}).
		Where("id_refresh_token = ?", id).
		Update("revoked", true).Error
}

func (r *authRepository) SaveOTP(o *UserOTP) error {
	return r.db.Create(o).Error
}

func (r *authRepository) FindValidOTP(userID, tenantID uint, code string, now time.Time) (*UserOTP, error) {
	var otp UserOTP
	err := r.db.Where(
		"id_user = ? AND id_tenant = ? AND code = ? AND used = ? AND expires_at >= ?",
		userID, tenantID, code, false, now,
	).First(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *authRepository) MarkOTPUsed(id uint) error {
	return r.db.Model(&UserOTP{}).
		Where("id_otp = ?", id).
		Update("used", true).Error
}
