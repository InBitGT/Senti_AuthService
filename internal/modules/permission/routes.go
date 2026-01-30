package permission

import (
	"AuthService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetUpPermissionRoutes(api *mux.Router, h *PermissionHandler) {
	r := api.PathPrefix("/permissions").Subrouter()

	// interno hard delete
	internal := r.PathPrefix("/internal").Subrouter()
	internal.Use(middleware.InternalKeyMiddleware)
	internal.HandleFunc("/{id}", h.HardDeleteInternal).Methods("DELETE")

	// público
	r.HandleFunc("", h.Create).Methods("POST")
	r.HandleFunc("", h.GetAll).Methods("GET")
	r.HandleFunc("/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/{id}", h.Delete).Methods("DELETE") // status=false
}
