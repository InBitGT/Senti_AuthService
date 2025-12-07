package server

import (
	"AuthService/internal/modules/auth"
	"AuthService/internal/modules/permission"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/tenant"
)

func (h *Handlers) GetAuthHandler() *auth.AuthHandler {
	return h.Auth
}

func (h *Handlers) GetRoleHandler() *role.RoleHandler {
	return h.Role
}

func (h *Handlers) GetPermissionHandler() *permission.PermissionHandler {
	return h.Permission
}

func (h *Handlers) GetTenantHandler() *tenant.TenantHandler {
	return h.Tenant
}
