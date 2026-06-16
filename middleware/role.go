package middleware

import (
	"net/http"
)

func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userRole, ok := r.Context().Value("user_role").(string)
			if !ok || userRole == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			allowed := false

			for _, role := range allowedRoles {
				if userRole == role {
					allowed = true
					break
				}
			}

			if !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}