package auth

import (
	"errors"
	"log"
	"time"

	"AuthService/internal/clients"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/tenant"

	"gorm.io/gorm"
)

var ErrRegisterCompany = errors.New("no se pudo registrar la empresa")

func (s *authService) RegisterCompany(req *RegisterCompanyRequest) (*RegisterCompanyResponse, error) {
	log.Println("[REGISTER_COMPANY] start")

	customer := clients.NewCustomerClient()
	payment := clients.NewPaymentClient()
	userc := clients.NewUserClient()

	// -------- VALIDACIONES --------
	if req.Tenant.Code == "" || req.Tenant.Name == "" {
		return nil, ErrRegisterCompany
	}
	if req.AdminUser.Email == "" || req.AdminUser.Password == "" {
		return nil, ErrRegisterCompany
	}
	if req.Subscription.PlanID == 0 {
		return nil, errors.New("plan_id es requerido")
	}

	// 0) Guard: evitar duplicado (idempotencia básica)
	if _, err := s.repo.GetActiveTenantByCode(req.Tenant.Code); err == nil {
		return nil, errors.New("tenant ya existe")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// -------- 1) ADDRESS --------
	log.Println("[REGISTER_COMPANY] creating address")
	addressID, err := customer.CreateAddress(clients.CreateAddressReq{
		Line1:      req.Address.Line1,
		Line2:      req.Address.Line2,
		City:       req.Address.City,
		State:      req.Address.State,
		Country:    req.Address.Country,
		PostalCode: req.Address.PostalCode,
	})
	if err != nil {
		log.Printf("[REGISTER_COMPANY] ERROR address: %v\n", err)
		return nil, err
	}

	// -------- 2) TENANT --------
	log.Println("[REGISTER_COMPANY] creating tenant")

	t := &tenant.Tenant{
		Code:      req.Tenant.Code,
		Name:      req.Tenant.Name,
		Picture:   req.Tenant.Picture,
		NIT:       req.Tenant.NIT,
		Phone:     req.Tenant.Phone,
		Email:     req.Tenant.Email,
		AddressID: addressID,
		Status:    true,
	}

	if err := s.repo.CreateTenant(t); err != nil {
		_ = customer.DeleteAddress(addressID)
		return nil, err
	}

	tenantID := t.ID

	// -------- 3) BRANCH GENERAL --------
	branchName := req.Branch.Name
	if branchName == "" {
		branchName = "General"
	}

	log.Println("[REGISTER_COMPANY] creating branch general")
	branchID, err := customer.CreateBranch(clients.CreateBranchReq{
		Name:        branchName,
		Description: req.Branch.Description,
		AddressID:   addressID,
		TenantID:    tenantID,
	})
	if err != nil {
		// compensación mínima
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = customer.DeleteAddress(addressID)
		return nil, err
	}

	// -------- 4) SUSCRIPTION --------
	now := time.Now()
	startedAt := now
	renewAt := now.Add(30 * 24 * time.Hour)

	if req.Subscription.StartedAt != nil {
		startedAt = *req.Subscription.StartedAt
	}
	if req.Subscription.RenewAt != nil {
		renewAt = *req.Subscription.RenewAt
	}

	log.Println("[REGISTER_COMPANY] creating suscription")
	suscriptionID, err := payment.CreateSuscription(clients.CreateSuscriptionReq{
		TenantID:  tenantID,
		PlanID:    req.Subscription.PlanID,
		StartedAt: startedAt,
		RenewAt:   renewAt,
		EndAt:     req.Subscription.EndAt,
	})
	if err != nil {
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = customer.DeleteAddress(addressID)
		return nil, err
	}

	// -------- 5) LINK TENANT -> SUSCRIPTION --------
	log.Println("[REGISTER_COMPANY] linking tenant suscription_id")
	if err := s.repo.UpdateTenantSuscription(tenantID, suscriptionID); err != nil {
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = customer.DeleteAddress(addressID)
		return nil, err
	}

	// -------- 6) ROLES (ADMIN/USER) --------
	log.Println("[REGISTER_COMPANY] ensuring roles exist")

	adminRole, err := s.repo.GetRoleByName(role.AdminName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			adminRole = &role.Role{
				Name:   role.AdminName,
				Desc:   "Administrador",
				Status: true,
			}
			if err := s.repo.CreateRole(adminRole); err != nil {
				_ = s.repo.HardDeleteTenant(tenantID)
				_ = customer.DeleteAddress(addressID)
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	usrRole, err := s.repo.GetRoleByName(role.UserName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			usrRole = &role.Role{
				Name:   role.UserName,
				Desc:   "Usuario",
				Status: true,
			}
			if err := s.repo.CreateRole(usrRole); err != nil {
				_ = s.repo.HardDeleteTenant(tenantID)
				_ = customer.DeleteAddress(addressID)
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// -------- 7) USER ADMIN (solo después de suscripción) --------
	log.Println("[REGISTER_COMPANY] creating admin user in UserService")

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
		_ = customer.DeleteAddress(addressID)
		return nil, err
	}

	// -------- 8) USER_BRANCH (asignar admin al branch general) --------
	log.Println("[REGISTER_COMPANY] creating user_branch assignment")
	_, err = customer.CreateUserBranch(clients.CreateUserBranchReq{
		UserID:   userID,
		BranchID: branchID,
	})
	if err != nil {
		_ = s.repo.HardDeleteTenant(tenantID)
		_ = customer.DeleteAddress(addressID)
		return nil, err
	}

	log.Println("[REGISTER_COMPANY] SUCCESS")

	return &RegisterCompanyResponse{
		TenantID:      tenantID,
		AddressID:     addressID,
		BranchID:      branchID,
		SuscriptionID: suscriptionID,
		UserID:        userID,
	}, nil
}
