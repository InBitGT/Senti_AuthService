package role

import "time"

type Role struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_role"`
	TenantID  uint       `json:"tenant_id" gorm:"not null"`
	Name      string     `json:"name" gorm:"type:varchar(50);not null"`
	Desc      string     `json:"description" gorm:"type:text"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (Role) TableName() string { return "roles" }
