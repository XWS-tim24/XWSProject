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
		//		Port:                     "8080",
		//		DBHost:                   "reservation_db",
		//		DBPort:                   "5432",
		//		DBName:                   "ReservationServiceDB",
		//		DBUser:                   "postgres",
		//		DBPass:                   "loki123",

		Port:                     os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_PORT"),
		DBHost:                   os.Getenv("ACCOMMODATION_RESERVATION_DB_HOST"),
		DBPort:                   os.Getenv("ACCOMMODATION_RESERVATION_DB_PORT"),
		DBName:                   os.Getenv("ACCOMMODATION_RESERVATION_DB_NAME"),
		DBUser:                   os.Getenv("ACCOMMODATION_RESERVATION_DB_USER"),
		DBPass:                   os.Getenv("ACCOMMODATION_RESERVATION_DB_PASS"),
		AccommodationServiceHost: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_HOST"),
		AccommodationServicePort: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_PORT"),
	}
}
