package tenant

type TenantService interface {
	Create(req *Tenant) (*Tenant, error)
	Update(id uint, req *Tenant) (*Tenant, error)
	Delete(id uint) error
	HardDelete(id uint) error
	GetByID(id uint) (*Tenant, error)
	GetByCode(code string) (*Tenant, error)
	GetAll() ([]Tenant, error)
}

type tenantService struct {
	repo TenantRepository
}

func NewTenantService(repo TenantRepository) TenantService {
	return &tenantService{repo}
}

func (s *tenantService) Create(req *Tenant) (*Tenant, error) {
	t := &Tenant{
		Code:          req.Code,
		Name:          req.Name,
		Picture:       req.Picture,
		NIT:           req.NIT,
		Phone:         req.Phone,
		Email:         req.Email,
		SuscriptionID: req.SuscriptionID,
		AddressID:     req.AddressID,
		Status:        true,
	}
	return t, s.repo.Create(t)
}

func (s *tenantService) Update(id uint, req *Tenant) (*Tenant, error) {
	return s.repo.Update(id, req)
}

func (s *tenantService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *tenantService) HardDelete(id uint) error {
	return s.repo.HardDelete(id)
}

func (s *tenantService) GetByID(id uint) (*Tenant, error) {
	return s.repo.GetByID(id)
}

func (s *tenantService) GetByCode(code string) (*Tenant, error) {
	return s.repo.GetByCode(code)
}

func (s *tenantService) GetAll() ([]Tenant, error) {
	return s.repo.GetAll()
}
