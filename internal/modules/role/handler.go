package role

import (
	"AuthService/internal/common"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RoleHandler struct {
	service RoleService
}

func NewRoleHandler(service RoleService) *RoleHandler {
	return &RoleHandler{service}
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRoleDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, 400, common.ERR_INVALID_JSON, "JSON inválido", nil)
		return
	}

	role, err := h.service.Create(req)
	if err != nil {
		common.ErrorResponse(w, 500, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, role, common.HTTP_CREATED)
}

func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	var req UpdateRoleDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, 400, common.ERR_INVALID_JSON, "JSON inválido", nil)
		return
	}

	role, err := h.service.Update(uint(id), req)
	if err != nil {
		common.ErrorResponse(w, 400, common.ERR_NOT_FOUND, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, role, common.HTTP_OK)
}

func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	if err := h.service.Delete(uint(id)); err != nil {
		common.ErrorResponse(w, 400, common.ERR_NOT_FOUND, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}

func (h *RoleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	role, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, 404, common.ERR_NOT_FOUND, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, role, common.HTTP_OK)
}

func (h *RoleHandler) GetByTenant(w http.ResponseWriter, r *http.Request) {
	tenantStr := mux.Vars(r)["tenant_id"]
	tenantID, _ := strconv.Atoi(tenantStr)

	roles, err := h.service.GetByTenant(uint(tenantID))
	if err != nil {
		common.ErrorResponse(w, 500, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, roles, common.HTTP_OK)
}
