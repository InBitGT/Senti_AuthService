package server

import (
	"AuthService/internal/modules/auth"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	DB     *gorm.DB
}

type Handlers struct {
	Auth *auth.AuthHandler
}
