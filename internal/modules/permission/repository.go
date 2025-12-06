package permission

import "gorm.io/gorm"

type PermissionRepository interface {
	Create(p *Permission) error
	Update(id uint, p *Permission) (*Permission, error)
	Delete(id uint) error
	GetByID(id uint) (*Permission, error)
	GetAll() ([]Permission, error)

	Assign(rolePermission *RolePermission) error
	RemoveAssignment(id uint) error
	GetPermissionsByRole(roleID uint) ([]Permission, error)
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
	if err := r.db.First(&existing, id).Error; err != nil {
		return nil, err
	}

	existing.Key = p.Key
	existing.Desc = p.Desc

	return &existing, r.db.Save(&existing).Error
}

func (r *permissionRepository) Delete(id uint) error {
	return r.db.Delete(&Permission{}, id).Error
}

func (r *permissionRepository) GetByID(id uint) (*Permission, error) {
	var out Permission
	err := r.db.First(&out, id).Error
	return &out, err
}

func (r *permissionRepository) GetAll() ([]Permission, error) {
	var list []Permission
	err := r.db.Find(&list).Error
	return list, err
}

func (r *permissionRepository) Assign(rp *RolePermission) error {
	return r.db.Create(rp).Error
}

func (r *permissionRepository) RemoveAssignment(id uint) error {
	return r.db.Delete(&RolePermission{}, id).Error
}

func (r *permissionRepository) GetPermissionsByRole(roleID uint) ([]Permission, error) {
	var list []Permission
	err := r.db.Raw(`
		SELECT p.* 
		FROM permissions p
		INNER JOIN role_permissions rp ON rp.id_permission = p.id_permission
		WHERE rp.id_role = ?
	`, roleID).Scan(&list).Error

	return list, err
}
