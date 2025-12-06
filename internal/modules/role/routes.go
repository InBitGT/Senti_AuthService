package role

import "github.com/gorilla/mux"

func SetUpRoleRoutes(api *mux.Router, handler *RoleHandler) {
	roleRouter := api.PathPrefix("/roles").Subrouter()

	roleRouter.HandleFunc("", handler.Create).Methods("POST")
	roleRouter.HandleFunc("/{id}", handler.Update).Methods("PUT")
	roleRouter.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
	roleRouter.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	roleRouter.HandleFunc("/tenant/{tenant_id}", handler.GetByTenant).Methods("GET")
}
