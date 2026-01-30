package permissionmodule

import "gorm.io/gorm"

type Repository interface {
	Create(pm *PermissionModule) error
	GetByID(id uint) (*PermissionModule, error)
	GetAll() ([]PermissionModule, error)
	Delete(id uint) error
	HardDelete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(pm *PermissionModule) error {
	return r.db.Create(pm).Error
}

func (r *repository) GetByID(id uint) (*PermissionModule, error) {
	var pm PermissionModule
	err := r.db.
		Where("id_permission_module = ? AND status = true", id).
		First(&pm).Error
	return &pm, err
}

func (r *repository) GetAll() ([]PermissionModule, error) {
	var list []PermissionModule
	err := r.db.Where("status = true").Find(&list).Error
	return list, err
}

func (r *repository) Delete(id uint) error {
	res := r.db.Model(&PermissionModule{}).
		Where("id_permission_module = ? AND status = true", id).
		Update("status", false)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *repository) HardDelete(id uint) error {
	return r.db.Unscoped().
		Where("id_permission_module = ?", id).
		Delete(&PermissionModule{}).Error
}
