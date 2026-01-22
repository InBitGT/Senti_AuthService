package permissionmodule

import (
	"AuthService/internal/common"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	out, err := h.service.Create(req)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, out, common.HTTP_CREATED)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	out, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, out, common.HTTP_OK)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	out, err := h.service.GetByID(uint(id64))
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, msg, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, out, common.HTTP_OK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	if err := h.service.Delete(uint(id64)); err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, msg, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}

func (h *Handler) HardDeleteInternal(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	if err := h.service.HardDelete(uint(id64)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, "ok", common.HTTP_OK)
}
