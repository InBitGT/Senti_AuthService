package auth

import "time"

type UserOTP struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_otp"`
	UserID    uint       `json:"user_id" gorm:"not null;column:id_user;index"`
	TenantID  uint       `json:"tenant_id" gorm:"not null;column:id_tenant;index"`
	Code      string     `json:"code" gorm:"type:varchar(10);not null;column:code"`
	Channel   string     `json:"channel" gorm:"type:varchar(20);not null;column:channel"` // sms,email,app
	ExpiresAt time.Time  `json:"expires_at" gorm:"column:expires_at;not null"`
	Used      bool       `json:"used" gorm:"column:used;default:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
}

func (UserOTP) TableName() string { return "user_otp" }
