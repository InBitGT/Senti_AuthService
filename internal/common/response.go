package common

import (
	"encoding/json"
	"net/http"
)

// Estructura base para todas las respuestas
type APIResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Información del error
type ErrorInfo struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Data    *string `json:"data,omitempty"`
}

// Respuesta exitosa
func SuccessResponse(w http.ResponseWriter, message string, data interface{}, code string) {
	response := APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Respuesta exitosa para creación
func CreatedResponse(w http.ResponseWriter, message string, data interface{}, code string) {
	response := APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Respuesta de error
func ErrorResponse(w http.ResponseWriter, statusCode int, code string, message string, data *string) {
	response := ErrorInfo{
		Code:    code,
		Message: message,
		Data:    data,
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
