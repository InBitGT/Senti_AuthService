package middlewarejwt

import (
	"net/http"

	"AuthService/internal/common"
)

func RequirePermission(permission string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			perms, ok := r.Context().Value(ContextPermsKey).([]string)
			if !ok {
				common.ErrorResponse(w, http.StatusForbidden, common.HTTP_FORBIDDEN, common.ERR_FORBIDDEN, nil)
				return
			}

			for _, p := range perms {
				if p == permission {
					next.ServeHTTP(w, r)
					return
				}
			}

			common.ErrorResponse(w, http.StatusForbidden, common.HTTP_FORBIDDEN, common.ERR_FORBIDDEN, nil)
		})
	}
}
