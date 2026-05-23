package rest

import (
	"net/http"

	"github.com/saurav11sarkar/resapi/rest/hendler"
	middlewares2 "github.com/saurav11sarkar/resapi/rest/middlewares"
)

func routes(mux *http.ServeMux) {
	manager := middlewares2.NewManager()
	testmiddlewares := manager.With(middlewares2.TestMiddleware)

	mux.Handle("GET /", testmiddlewares.Apply(http.HandlerFunc(hendler.HomeHandler)))
	mux.HandleFunc("POST /user", hendler.PostUser)
	mux.HandleFunc("GET /user", hendler.GetAllUsers)
	mux.HandleFunc("GET /user/{id}", hendler.GetUserById)
	mux.HandleFunc("PUT /user/{id}", hendler.UpdateUserById)
	mux.HandleFunc("DELETE /user/{id}", hendler.DeleteUserById)
	mux.HandleFunc("POST /product", hendler.CreateProduct)
	mux.HandleFunc("GET /product", hendler.GetAllProducts)
	mux.HandleFunc("GET /product/{id}", hendler.GetProductById)
	mux.HandleFunc("POST /login", hendler.LoginUser)
}
