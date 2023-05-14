package config

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
		Port:   "8080",
		DBHost: "localhost",
		DBPort: "5432",
		DBName: "ReservationServiceDB",
		DBUser: "postgres",
		DBPass: "loki123",
	}
}
