package config

import "os"

type Config struct {
	Port                         string
	AccommodationHost            string
	AccommodationPort            string
	AccommodationReservationHost string
	AccommodationReservationPort string
}

func NewConfig() *Config {
	return &Config{
		/*	Port:                         "8000",
			AccommodationReservationHost: "localhost",
			AccommodationReservationPort: "8080",
			AccommodationHost:            os.Getenv("ACCOMMODATION_SERVICE_HOST"),
			AccommodationPort:            os.Getenv("ACCOMMODATION_SERVICE_PORT"),*/
		Port:                         os.Getenv("GATEWAY_PORT"),
		AccommodationHost:            os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort:            os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationReservationHost: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_HOST"),
		AccommodationReservationPort: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_PORT"),
	}
}
