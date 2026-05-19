package cmd

import (
	"github.com/saurav11sarkar/resapi/config"
	"github.com/saurav11sarkar/resapi/rest"
)

func Serve() {
	cfg := config.LoadConfig()
	rest.Start(cfg)
}
