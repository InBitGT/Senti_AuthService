package permission

type PermissionService interface {
	Create(req *Permission) (*Permission, error)
	Update(id uint, req *Permission) (*Permission, error)
	Delete(id uint) error
	GetByID(id uint) (*Permission, error)
	GetAll() ([]Permission, error)

	AssignRolePermission(roleID uint, permID uint) (*RolePermission, error)
	RemoveAssignment(id uint) error
}

type permissionService struct {
	repo PermissionRepository
}

func NewPermissionService(repo PermissionRepository) PermissionService {
	return &permissionService{repo}
}

func (s *permissionService) Create(req *Permission) (*Permission, error) {
	p := &Permission{
		Key:  req.Key,
		Desc: req.Desc,
	}
	return p, s.repo.Create(p)
}

func (s *permissionService) Update(id uint, req *Permission) (*Permission, error) {
	return s.repo.Update(id, req)
}

func (s *permissionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *permissionService) GetByID(id uint) (*Permission, error) {
	return s.repo.GetByID(id)
}

func (s *permissionService) GetAll() ([]Permission, error) {
	return s.repo.GetAll()
}

func (s *permissionService) AssignRolePermission(roleID uint, permID uint) (*RolePermission, error) {
	rp := &RolePermission{
		RoleID:       roleID,
		PermissionID: permID,
	}
	return rp, s.repo.Assign(rp)
}

func (s *permissionService) RemoveAssignment(id uint) error {
	return s.repo.RemoveAssignment(id)
}
