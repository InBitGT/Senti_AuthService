package server

import (
	"AuthService/internal/modules/auth"
	"AuthService/internal/modules/role"
)

func (h *Handlers) GetAuthHandler() *auth.AuthHandler {
	return h.Auth
}

func (h *Handlers) GetRoleHandler() *role.RoleHandler {
	return h.Role
}
