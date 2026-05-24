package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/saurav11sarkar/resapi/config"
	"github.com/saurav11sarkar/resapi/utils"
)

func AuthMiddleware(cfg *config.Config, roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			parts := strings.Split(token, " ")
			if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			decoded, err := utils.VerifyToken(parts[1], cfg.JWTSecret)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			if len(roles) > 0 && !hasRole(decoded.Role, roles) {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), utils.ContextUserKey, decoded)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func hasRole(userRole string, allowedRoles []string) bool {
	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
