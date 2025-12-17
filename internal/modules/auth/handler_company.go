package auth

import (
	"AuthService/internal/common"
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) RegisterCompany(w http.ResponseWriter, r *http.Request) {
	var req RegisterCompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	out, err := h.service.RegisterCompany(&req)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, out, common.HTTP_CREATED)
}
