package server

import (
	"AuthService/internal/modules/auth"
	"AuthService/internal/modules/permission"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/tenant"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	DB     *gorm.DB
}

type Handlers struct {
	Auth       *auth.AuthHandler
	Role       *role.RoleHandler
	Permission *permission.PermissionHandler
	Tenant     *tenant.TenantHandler
}
