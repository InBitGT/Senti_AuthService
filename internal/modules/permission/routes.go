package permission

import "github.com/gorilla/mux"

func SetUpPermissionRoutes(api *mux.Router, h *PermissionHandler) {
	r := api.PathPrefix("/permissions").Subrouter()

	r.HandleFunc("", h.Create).Methods("POST")
	r.HandleFunc("", h.GetAll).Methods("GET")
	r.HandleFunc("/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/{id}", h.Delete).Methods("DELETE")

	r.HandleFunc("/assign/{role_id}/{permission_id}", h.AssignRolePermission).Methods("POST")
	r.HandleFunc("/assign/{id}", h.RemoveRolePermission).Methods("DELETE")
}
