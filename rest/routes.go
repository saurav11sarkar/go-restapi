package rest

import (
	"net/http"

	"github.com/saurav11sarkar/resapi/config"
	"github.com/saurav11sarkar/resapi/rest/hendler/auth"
	"github.com/saurav11sarkar/resapi/rest/hendler/product"
	"github.com/saurav11sarkar/resapi/rest/hendler/user"
)

func routes(prefix string, mux *http.ServeMux, cfg *config.Config) {
	mountRouter(mux, "/"+prefix+"/user", user.Routes(cfg))
	mountRouter(mux, "/"+prefix+"/auth", auth.Routes(cfg))
	mountRouter(mux, "/"+prefix+"/product", product.Routes(cfg))
}

func mountRouter(mux *http.ServeMux, path string, router http.Handler) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "" {
				r.URL.Path = "/"
			}
			router.ServeHTTP(w, r)
		})).ServeHTTP(w, r)
	})

	mux.Handle(path, handler)
	mux.Handle(path+"/", handler)
}
