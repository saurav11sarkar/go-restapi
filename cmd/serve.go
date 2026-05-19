package cmd

import (
	"fmt"
	"net/http"

	"github.com/saurav11sarkar/resapi/config"
	"github.com/saurav11sarkar/resapi/middlewares"
)

func Serve() {
	mux := http.NewServeMux()

	routes(mux)

	manager := middlewares.NewManager()

	handler := manager.
		With(
			middlewares.Logger,
			middlewares.EnableCORS,
		).
		Apply(mux)

	cfg := config.LoadConfig()

	fmt.Println("Server is running on http://localhost:" + cfg.HttpPort)
	if err := http.ListenAndServe(":"+cfg.HttpPort, handler); err != nil {
		fmt.Println(err)
	}
}
