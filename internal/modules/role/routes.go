package role

import (
	"AuthService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetUpRoleRoutes(api *mux.Router, handler *RoleHandler) {
	r := api.PathPrefix("/roles").Subrouter()

	// interno hard delete
	internal := r.PathPrefix("/internal").Subrouter()
	internal.Use(middleware.InternalKeyMiddleware)
	internal.HandleFunc("/{id}", handler.HardDeleteInternal).Methods("DELETE")

	// público
	r.HandleFunc("", handler.Create).Methods("POST")
	r.HandleFunc("", handler.GetAll).Methods("GET")
	r.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	r.HandleFunc("/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/{id}", handler.Delete).Methods("DELETE") // status=false
}
