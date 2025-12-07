package tenant

import "github.com/gorilla/mux"

func SetUpTenantRoutes(api *mux.Router, h *TenantHandler) {
	r := api.PathPrefix("/tenants").Subrouter()

	r.HandleFunc("", h.Create).Methods("POST")
	r.HandleFunc("", h.GetAll).Methods("GET")
	r.HandleFunc("/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/{id}", h.Delete).Methods("DELETE")
}
