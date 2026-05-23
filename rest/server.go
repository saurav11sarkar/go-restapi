package rest

import (
	"fmt"
	"net/http"

	"github.com/saurav11sarkar/resapi/config"
	"github.com/saurav11sarkar/resapi/rest/hendler"
	middlewares2 "github.com/saurav11sarkar/resapi/rest/middlewares"
)

func Start(cfg config.Config) {
	mux := http.NewServeMux()

	hendler.ConfigureAuth(cfg.JWTSecret, cfg.JWTExpires)
	routes(mux)

	manager := middlewares2.NewManager()

	handler := manager.
		With(
			middlewares2.Logger,
			middlewares2.EnableCORS,
		).
		Apply(mux)

	fmt.Println("Server is running on http://localhost:" + cfg.HttpPort)
	if err := http.ListenAndServe(":"+cfg.HttpPort, handler); err != nil {
		fmt.Println(err)
	}
}
