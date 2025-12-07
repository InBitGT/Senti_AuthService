package tenant

import "gorm.io/gorm"

type TenantRepository interface {
	Create(t *Tenant) error
	Update(id uint, t *Tenant) (*Tenant, error)
	Delete(id uint) error
	GetByID(id uint) (*Tenant, error)
	GetByCode(code string) (*Tenant, error)
	GetAll() ([]Tenant, error)
}

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{db}
}

func (r *tenantRepository) Create(t *Tenant) error {
	return r.db.Create(t).Error
}

func (r *tenantRepository) Update(id uint, t *Tenant) (*Tenant, error) {
	var existing Tenant
	if err := r.db.First(&existing, id).Error; err != nil {
		return nil, err
	}

	existing.Name = t.Name
	existing.Code = t.Code
	existing.IsActive = t.IsActive

	return &existing, r.db.Save(&existing).Error
}

func (r *tenantRepository) Delete(id uint) error {
	return r.db.Delete(&Tenant{}, id).Error
}

func (r *tenantRepository) GetByID(id uint) (*Tenant, error) {
	var t Tenant
	err := r.db.First(&t, id).Error
	return &t, err
}

func (r *tenantRepository) GetByCode(code string) (*Tenant, error) {
	var t Tenant
	err := r.db.Where("code = ?", code).First(&t).Error
	return &t, err
}

func (r *tenantRepository) GetAll() ([]Tenant, error) {
	var t []Tenant
	err := r.db.Find(&t).Error
	return t, err
}
