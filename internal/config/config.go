package config

import "os"

// Config holds the settings the program reads from its environment.
type Config struct {
	DatabaseURL string
	Port        string
	JWTSecret   string
}

// Load reads the configuration from environment variables.
func Load() (*Config, error) {
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}, nil
}
