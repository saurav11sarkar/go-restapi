package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
	JWTSecret   string
	JWTExpires  time.Duration
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if len(jwtSecret) < 32 {
		log.Fatal("JWT_SECRET must be set and at least 32 characters long")
	}

	jwtExpires := 15 * time.Minute
	if raw := os.Getenv("JWT_EXPIRES_IN"); raw != "" {
		parsed, err := time.ParseDuration(raw)
		if err != nil {
			log.Fatal("JWT_EXPIRES_IN must be a valid duration, for example 15m or 1h")
		}
		jwtExpires = parsed
	}

	return Config{
		Version:     os.Getenv("VERSION"),
		ServiceName: os.Getenv("SERVICENAME"),
		HttpPort:    os.Getenv("HTTPPORT"),
		JWTSecret:   jwtSecret,
		JWTExpires:  jwtExpires,
	}
}
