package permissionmodule

import "time"

type PermissionModule struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement;column:id_permission_module"`
	PermissionID uint      `json:"permission_id" gorm:"not null;column:permission_id"`
	ModuleID     uint      `json:"module_id" gorm:"not null;column:module_id"`
	Status       bool      `json:"status" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt    time.Time `json:"update_at" gorm:"column:update_at;type:timestamptz;autoUpdateTime"`
}

func (PermissionModule) TableName() string {
	return "permission_module"
}
