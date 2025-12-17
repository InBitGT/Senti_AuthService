package auth

import "github.com/gorilla/mux"

func SetUpAuthRoutes(api *mux.Router, handler *AuthHandler) {
	authRouter := api.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/register", handler.Register).Methods("POST")
	authRouter.HandleFunc("/login", handler.Login).Methods("POST")
	authRouter.HandleFunc("/refresh", handler.Refresh).Methods("POST")
	authRouter.HandleFunc("/otp/send", handler.SendOTP).Methods("POST")
	authRouter.HandleFunc("/2fa/{id_user}", handler.ToggleTwoFA).Methods("PUT")
	authRouter.HandleFunc("/register-company", handler.RegisterCompany).Methods("POST")
}
