package tenant

import "time"

type Tenant struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_tenant"`
	Code      string     `json:"code" gorm:"type:varchar(50);unique;not null"`
	Name      string     `json:"name" gorm:"type:varchar(150);not null"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (Tenant) TableName() string { return "tenants" }
