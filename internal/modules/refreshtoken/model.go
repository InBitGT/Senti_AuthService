package refreshtoken

import "time"

type RefreshToken struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_refresh_token"`
	UserID    uint       `json:"user_id" gorm:"not null"`
	Token     string     `json:"token" gorm:"type:text;not null"`
	ExpiresAt time.Time  `json:"expires_at"`
	Revoked   bool       `json:"revoked" gorm:"default:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
}

func (RefreshToken) TableName() string { return "refresh_tokens" }
