package cmd

import (
	"fmt"
	"net/http"

	"github.com/saurav11sarkar/resapi/middlewares"
)

func Serve() {
	mux := http.NewServeMux()

	routes(mux)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", middlewares.EnableCORS(mux)); err != nil {
		fmt.Println(err)
	}
}
