package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	JWTSecret    string
	Database_URL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errror loading .env file")
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

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
