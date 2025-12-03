// main.go
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"AuthService/db"
	"AuthService/internal/config"
	"AuthService/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	config.Init()

	database := db.Database()

	// migration.Migration()

	srv := server.NewServer(database)

	_ = godotenv.Load()

	httpServer := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      srv.Router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(httpServer.ListenAndServe())
}
