package permission

type PermissionService interface {
	Create(dto CreatePermissionDTO) (*Permission, error)
	Update(id uint, dto UpdatePermissionDTO) (*Permission, error)
	Delete(id uint) error
	HardDelete(id uint) error
	GetByID(id uint) (*Permission, error)
	GetAll() ([]Permission, error)
}

type permissionService struct {
	repo PermissionRepository
}

func NewPermissionService(repo PermissionRepository) PermissionService {
	return &permissionService{repo}
}

type CreatePermissionDTO struct {
	Key  string `json:"key"`
	Desc string `json:"description"`
}

type UpdatePermissionDTO struct {
	Key  string `json:"key"`
	Desc string `json:"description"`
}

func (s *permissionService) Create(dto CreatePermissionDTO) (*Permission, error) {
	p := &Permission{
		Key:    dto.Key,
		Desc:   dto.Desc,
		Status: true,
	}

	if err := s.repo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *permissionService) Update(id uint, dto UpdatePermissionDTO) (*Permission, error) {
	p := &Permission{
		Key:  dto.Key,
		Desc: dto.Desc,
	}
	return s.repo.Update(id, p)
}

func (s *permissionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *permissionService) HardDelete(id uint) error {
	return s.repo.HardDelete(id)
}

func (s *permissionService) GetByID(id uint) (*Permission, error) {
	return s.repo.GetByID(id)
}

func (s *permissionService) GetAll() ([]Permission, error) {
	return s.repo.GetAll()
}
