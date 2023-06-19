package config

import "os"

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
	//ReservationServiceHost string
	//ReservationServicePort string
}

func NewConfig() *Config {
	return &Config{
		/*Port:   "8080",
		DBHost: "localhost",
		DBPort: "5432",
		DBName: "AccommodationServiceDB",
		DBUser: "postgres",
		DBPass: "loki123",*/
		Port:   os.Getenv("RATING_SERVICE_PORT"),
		DBHost: os.Getenv("RATING_DB_HOST"),
		DBPort: os.Getenv("RATING_DB_PORT"),
		DBName: os.Getenv("RATING_DB_NAME"),
		DBUser: os.Getenv("RATING_DB_USER"),
		DBPass: os.Getenv("RATING_DB_PASS"),
		//ReservationServiceHost: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_HOST"),
		//ReservationServicePort: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_PORT"),
	}
}
