package auth

import (
	"net/http"

	"github.com/saurav11sarkar/resapi/config"
	"github.com/saurav11sarkar/resapi/rest/hendler/user"
)

func Routes(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", user.LoginUser)
	mux.HandleFunc("POST /login", user.LoginUser)

	return mux
}
