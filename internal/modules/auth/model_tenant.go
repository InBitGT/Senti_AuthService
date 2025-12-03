package auth

import "time"

type Tenant struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_tenant"`
	Code      string     `json:"code" gorm:"type:varchar(50);unique;not null;column:code"`
	Name      string     `json:"name" gorm:"type:varchar(150);not null;column:name"`
	IsActive  bool       `json:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;autoUpdateTime"`
}

func (Tenant) TableName() string { return "tenants" }
