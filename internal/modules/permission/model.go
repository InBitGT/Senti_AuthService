package permission

import "time"

type Permission struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;column:id_permission"`
	Key       string    `json:"key" gorm:"type:varchar(100);unique;not null"`
	Desc      string    `json:"description" gorm:"type:text"`
	Status    bool      `json:"status" gorm:"type:boolean;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at;type:timestamptz;autoUpdateTime"`
}

func (Permission) TableName() string { return "permissions" }
