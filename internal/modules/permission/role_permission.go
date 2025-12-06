package permission

type RolePermission struct {
	ID           uint `json:"id" gorm:"primaryKey;autoIncrement;column:id_role_permission"`
	RoleID       uint `json:"role_id" gorm:"column:id_role;not null"`
	PermissionID uint `json:"permission_id" gorm:"column:id_permission;not null"`
}

func (RolePermission) TableName() string { return "role_permissions" }
