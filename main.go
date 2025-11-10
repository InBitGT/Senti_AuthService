package main

import (
	"AuthService/internal/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handler.HealthCheck)
	fmt.Println("Authorization service running on :8000")
	http.ListenAndServe(":8000", nil)
}
