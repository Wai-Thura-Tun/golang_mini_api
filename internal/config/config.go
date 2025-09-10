package config

import (
	"fmt"
	"os"
)

type Config struct {
	ServerPort   string
	JWTSecret    string
	Database_URL string
}

func LoadConfig() (*Config, error) {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	config := &Config{
		ServerPort:   os.Getenv("APP_PORT"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		Database_URL: dsn,
	}

	// validation
	if config.ServerPort == "" {
		return nil, fmt.Errorf("APP_PORT is not set")
	}

	if config.JWTSecret == "" {
		return nil, fmt.Errorf("JWT secret is not set")
	}

	if config.Database_URL == "" {
		return nil, fmt.Errorf("Database URL is not set")
	}

	return config, nil
}
