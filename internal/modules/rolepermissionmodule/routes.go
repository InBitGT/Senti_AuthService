package rolepermissionmodule

import "github.com/gorilla/mux"

func SetUpRolePermissionModuleRoutes(api *mux.Router, h *Handler) {
	r := api.PathPrefix("/role-permission-modules").Subrouter()

	r.HandleFunc("/assign/{role_id}/{permission_module_id}", h.Assign).Methods("POST")
	r.HandleFunc("/{id}", h.Remove).Methods("DELETE") // soft delete
	r.HandleFunc("/role/{role_id}", h.GetByRole).Methods("GET")
}
