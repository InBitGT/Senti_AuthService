package auth

type Permission struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement;column:id_permission"`
	Key  string `json:"key" gorm:"type:varchar(100);unique;not null;column:key"`
	Desc string `json:"description" gorm:"type:text;column:description"`
}

func (Permission) TableName() string { return "permissions" }

type RolePermission struct {
	RoleID       uint `json:"role_id" gorm:"primaryKey;column:id_role"`
	PermissionID uint `json:"permission_id" gorm:"primaryKey;column:id_permission"`
}

func (RolePermission) TableName() string { return "role_permissions" }
