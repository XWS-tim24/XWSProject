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
		Port:                         os.Getenv("GATEWAY_PORT=8000"),
		AccommodationHost:            os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort:            os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationReservationHost: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_HOST"),
		AccommodationReservationPort: os.Getenv("ACCOMMODATION_RESERVATION_SERVICE_PORT"),
	}
}
