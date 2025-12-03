package routes

import (
	"AuthService/internal/middleware"
	"AuthService/internal/modules/auth"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetAuthHandler() *auth.AuthHandler
}

func SetupRoutes(router *mux.Router, handlers RouteHandlers) {
	router.Use(middleware.ContentTypeJSON) // copias del proyecto base
	router.Use(middleware.Logger)
	router.Use(middleware.Recovery)

	api := router.PathPrefix("/api").Subrouter()

	auth.SetUpAuthRoutes(api, handlers.GetAuthHandler())
}
