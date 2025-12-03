package otp

import "time"

type UserOTP struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_otp"`
	UserID    uint       `json:"user_id" gorm:"not null"`
	TenantID  uint       `json:"tenant_id" gorm:"not null"`
	Code      string     `json:"code" gorm:"type:varchar(10);not null"`
	Channel   string     `json:"channel" gorm:"type:varchar(20)"`
	ExpiresAt time.Time  `json:"expires_at"`
	Used      bool       `json:"used" gorm:"default:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
}

func (UserOTP) TableName() string { return "user_otp" }
