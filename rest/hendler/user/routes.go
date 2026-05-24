package user

import (
	"net/http"

	"github.com/saurav11sarkar/resapi/config"
	"github.com/saurav11sarkar/resapi/rest/middlewares"
)

func Routes(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", PostUser)
	mux.HandleFunc("GET /", GetAllUsers)
	mux.HandleFunc("GET /{id}", GetUserById)
	mux.HandleFunc("PUT /{id}", UpdateUserById)
	mux.HandleFunc("DELETE /{id}", DeleteUserById)
	mux.HandleFunc("POST /login", LoginUser)
	mux.Handle("GET /profile", ProfileHandler(cfg))

	return mux
}

func ProfileHandler(cfg *config.Config) http.Handler {
	return middlewares.AuthMiddleware(cfg, "admin", "user")(http.HandlerFunc(Profile))
}
