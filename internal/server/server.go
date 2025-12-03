package server

import (
	"net/http"

	"AuthService/internal/middleware"
	"AuthService/internal/routes"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		DB:     db,
	}

	s.Router.Use(middleware.CORS)

	handlers := s.initializeHandlers()
	routes.SetupRoutes(s.Router, handlers)

	s.Router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	return s
}
