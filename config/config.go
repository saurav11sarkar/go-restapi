package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
		os.Exit(1)
	}

	return Config{
		Version:     os.Getenv("VERSION"),
		ServiceName: os.Getenv("SERVICENAME"),
		HttpPort:    os.Getenv("HTTPPORT"),
	}
}
