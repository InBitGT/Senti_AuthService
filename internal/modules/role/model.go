package role

import "time"

type Role struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_role"`
	Name      string     `json:"name" gorm:"type:varchar(50);not null"`
	Desc      string     `json:"description" gorm:"type:text"`
	Status    bool       `json:"status" gorm:"default:true"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"update_at,omitempty" gorm:"column:update_at;autoUpdateTime"`
}

func (Role) TableName() string { return "roles" }
