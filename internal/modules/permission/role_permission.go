package permission

type RolePermission struct {
	RoleID       uint `json:"role_id" gorm:"primaryKey;column:id_role"`
	PermissionID uint `json:"permission_id" gorm:"primaryKey;column:id_permission"`
}

func (RolePermission) TableName() string { return "role_permissions" }
