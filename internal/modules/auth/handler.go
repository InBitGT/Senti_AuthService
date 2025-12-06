package auth

import (
	"AuthService/internal/common"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "JSON inválido", nil)
		return
	}

	user, err := h.service.Register(&req)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_REGISTER, user, common.HTTP_CREATED)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "JSON inválido", nil)
		return
	}

	resp, err := h.service.Login(&req)
	if err != nil {
		status := http.StatusUnauthorized
		if err == ErrInvalidOTP {
			status = http.StatusForbidden
		}
		common.ErrorResponse(w, status, common.ERR_INVALID_LOGIN, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_LOGIN, resp, common.HTTP_OK)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "JSON inválido", nil)
		return
	}

	resp, err := h.service.Refresh(req.RefreshToken)
	if err != nil {
		common.ErrorResponse(w, http.StatusUnauthorized, common.ERR_INVALID_TOKEN, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, resp, common.HTTP_OK)
}

func (h *AuthHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TenantCode string `json:"tenant_code"`
		Email      string `json:"email"`
		Channel    string `json:"channel"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "JSON inválido", nil)
		return
	}

	otpObj, err := h.service.GenerateOTP(req.TenantCode, req.Email, req.Channel)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_LOGIN, err.Error(), nil)
		return
	}

	resp := map[string]interface{}{
		"otp":        otpObj.Code,
		"expires_at": otpObj.ExpiresAt,
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, resp, common.HTTP_OK)
}

func (h *AuthHandler) ToggleTwoFA(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id_user"]
	id, _ := strconv.Atoi(idStr)

	var req struct {
		Enabled bool `json:"enabled"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "JSON inválido", nil)
		return
	}

	if err := h.service.ToggleTwoFA(uint(id), req.Enabled); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, map[string]bool{
		"two_fa_enabled": req.Enabled,
	}, common.HTTP_OK)
}
