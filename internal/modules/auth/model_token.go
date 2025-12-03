package auth

import "time"

type RefreshToken struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_refresh_token"`
	UserID    uint       `json:"user_id" gorm:"not null;column:id_user;index"`
	Token     string     `json:"token" gorm:"type:text;not null;column:token"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"column:expires_at;not null"`
	Revoked   bool       `json:"revoked" gorm:"column:revoked;default:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
}

func (RefreshToken) TableName() string { return "refresh_tokens" }
