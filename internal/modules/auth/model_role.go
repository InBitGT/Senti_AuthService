package auth

import "time"

type Role struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_role"`
	TenantID  uint       `json:"tenant_id" gorm:"not null;column:id_tenant;index"`
	Name      string     `json:"name" gorm:"type:varchar(50);not null;column:name"`
	Desc      string     `json:"description" gorm:"type:text;column:description"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;autoUpdateTime"`
}

func (Role) TableName() string { return "roles" }
