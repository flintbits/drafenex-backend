package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	JWTSecret   string
}

func Load() (*Config, error) {
	var err error = godotenv.Load()

	if err != nil {
		log.Println("Warning: .env file not found, using enviornment variables provided by the OS, Docker, or Kubernetes")
	}

	var config *Config = &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}

	return config, nil
}

//CREATE USER drafenex_platform_user WITH PASSWORD 'D3ad50ul_drafenex';
