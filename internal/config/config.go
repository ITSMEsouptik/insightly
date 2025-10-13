package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func Load() *Config {
	_ = godotenv.Load() // Load .env file if it exists
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "development" // Default environment if not specified
	}
	cfg := &Config{
		Port: port,
		Env:  env,
	}
	log.Printf("Config loaded: env=%s port=%s", cfg.Env, cfg.Port)
	return cfg
}
