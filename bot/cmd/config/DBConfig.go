package config

import (
	"os"
)

type Config struct {
	DBName     string
	DBUser     string
	DBPort     string
	DBPassword string
	DBHost     string
}

func NewConfig() (*Config, error) {
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	return &Config{DBName: name, DBUser: user, DBPassword: password, DBPort: port, DBHost: host}, nil
}
