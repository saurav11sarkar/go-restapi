package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/saurav11sarkar/resapi/internal/model"
	"github.com/saurav11sarkar/resapi/utils"
)

type contextKey string

const authUserKey contextKey = "authUser"

func Auth(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.SendJson(w, http.StatusUnauthorized, model.Response{
					Success: false,
					Message: "Authorization header is required",
					Data:    nil,
				})
				return
			}

			parts := strings.Fields(authHeader)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				utils.SendJson(w, http.StatusUnauthorized, model.Response{
					Success: false,
					Message: "Authorization header must be Bearer token",
					Data:    nil,
				})
				return
			}

			claims, err := utils.VerifyToken(parts[1], secret)
			if err != nil {
				utils.SendJson(w, http.StatusUnauthorized, model.Response{
					Success: false,
					Message: "Invalid or expired token",
					Data:    nil,
				})
				return
			}

			ctx := context.WithValue(r.Context(), authUserKey, model.PublicUser{
				Id:    claims.UserID,
				Email: claims.Email,
			})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthUser(r *http.Request) (model.PublicUser, bool) {
	user, ok := r.Context().Value(authUserKey).(model.PublicUser)
	return user, ok
}
