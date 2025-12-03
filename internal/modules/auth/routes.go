package auth

import "github.com/gorilla/mux"

func SetUpAuthRoutes(api *mux.Router, handler *AuthHandler) {
	authRouter := api.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/register", handler.Register).Methods("POST")
	authRouter.HandleFunc("/login", handler.Login).Methods("POST")

}
