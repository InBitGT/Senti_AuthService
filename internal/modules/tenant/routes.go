package tenant

import (
	"AuthService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetUpTenantRoutes(api *mux.Router, h *TenantHandler) {
	r := api.PathPrefix("/tenants").Subrouter()

	// Interno (hard delete) - protegido por X-Internal-Key
	internal := r.PathPrefix("/internal").Subrouter()
	internal.Use(middleware.InternalKeyMiddleware)
	internal.HandleFunc("/{id}", h.HardDeleteInternal).Methods("DELETE")

	// Público (soft delete)
	r.HandleFunc("", h.Create).Methods("POST")
	r.HandleFunc("", h.GetAll).Methods("GET")
	r.HandleFunc("/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/{id}", h.Delete).Methods("DELETE") // status=false
}
