package server

import "AuthService/internal/modules/auth"

func (h *Handlers) GetAuthHandler() *auth.AuthHandler {
	return h.Auth
}
