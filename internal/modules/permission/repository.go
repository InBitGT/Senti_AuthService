package permission

import "gorm.io/gorm"

type PermissionRepository interface {
	Create(p *Permission) error
	Update(id uint, p *Permission) (*Permission, error)
	Delete(id uint) error     // soft delete
	HardDelete(id uint) error // físico
	GetByID(id uint) (*Permission, error)
	GetAll() ([]Permission, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db}
}

func (r *permissionRepository) Create(p *Permission) error {
	return r.db.Create(p).Error
}

func (r *permissionRepository) Update(id uint, p *Permission) (*Permission, error) {
	var existing Permission
	if err := r.db.Where("id_permission = ? AND status = true", id).First(&existing).Error; err != nil {
		return nil, err
	}

	updates := map[string]interface{}{
		"key":         p.Key,
		"description": p.Desc,
	}

	if err := r.db.Model(&Permission{}).
		Where("id_permission = ?", id).
		Updates(updates).Error; err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

func (r *permissionRepository) Delete(id uint) error {
	res := r.db.Model(&Permission{}).
		Where("id_permission = ? AND status = true", id).
		Updates(map[string]interface{}{"status": false})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *permissionRepository) HardDelete(id uint) error {
	return r.db.Unscoped().
		Where("id_permission = ?", id).
		Delete(&Permission{}).Error
}

func (r *permissionRepository) GetByID(id uint) (*Permission, error) {
	var p Permission
	err := r.db.Where("id_permission = ? AND status = true", id).First(&p).Error
	return &p, err
}

func (r *permissionRepository) GetAll() ([]Permission, error) {
	var list []Permission
	err := r.db.Where("status = true").Find(&list).Error
	return list, err
}
