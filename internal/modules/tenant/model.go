package tenant

import "time"

type Tenant struct {
	ID            uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_tenant"`
	Code          string     `json:"code" gorm:"type:varchar(50);unique;not null"`
	Name          string     `json:"name" gorm:"type:varchar(150);not null"`
	Picture       string     `json:"picture,omitempty" gorm:"type:text"`
	NIT           string     `json:"nit,omitempty" gorm:"type:varchar(50)"`
	Phone         string     `json:"phone,omitempty" gorm:"type:varchar(30)"`
	Email         string     `json:"email,omitempty" gorm:"type:varchar(150)"`
	SuscriptionID *uint      `json:"suscription_id,omitempty" gorm:"column:suscription_id"`
	AddressID     uint       `json:"address_id" gorm:"not null;column:address_id"`
	Status        bool       `json:"status" gorm:"default:true"`
	CreatedAt     *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     *time.Time `json:"update_at,omitempty" gorm:"column:update_at;autoUpdateTime"`
}

func (Tenant) TableName() string { return "tenants" }
