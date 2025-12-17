package middleware

import (
	"net/http"
	"os"
)

func InternalKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := os.Getenv("INTERNAL_API_KEY")
		if expected == "" {
			http.Error(w, "internal key not configured", http.StatusInternalServerError)
			return
		}

		got := r.Header.Get("X-Internal-Key")
		if got == "" || got != expected {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
