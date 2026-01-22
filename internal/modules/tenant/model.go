package tenant

import "time"

type Tenant struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement;column:id_tenant"`
	Code          string    `json:"code" gorm:"type:varchar(50);unique;not null"`
	Name          string    `json:"name" gorm:"type:varchar(150);not null"`
	Picture       string    `json:"picture" gorm:"type:text"`
	NIT           string    `json:"nit" gorm:"type:varchar(50)"`
	Phone         string    `json:"phone" gorm:"type:varchar(30)"`
	Email         string    `json:"email" gorm:"type:varchar(150)"`
	SuscriptionID uint      `json:"suscription_id" gorm:"column:suscription_id"`
	AddressID     uint      `json:"address_id" gorm:"column:address_id"`
	Status        bool      `json:"status" gorm:"type:boolean;default:true"`
	CreatedAt     time.Time `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt     time.Time `json:"update_at" gorm:"column:update_at;type:timestamptz;autoUpdateTime"`
}

func (Tenant) TableName() string { return "tenants" }
