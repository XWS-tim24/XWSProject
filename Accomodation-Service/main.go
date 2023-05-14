package main

import (
	"Accomodation-Service/startup"
	cfg "Accomodation-Service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
