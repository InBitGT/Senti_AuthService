package auth

import (
	"errors"

	"AuthService/internal/clients"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/tenant"
)

var ErrRegisterCompany = errors.New("no se pudo registrar la empresa")

func (s *authService) RegisterCompany(req *RegisterCompanyRequest) (*RegisterCompanyResponse, error) {
	senti := clients.NewSentiClient()
	userc := clients.NewUserClient()

	if req.Tenant.Code == "" || req.Tenant.Name == "" || req.AdminUser.Email == "" || req.AdminUser.Password == "" {
		return nil, ErrRegisterCompany
	}

	// 1) Crear una sola address (compartida)
	addressID, err := senti.CreateAddress(clients.CreateAddressReq{
		Line1:      req.Address.Line1,
		Line2:      req.Address.Line2,
		City:       req.Address.City,
		State:      req.Address.State,
		Country:    req.Address.Country,
		PostalCode: req.Address.PostalCode,
	})
	if err != nil {
		return nil, err
	}

	var tenantID uint
	defer func() {
		// Si algo falla después, esto evita basura.
		// Se activa solo si devolvemos error y ya habíamos creado cosas.
	}()

	// 2) Crear tenant en AuthService
	t := &tenant.Tenant{
		Code:      req.Tenant.Code,
		Name:      req.Tenant.Name,
		NIT:       req.Tenant.NIT,
		Phone:     req.Tenant.Phone,
		Email:     req.Tenant.Email,
		AddressID: addressID,
		IsActive:  true,
	}

	if err := s.repo.CreateTenant(t); err != nil {
		_ = senti.DeleteAddress(addressID)
		return nil, err
	}
	tenantID = t.ID

	// 3) Crear roles base del tenant
	adminRole := &role.Role{
		TenantID: tenantID,
		Name:     role.AdminName,
		Desc:     "Administrador del tenant",
	}

	if err := s.repo.CreateRole(adminRole); err != nil {
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = senti.DeleteAddress(addressID)
		return nil, err
	}

	userRole := &role.Role{
		TenantID: tenantID,
		Name:     role.UserName,
		Desc:     "Usuario estándar",
	}

	if err := s.repo.CreateRole(userRole); err != nil {
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = senti.DeleteAddress(addressID)
		return nil, err
	}

	// 4) Crear usuario admin en UserService (misma address)
	userID, err := userc.CreateAdmin(clients.CreateAdminReq{
		TenantID:  tenantID,
		AddressID: addressID,
		Email:     req.AdminUser.Email,
		Password:  req.AdminUser.Password,
		FirstName: req.AdminUser.FirstName,
		LastName:  req.AdminUser.LastName,
		Phone:     req.AdminUser.Phone,
		RoleID:    adminRole.ID,
	})
	if err != nil {
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = senti.DeleteAddress(addressID)
		return nil, err
	}

	return &RegisterCompanyResponse{
		TenantID:  tenantID,
		AddressID: addressID,
		UserID:    userID,
	}, nil
}
