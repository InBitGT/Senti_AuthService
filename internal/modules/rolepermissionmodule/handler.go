package rolepermissionmodule

import (
	"AuthService/internal/common"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) Assign(w http.ResponseWriter, r *http.Request) {
	roleID, _ := strconv.Atoi(mux.Vars(r)["role_id"])
	pmID, _ := strconv.Atoi(mux.Vars(r)["permission_module_id"])

	out, err := h.svc.Assign(uint(roleID), uint(pmID))
	if err != nil {
		common.ErrorResponse(w, 500, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, out, common.HTTP_CREATED)
}

func (h *Handler) Remove(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := h.svc.Remove(uint(id)); err != nil {
		common.ErrorResponse(w, 500, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}

func (h *Handler) GetByRole(w http.ResponseWriter, r *http.Request) {
	roleID, _ := strconv.Atoi(mux.Vars(r)["role_id"])

	out, err := h.svc.GetByRole(uint(roleID))
	if err != nil {
		common.ErrorResponse(w, 500, common.ERR_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, out, common.HTTP_OK)
}
