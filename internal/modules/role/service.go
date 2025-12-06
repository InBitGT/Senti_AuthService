package role

type RoleService interface {
	Create(dto CreateRoleDTO) (*Role, error)
	Update(id uint, dto UpdateRoleDTO) (*Role, error)
	Delete(id uint) error
	GetByID(id uint) (*Role, error)
	GetByTenant(tenantID uint) ([]Role, error)
}

type roleService struct {
	repo RoleRepository
}

func NewRoleService(repo RoleRepository) RoleService {
	return &roleService{repo}
}

type CreateRoleDTO struct {
	TenantID uint   `json:"tenant_id"`
	Name     string `json:"name"`
}

type UpdateRoleDTO struct {
	Name string `json:"name"`
}

func (s *roleService) Create(dto CreateRoleDTO) (*Role, error) {
	role := &Role{
		TenantID: dto.TenantID,
		Name:     dto.Name,
	}

	err := s.repo.Create(role)
	return role, err
}

func (s *roleService) Update(id uint, dto UpdateRoleDTO) (*Role, error) {
	role, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	role.Name = dto.Name

	err = s.repo.Update(role)
	return role, err
}

func (s *roleService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *roleService) GetByID(id uint) (*Role, error) {
	return s.repo.GetByID(id)
}

func (s *roleService) GetByTenant(tenantID uint) ([]Role, error) {
	return s.repo.GetByTenant(tenantID)
}
