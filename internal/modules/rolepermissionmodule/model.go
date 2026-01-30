package rolepermissionmodule

import "time"

type RolePermissionModule struct {
	ID                 uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_role_permission_module"`
	RoleID             uint       `json:"role_id" gorm:"not null;column:role_id"`
	PermissionModuleID uint       `json:"permission_module_id" gorm:"not null;column:permission_module_id"`
	Status             bool       `json:"status" gorm:"default:true"`
	CreatedAt          *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (RolePermissionModule) TableName() string {
	return "role_permission_module"
}
