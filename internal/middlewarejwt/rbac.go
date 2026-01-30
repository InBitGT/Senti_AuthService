package middlewarejwt

import (
	"net/http"

	"AuthService/internal/common"
)

// ✅ valida "perm" dentro de un moduleID específico
func RequireModulePermission(moduleID uint, perm string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			mp, ok := r.Context().Value(ContextModulePermsKey).(map[uint][]string)
			if !ok || mp == nil {
				common.ErrorResponse(w, http.StatusForbidden, common.HTTP_FORBIDDEN, common.ERR_FORBIDDEN, nil)
				return
			}

			perms := mp[moduleID]
			for _, p := range perms {
				if p == perm {
					next.ServeHTTP(w, r)
					return
				}
			}

			common.ErrorResponse(w, http.StatusForbidden, common.HTTP_FORBIDDEN, common.ERR_FORBIDDEN, nil)
		})
	}
}
