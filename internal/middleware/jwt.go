package middleware

import (
	"context"
	"net/http"
	"strings"

	"AuthService/internal/common"
	"AuthService/internal/config"
	"AuthService/internal/modules/auth"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	ContextUserIDKey   contextKey = "user_id"
	ContextTenantIDKey contextKey = "tenant_id"
	ContextRoleKey     contextKey = "role"
	ContextPermsKey    contextKey = "permissions"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		claims := &auth.AuthClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, ContextTenantIDKey, claims.TenantID)
		ctx = context.WithValue(ctx, ContextRoleKey, claims.Role)
		ctx = context.WithValue(ctx, ContextPermsKey, claims.Permissions)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
