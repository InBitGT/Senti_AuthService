package user

import "time"

type User struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_user"`
	TenantID     uint       `json:"tenant_id" gorm:"not null;column:tenant_id"`
	Email        string     `json:"email" gorm:"type:varchar(150);not null"`
	PasswordHash string     `json:"-" gorm:"type:varchar(255);not null"`
	Phone        string     `json:"phone" gorm:"type:varchar(20)"`
	FirstName    string     `json:"first_name" gorm:"type:varchar(100)"`
	LastName     string     `json:"last_name" gorm:"type:varchar(100)"`
	AddressID    uint       `json:"address_id" gorm:"column:id_address"`
	RoleID       uint       `json:"role_id" gorm:"not null;column:role_id"`
	IsActive     bool       `json:"is_active" gorm:"default:true"`
	TwoFAEnabled bool       `json:"two_fa_enabled" gorm:"default:false"`
	CreatedAt    *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (User) TableName() string { return "users" }
