package server

import (
	"AuthService/internal/config"
	"AuthService/internal/modules/auth"
	"AuthService/internal/modules/permission"
	"AuthService/internal/modules/permissionmodule"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/rolepermissionmodule"
	"AuthService/internal/modules/tenant"
)

func (s *Server) initializeHandlers() *Handlers {
	// AUTH
	authRepo := auth.NewAuthRepository(s.DB)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	// ROLES
	roleRepo := role.NewRoleRepository(s.DB)
	roleService := role.NewRoleService(roleRepo)
	roleHandler := role.NewRoleHandler(roleService)

	// PERMISSIONS
	permissionRepo := permission.NewPermissionRepository(s.DB)
	permissionService := permission.NewPermissionService(permissionRepo)
	permissionHandler := permission.NewPermissionHandler(permissionService)

	// PERMISSION_MODULE
	pmRepo := permissionmodule.NewRepository(s.DB)
	pmService := permissionmodule.NewService(pmRepo)
	pmHandler := permissionmodule.NewHandler(pmService)

	// TENANTS
	tenantRepo := tenant.NewTenantRepository(s.DB)
	tenantService := tenant.NewTenantService(tenantRepo)
	tenantHandler := tenant.NewTenantHandler(tenantService)

	rpmRepo := rolepermissionmodule.NewRepository(s.DB)
	rpmService := rolepermissionmodule.NewService(rpmRepo)
	rpmHandler := rolepermissionmodule.NewHandler(rpmService)

	_ = config.JwtSecret

	return &Handlers{
		Auth:                 authHandler,
		Role:                 roleHandler,
		Permission:           permissionHandler,
		PermissionModule:     pmHandler,
		Tenant:               tenantHandler,
		RolePermissionModule: rpmHandler,
	}
}
