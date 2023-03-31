package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBName     string
	DBUser     string
	DBPort     int
	DBPassword string
	DBHost     string
}

func NewConfig() (*Config, error) {
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	host := os.Getenv("DB_HOST")
	if err != nil {
		return nil, err
	}
	return &Config{DBName: name, DBUser: user, DBPassword: password, DBPort: port, DBHost: host}, nil
}
