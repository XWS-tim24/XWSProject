package config

import "os"

type Config struct {
	Port                     string
	DBHost                   string
	DBPort                   string
	DBName                   string
	DBUser                   string
	DBPass                   string
	AccommodationServiceHost string
	AccommodationServicePort string
}

func NewConfig() *Config {
	return &Config{
		Port:                     "8080",
		DBHost:                   "localhost",
		DBPort:                   "5432",
		DBName:                   "ReservationServiceDB",
		DBUser:                   "postgres",
		DBPass:                   "loki123",
		AccommodationServiceHost: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_HOST"),
		AccommodationServicePort: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_PORT"),
	}
}
