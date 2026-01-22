package routes

import (
	"AuthService/internal/middleware"
	"AuthService/internal/modules/auth"
	"AuthService/internal/modules/permission"
	"AuthService/internal/modules/permissionmodule"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/rolepermissionmodule"
	"AuthService/internal/modules/tenant"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetAuthHandler() *auth.AuthHandler
	GetRoleHandler() *role.RoleHandler
	GetPermissionHandler() *permission.PermissionHandler
	GetPermissionModuleHandler() *permissionmodule.Handler
	GetTenantHandler() *tenant.TenantHandler
	GetRolePermissionModuleHandler() *rolepermissionmodule.Handler
}

func SetupRoutes(router *mux.Router, handlers RouteHandlers) {
	router.Use(middleware.ContentTypeJSON)
	router.Use(middleware.Logger)
	router.Use(middleware.Recovery)

	api := router.PathPrefix("/api").Subrouter()

	auth.SetUpAuthRoutes(api, handlers.GetAuthHandler())
	role.SetUpRoleRoutes(api, handlers.GetRoleHandler())
	permission.SetUpPermissionRoutes(api, handlers.GetPermissionHandler())
	permissionmodule.SetUpPermissionModuleRoutes(api, handlers.GetPermissionModuleHandler())
	tenant.SetUpTenantRoutes(api, handlers.GetTenantHandler())
	rolepermissionmodule.SetUpRolePermissionModuleRoutes(api, handlers.GetRolePermissionModuleHandler())

}
