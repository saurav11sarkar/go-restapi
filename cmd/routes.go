package cmd

import (
	"net/http"

	"github.com/saurav11sarkar/resapi/handler"
	"github.com/saurav11sarkar/resapi/middlewares"
)

func routes(mux *http.ServeMux) {
	manager := middlewares.NewManager()
	testmiddlewares := manager.With(middlewares.TestMiddleware)
	mux.Handle("GET /", testmiddlewares.Apply(http.HandlerFunc(handler.HomeHandler)))
	mux.HandleFunc("POST /user", handler.PostUser)
	mux.HandleFunc("GET /user", handler.GetAllUsers)
	mux.HandleFunc("GET /user/{id}", handler.GetUserById)
	mux.HandleFunc("PUT /user/{id}", handler.UpdateUserById)
	mux.HandleFunc("DELETE /user/{id}", handler.DeleteUserById)
}
