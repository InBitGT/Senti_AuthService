package permission

import (
	"AuthService/internal/common"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PermissionHandler struct {
	svc PermissionService
}

func NewPermissionHandler(s PermissionService) *PermissionHandler {
	return &PermissionHandler{s}
}

func (h *PermissionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req Permission
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

func (h *PermissionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	out, err := h.svc.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, out, common.HTTP_OK)
}

func (h *PermissionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	out, err := h.svc.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.ERR_NOT_FOUND, "Permission not found", nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, out, common.HTTP_OK)
}

func (h *PermissionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var req Permission
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

func (h *PermissionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := h.svc.Delete(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}

func (h *PermissionHandler) AssignRolePermission(w http.ResponseWriter, r *http.Request) {
	roleID, _ := strconv.Atoi(mux.Vars(r)["role_id"])
	permID, _ := strconv.Atoi(mux.Vars(r)["permission_id"])

	rp, err := h.svc.AssignRolePermission(uint(roleID), uint(permID))
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, rp, common.HTTP_CREATED)
}

func (h *PermissionHandler) RemoveRolePermission(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := h.svc.RemoveAssignment(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}
