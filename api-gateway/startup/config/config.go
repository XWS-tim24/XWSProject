package config

type Config struct {
	Port                         string
	AccommodationHost            string
	AccommodationPort            string
	AccommodationReservationHost string
	AccommodationReservationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:                         "8000",
		AccommodationHost:            "localhost", //"accommodation_service",
		AccommodationPort:            "8080",
		AccommodationReservationHost: "localhost", //"accommodation_reservation_service",
		AccommodationReservationPort: "8080",
	}
}
