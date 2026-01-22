package role

import "gorm.io/gorm"

type RoleRepository interface {
	Create(role *Role) error
	Update(id uint, role *Role) (*Role, error)
	Delete(id uint) error     // soft delete => status=false
	HardDelete(id uint) error // físico
	GetByID(id uint) (*Role, error)
	GetAll() ([]Role, error)
}

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepo{db}
}

func (r *roleRepo) Create(role *Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepo) Update(id uint, role *Role) (*Role, error) {
	var existing Role
	if err := r.db.Where("id_role = ? AND status = true", id).First(&existing).Error; err != nil {
		return nil, err
	}

	updates := map[string]interface{}{
		"name":        role.Name,
		"description": role.Desc,
		"status":      role.Status, // si querés permitir re-activar desde update
	}

	if err := r.db.Model(&Role{}).Where("id_role = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

func (r *roleRepo) Delete(id uint) error {
	res := r.db.Model(&Role{}).
		Where("id_role = ? AND status = true", id).
		Updates(map[string]interface{}{"status": false})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *roleRepo) HardDelete(id uint) error {
	return r.db.Unscoped().Where("id_role = ?", id).Delete(&Role{}).Error
}

func (r *roleRepo) GetByID(id uint) (*Role, error) {
	var rl Role
	err := r.db.Where("id_role = ? AND status = true", id).First(&rl).Error
	return &rl, err
}

func (r *roleRepo) GetAll() ([]Role, error) {
	var roles []Role
	err := r.db.Where("status = true").Find(&roles).Error
	return roles, err
}
