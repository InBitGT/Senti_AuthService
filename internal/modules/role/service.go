package role

type RoleService interface {
	Create(dto CreateRoleDTO) (*Role, error)
	Update(id uint, dto UpdateRoleDTO) (*Role, error)
	Delete(id uint) error
	HardDelete(id uint) error
	GetByID(id uint) (*Role, error)
	GetAll() ([]Role, error)
}

type roleService struct {
	repo RoleRepository
}

func NewRoleService(repo RoleRepository) RoleService {
	return &roleService{repo}
}

type CreateRoleDTO struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

type UpdateRoleDTO struct {
	Name   string `json:"name"`
	Desc   string `json:"description"`
	Status *bool  `json:"status,omitempty"` // opcional
}

func (s *roleService) Create(dto CreateRoleDTO) (*Role, error) {
	role := &Role{
		Name:   dto.Name,
		Desc:   dto.Desc,
		Status: true,
	}

	if err := s.repo.Create(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *roleService) Update(id uint, dto UpdateRoleDTO) (*Role, error) {
	// armamos objeto con lo que venga
	r := &Role{
		Name: dto.Name,
		Desc: dto.Desc,
	}

	// si viene status, lo seteamos; si no, no lo tocamos (lo dejará como default false si no lo pasamos),
	// por eso el repo.Update usa existing+updates, y aquí no forzamos status si no viene
	if dto.Status != nil {
		r.Status = *dto.Status
	}

	return s.repo.Update(id, r)
}

func (s *roleService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *roleService) HardDelete(id uint) error {
	return s.repo.HardDelete(id)
}

func (s *roleService) GetByID(id uint) (*Role, error) {
	return s.repo.GetByID(id)
}

func (s *roleService) GetAll() ([]Role, error) {
	return s.repo.GetAll()
}
