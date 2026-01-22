package role

import "time"

type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;column:id_role"`
	Name      string    `json:"name" gorm:"type:varchar(50);not null"`
	Desc      string    `json:"description" gorm:"type:text"`
	Status    bool      `json:"status" gorm:"type:boolean;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at;type:timestamptz;autoUpdateTime"`
}

func (Role) TableName() string { return "roles" }
