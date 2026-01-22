package auth

import (
	"errors"
	"log"

	"AuthService/internal/clients"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/tenant"
)

var ErrRegisterCompany = errors.New("no se pudo registrar la empresa")

func (s *authService) RegisterCompany(req *RegisterCompanyRequest) (*RegisterCompanyResponse, error) {
	log.Println("[REGISTER_COMPANY] start")

	senti := clients.NewSentiClient()
	userc := clients.NewUserClient()

	// -------- VALIDACIÓN INICIAL --------
	if req.Tenant.Code == "" || req.Tenant.Name == "" {
		log.Println("[REGISTER_COMPANY] tenant inválido")
		return nil, ErrRegisterCompany
	}
	if req.AdminUser.Email == "" || req.AdminUser.Password == "" {
		log.Println("[REGISTER_COMPANY] admin user inválido")
		return nil, ErrRegisterCompany
	}

	// -------- 1) ADDRESS --------
	log.Println("[REGISTER_COMPANY] creando address")

	addressID, err := senti.CreateAddress(clients.CreateAddressReq{
		Line1:      req.Address.Line1,
		Line2:      req.Address.Line2,
		City:       req.Address.City,
		State:      req.Address.State,
		Country:    req.Address.Country,
		PostalCode: req.Address.PostalCode,
	})
	if err != nil {
		log.Printf("[REGISTER_COMPANY] ERROR creando address: %v\n", err)
		return nil, err
	}

	log.Printf("[REGISTER_COMPANY] address creado id=%d\n", addressID)

	// -------- 2) TENANT --------
	t := &tenant.Tenant{
		Code:    req.Tenant.Code,
		Name:    req.Tenant.Name,
		Picture: "", // opcional por ahora
		NIT:     req.Tenant.NIT,
		Phone:   req.Tenant.Phone,
		Email:   req.Tenant.Email,
		// SuscriptionID: 0, // opcional por ahora
		AddressID: addressID,
		Status:    true,
	}

	log.Println("[REGISTER_COMPANY] creando tenant")

	if err := s.repo.CreateTenant(t); err != nil {
		log.Printf("[REGISTER_COMPANY] ERROR creando tenant: %v\n", err)
		_ = senti.DeleteAddress(addressID)
		return nil, err
	}

	tenantID := t.ID
	log.Printf("[REGISTER_COMPANY] tenant creado id=%d\n", tenantID)

	// -------- 3) ROLES --------
	log.Println("[REGISTER_COMPANY] creando rol ADMIN")

	adminRole := &role.Role{
		Name:   role.AdminName,
		Desc:   "Administrador del tenant",
		Status: true,
	}

	if err := s.repo.CreateRole(adminRole); err != nil {
		log.Printf("[REGISTER_COMPANY] ERROR creando rol ADMIN: %v\n", err)
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = senti.DeleteAddress(addressID)
		return nil, err
	}

	log.Printf("[REGISTER_COMPANY] rol ADMIN creado id=%d\n", adminRole.ID)

	log.Println("[REGISTER_COMPANY] creando rol USER")

	userRole := &role.Role{
		Name:   role.UserName,
		Desc:   "Usuario estándar",
		Status: true,
	}

	if err := s.repo.CreateRole(userRole); err != nil {
		log.Printf("[REGISTER_COMPANY] ERROR creando rol USER: %v\n", err)
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = senti.DeleteAddress(addressID)
		return nil, err
	}

	// -------- 4) USER ADMIN --------
	log.Println("[REGISTER_COMPANY] creando usuario admin en UserService")

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
		log.Printf("[REGISTER_COMPANY] ERROR creando usuario admin: %v\n", err)
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = senti.DeleteAddress(addressID)
		return nil, err
	}

	log.Printf("[REGISTER_COMPANY] usuario admin creado id=%d\n", userID)

	log.Println("[REGISTER_COMPANY] SUCCESS")

	return &RegisterCompanyResponse{
		TenantID:  tenantID,
		AddressID: addressID,
		UserID:    userID,
	}, nil
}
