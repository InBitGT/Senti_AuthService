package server

import (
	"AuthService/internal/config"
	"AuthService/internal/modules/auth"
	"AuthService/internal/modules/role"
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

	_ = config.JwtSecret

	return &Handlers{
		Auth: authHandler,
		Role: roleHandler,
	}
}
