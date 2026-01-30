package rolepermissionmodule

import "gorm.io/gorm"

type Repository interface {
	Create(rp *RolePermissionModule) error
	SoftDelete(id uint) error
	GetByRole(roleID uint) ([]RolePermissionModule, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(rp *RolePermissionModule) error {
	return r.db.Create(rp).Error
}

func (r *repository) SoftDelete(id uint) error {
	return r.db.Model(&RolePermissionModule{}).
		Where("id_role_permission_module = ?", id).
		Update("status", false).Error
}

func (r *repository) GetByRole(roleID uint) ([]RolePermissionModule, error) {
	var list []RolePermissionModule
	err := r.db.
		Where("role_id = ? AND status = true", roleID).
		Find(&list).Error
	return list, err
}
