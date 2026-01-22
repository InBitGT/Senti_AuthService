package server

import (
	"AuthService/internal/modules/auth"
	"AuthService/internal/modules/permission"
	"AuthService/internal/modules/permissionmodule"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/rolepermissionmodule"
	"AuthService/internal/modules/tenant"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	DB     *gorm.DB
}

type Handlers struct {
	Auth                 *auth.AuthHandler
	Role                 *role.RoleHandler
	Permission           *permission.PermissionHandler
	PermissionModule     *permissionmodule.Handler
	Tenant               *tenant.TenantHandler
	RolePermissionModule *rolepermissionmodule.Handler
}
