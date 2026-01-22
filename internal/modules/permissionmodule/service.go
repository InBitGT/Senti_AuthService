package permissionmodule

type Service interface {
	Create(dto CreateDTO) (*PermissionModule, error)
	GetByID(id uint) (*PermissionModule, error)
	GetAll() ([]PermissionModule, error)
	Delete(id uint) error
	HardDelete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

type CreateDTO struct {
	PermissionID uint `json:"permission_id"`
	ModuleID     uint `json:"module_id"`
}

func (s *service) Create(dto CreateDTO) (*PermissionModule, error) {
	pm := &PermissionModule{
		PermissionID: dto.PermissionID,
		ModuleID:     dto.ModuleID,
		Status:       true,
	}

	if err := s.repo.Create(pm); err != nil {
		return nil, err
	}
	return pm, nil
}

func (s *service) GetByID(id uint) (*PermissionModule, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAll() ([]PermissionModule, error) {
	return s.repo.GetAll()
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *service) HardDelete(id uint) error {
	return s.repo.HardDelete(id)
}
