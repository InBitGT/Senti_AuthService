package routes

import (
	"AuthService/internal/middleware"
	"AuthService/internal/modules/auth"
	"AuthService/internal/modules/permission"
	"AuthService/internal/modules/role"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetAuthHandler() *auth.AuthHandler
	GetRoleHandler() *role.RoleHandler
	GetPermissionHandler() *permission.PermissionHandler
}

func SetupRoutes(router *mux.Router, handlers RouteHandlers) {
	router.Use(middleware.ContentTypeJSON)
	router.Use(middleware.Logger)
	router.Use(middleware.Recovery)

	api := router.PathPrefix("/api").Subrouter()

	// Auth routes
	auth.SetUpAuthRoutes(api, handlers.GetAuthHandler())

	// Role routes
	role.SetUpRoleRoutes(api, handlers.GetRoleHandler())

	// Permission routes
	permission.SetUpPermissionRoutes(api, handlers.GetPermissionHandler())
}
