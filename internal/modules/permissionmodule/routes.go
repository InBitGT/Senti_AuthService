package permissionmodule

import (
	"AuthService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetUpPermissionModuleRoutes(api *mux.Router, h *Handler) {
	r := api.PathPrefix("/permission-modules").Subrouter()

	// interno hard delete
	internal := r.PathPrefix("/internal").Subrouter()
	internal.Use(middleware.InternalKeyMiddleware)
	internal.HandleFunc("/{id}", h.HardDeleteInternal).Methods("DELETE")

	// público
	r.HandleFunc("", h.Create).Methods("POST")
	r.HandleFunc("", h.GetAll).Methods("GET")
	r.HandleFunc("/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/{id}", h.Delete).Methods("DELETE")
}
