package auth

import (
	"AuthService/internal/common"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	user, err := h.service.Register(&req)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_CREATED, user, common.HTTP_CREATED)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	resp, err := h.service.Login(&req)
	if err != nil {
		status := http.StatusUnauthorized
		if err == ErrInvalidOTP {
			status = http.StatusForbidden
		}
		common.ErrorResponse(w, status, common.HTTP_UNAUTHORIZED, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, resp, common.HTTP_OK)
}
