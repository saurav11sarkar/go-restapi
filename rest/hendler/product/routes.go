package product

import (
	"net/http"

	"github.com/saurav11sarkar/resapi/config"
)

func Routes(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", CreateProduct)
	mux.HandleFunc("GET /", GetAllProducts)
	mux.HandleFunc("GET /{id}", GetProductById)

	return mux
}
