package rest

import (
	"net/http"

	"github.com/saurav11sarkar/resapi/config"
	"github.com/saurav11sarkar/resapi/rest/hendler"
	middlewares2 "github.com/saurav11sarkar/resapi/rest/middlewares"
)

func routes(mux *http.ServeMux, cfg config.Config) {
	manager := middlewares2.NewManager()
	testmiddlewares := manager.With(middlewares2.TestMiddleware)
	authMiddlewares := middlewares2.NewManager().With(middlewares2.Auth(cfg.JWTSecret))

	mux.Handle("GET /", testmiddlewares.Apply(http.HandlerFunc(hendler.HomeHandler)))
	mux.HandleFunc("POST /user", hendler.PostUser)
	mux.Handle("GET /user", authMiddlewares.Apply(http.HandlerFunc(hendler.GetAllUsers)))
	mux.Handle("GET /user/{id}", authMiddlewares.Apply(http.HandlerFunc(hendler.GetUserById)))
	mux.Handle("PUT /user/{id}", authMiddlewares.Apply(http.HandlerFunc(hendler.UpdateUserById)))
	mux.Handle("DELETE /user/{id}", authMiddlewares.Apply(http.HandlerFunc(hendler.DeleteUserById)))
	mux.Handle("POST /product", authMiddlewares.Apply(http.HandlerFunc(hendler.CreateProduct)))
	mux.HandleFunc("GET /product", hendler.GetAllProducts)
	mux.HandleFunc("GET /product/{id}", hendler.GetProductById)
	mux.HandleFunc("POST /login", hendler.LoginUser)
}
