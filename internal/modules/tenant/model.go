package tenant

import "time"

type Tenant struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_tenant"`
	Code      string     `json:"code" gorm:"type:varchar(50);unique;not null"`
	Name      string     `json:"name" gorm:"type:varchar(150);not null"`
	NIT       string     `json:"nit" gorm:"type:varchar(50)"`
	Phone     string     `json:"phone" gorm:"type:varchar(30)"`
	Email     string     `json:"email" gorm:"type:varchar(150)"`
	AddressID uint       `json:"address_id" gorm:"not null;column:id_address"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (Tenant) TableName() string { return "tenants" }
