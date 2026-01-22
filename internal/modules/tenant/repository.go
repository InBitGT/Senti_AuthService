package tenant

import "gorm.io/gorm"

type TenantRepository interface {
	Create(t *Tenant) error
	Update(id uint, t *Tenant) (*Tenant, error)
	Delete(id uint) error     // soft delete => status=false
	HardDelete(id uint) error // físico
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
	if err := r.db.Where("id_tenant = ? AND status = true", id).First(&existing).Error; err != nil {
		return nil, err
	}

	updates := map[string]interface{}{
		"code":           t.Code,
		"name":           t.Name,
		"picture":        t.Picture,
		"nit":            t.NIT,
		"phone":          t.Phone,
		"email":          t.Email,
		"suscription_id": t.SuscriptionID,
		"address_id":     t.AddressID,
		"status":         t.Status, // si querés permitir re-activar
	}

	if err := r.db.Model(&Tenant{}).Where("id_tenant = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

// ✅ soft delete
func (r *tenantRepository) Delete(id uint) error {
	res := r.db.Model(&Tenant{}).
		Where("id_tenant = ? AND status = true", id).
		Updates(map[string]interface{}{"status": false})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// ✅ hard delete físico
func (r *tenantRepository) HardDelete(id uint) error {
	return r.db.Unscoped().Where("id_tenant = ?", id).Delete(&Tenant{}).Error
}

func (r *tenantRepository) GetByID(id uint) (*Tenant, error) {
	var t Tenant
	err := r.db.Where("id_tenant = ? AND status = true", id).First(&t).Error
	return &t, err
}

func (r *tenantRepository) GetByCode(code string) (*Tenant, error) {
	var t Tenant
	err := r.db.Where("code = ? AND status = true", code).First(&t).Error
	return &t, err
}

func (r *tenantRepository) GetAll() ([]Tenant, error) {
	var list []Tenant
	err := r.db.Where("status = true").Find(&list).Error
	return list, err
}
