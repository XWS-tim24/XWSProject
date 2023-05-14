package config

import "os"

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
}

func NewConfig() *Config {
	return &Config{
		Port:   os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		DBHost: os.Getenv("ACCOMMODATION_DB_HOST"),
		DBPort: os.Getenv("ACCOMMODATION_DB_PORT"),
		DBName: os.Getenv("ACCOMMODATION_DB_NAME"),
		DBUser: os.Getenv("ACCOMMODATION_DB_USER"),
		DBPass: os.Getenv("ACCOMMODATION_DB_PASS"),
	}
}
