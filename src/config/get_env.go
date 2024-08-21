package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port           int
	DatabaseConfig string
}

func GetEnv() *Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	// databaseConfig := os.Getenv("DATABASE_CONFIG")
	databaseConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)

	return &Config{
		Port:           port,
		DatabaseConfig: databaseConfig,
	}
}
