package rolepermissionmodule

type Service interface {
	Assign(roleID, permissionModuleID uint) (*RolePermissionModule, error)
	Remove(id uint) error
	GetByRole(roleID uint) ([]RolePermissionModule, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Assign(roleID, permissionModuleID uint) (*RolePermissionModule, error) {
	rp := &RolePermissionModule{
		RoleID:             roleID,
		PermissionModuleID: permissionModuleID,
		Status:             true,
	}
	return rp, s.repo.Create(rp)
}

func (s *service) Remove(id uint) error {
	return s.repo.SoftDelete(id)
}

func (s *service) GetByRole(roleID uint) ([]RolePermissionModule, error) {
	return s.repo.GetByRole(roleID)
}
