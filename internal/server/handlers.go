package server

import (
	"AuthService/internal/config"
	"AuthService/internal/modules/auth"
)

func (s *Server) initializeHandlers() *Handlers {
	authRepo := auth.NewAuthRepository(s.DB)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	_ = config.JwtSecret // sólo para que sepas que ya está inicializado en main

	return &Handlers{
		Auth: authHandler,
	}
}
