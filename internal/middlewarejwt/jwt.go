package middlewarejwt

import (
	"context"
	"net/http"
	"strings"

	"AuthService/internal/common"
	"AuthService/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	ContextUserIDKey      contextKey = "user_id"
	ContextTenantIDKey    contextKey = "tenant_id"
	ContextRoleIDKey      contextKey = "role_id"
	ContextModulePermsKey contextKey = "module_permissions"
)

// ✅ Claims NUEVOS (por módulo)
type AuthClaims struct {
	UserID            uint              `json:"sub"`
	TenantID          uint              `json:"tenant"`
	RoleID            uint              `json:"role_id"`
	ModulePermissions map[uint][]string `json:"module_permissions"`
	jwt.RegisteredClaims
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		claims := &AuthClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		if claims.ModulePermissions == nil {
			claims.ModulePermissions = map[uint][]string{}
		}

		ctx := context.WithValue(r.Context(), ContextUserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, ContextTenantIDKey, claims.TenantID)
		ctx = context.WithValue(ctx, ContextRoleIDKey, claims.RoleID)
		ctx = context.WithValue(ctx, ContextModulePermsKey, claims.ModulePermissions)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
