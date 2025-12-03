package user

import "time"

type User struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_user"`
	TenantID     uint       `json:"tenant_id" gorm:"not null"`
	Email        string     `json:"email" gorm:"type:varchar(150);not null"`
	PasswordHash string     `json:"-" gorm:"type:varchar(255);not null"`
	RoleID       uint       `json:"role_id" gorm:"not null"`
	IsActive     bool       `json:"is_active" gorm:"default:true"`
	TwoFAEnabled bool       `json:"two_fa_enabled" gorm:"default:false"`
	CreatedAt    *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (User) TableName() string { return "users" }
