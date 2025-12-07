package tenant

import (
	"AuthService/internal/common"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TenantHandler struct {
	svc TenantService
}

func NewTenantHandler(s TenantService) *TenantHandler {
	return &TenantHandler{s}
}

func (h *TenantHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req Tenant
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "Invalid JSON", nil)
		return
	}

	out, err := h.svc.Create(&req)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, out, common.HTTP_CREATED)
}

func (h *TenantHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	out, err := h.svc.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, out, common.HTTP_OK)
}

func (h *TenantHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	out, err := h.svc.GetByID(uint(id))

	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.ERR_NOT_FOUND, "Tenant not found", nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, out, common.HTTP_OK)
}

func (h *TenantHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var req Tenant
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "Invalid JSON", nil)
		return
	}

	out, err := h.svc.Update(uint(id), &req)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, out, common.HTTP_OK)
}

func (h *TenantHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := h.svc.Delete(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}
