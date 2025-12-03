package auth

import (
	"time"
)

type User struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_user"`
	TenantID     uint       `json:"tenant_id" gorm:"not null;column:id_tenant;index"`
	Email        string     `json:"email" gorm:"type:varchar(150);not null;column:email"`
	PasswordHash string     `json:"-" gorm:"type:varchar(255);not null;column:password_hash"`
	RoleID       uint       `json:"role_id" gorm:"not null;column:id_role;index"`
	IsActive     bool       `json:"is_active" gorm:"column:is_active;default:true"`
	MustChangePw bool       `json:"must_change_pw" gorm:"column:must_change_pw;default:false"`
	TwoFAEnabled bool       `json:"two_fa_enabled" gorm:"column:two_fa_enabled;default:false"`
	CreatedAt    *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string { return "users" }
