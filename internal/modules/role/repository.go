package role

import "gorm.io/gorm"

type RoleRepository interface {
	Create(role *Role) error
	Update(role *Role) error
	Delete(id uint) error
	GetByID(id uint) (*Role, error)
	GetByTenant(tenantID uint) ([]Role, error)
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

func (r *roleRepo) Update(role *Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepo) Delete(id uint) error {
	return r.db.Delete(&Role{}, id).Error
}

func (r *roleRepo) GetByID(id uint) (*Role, error) {
	var rl Role
	err := r.db.First(&rl, id).Error
	return &rl, err
}

func (r *roleRepo) GetByTenant(tenantID uint) ([]Role, error) {
	var roles []Role
	err := r.db.Where("tenant_id = ?", tenantID).Find(&roles).Error
	return roles, err
}
