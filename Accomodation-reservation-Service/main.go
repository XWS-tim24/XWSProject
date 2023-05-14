package main

/*
import (
	"Accomodation-reservation-Service/domain"
	"Accomodation-reservation-Service/handler"
	"Accomodation-reservation-Service/repo"
	"Accomodation-reservation-Service/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=loki123 dbname=ReservationServiceDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	err = db.AutoMigrate(&domain.Reservation{})
	if err != nil {
		print(err)
		return nil
	}

	err = db.AutoMigrate(&domain.ReservationRequest{})
	if err != nil {
		print(err)
		return nil
	}

	return db
}

func startServer(resHandler *handler.ReservationHandler, reqHandler *handler.ReservationRequestHandler) {
	router := mux.NewRouter().StrictSlash(true)

.	router.HandleFunc("/reservationRequests/{id}", reqHandler.GetById).Methods("GET")
.	router.HandleFunc("/reservationRequests/pending/user/{userId}", reqHandler.GetAllPendingForUser).Methods("GET")
.	router.HandleFunc("/reservationRequests/pending/accomodation/{accomodationId}", reqHandler.GetAllPendingForAccomodation).Methods("GET")

.	router.HandleFunc("/reservations/{id}", resHandler.GetById).Methods("GET")
.	router.HandleFunc("/reservations/canceled/{userId}", resHandler.GetNumberOfCanceled).Methods("GET")
	router.HandleFunc("/reservations", resHandler.Create).Methods("POST")
.	router.HandleFunc("/reservationRequests", reqHandler.Create).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Hello, World")
	//	config := config.NewConfig()
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	resRepo := &repo.ReservationRepository{DatabaseConnection: database}
	reqRepo := &repo.ReservationRequestRepository{DatabaseConnection: database}

	resService := &service.ReservationService{ReservationRepo: resRepo, ReservationRequestRepo: reqRepo}
	reqService := &service.ReservationRequestService{ReservationRequestRepo: reqRepo, ReservationService: resService}

	resHandler := &handler.ReservationHandler{ReservationService: resService, ReservationRequestService: reqService}

	//err := reqService.Accept("40369432-3454-47b3-9042-2df11c272b52")
	err := resService.Cancel("485624e5-9107-4330-94aa-9c01842f2858", "user2")
	fmt.Print(err)
	//startServer(resHandler, reqHandler)

	//	server := startup.NewServer(config)
	//	server.Start()
}

/*
	func (server *Server) initPostgresClient() *gorm.DB {
		client, err := persistence.GetClient(
			server.config.InventoryDBHost, server.config.InventoryDBUser,
			server.config.InventoryDBPass, server.config.InventoryDBName,
			server.config.InventoryDBPort)
		if err != nil {
			log.Fatal(err)
		}
		return client
	}
*/

import (
	"Accomodation-reservation-Service/startup"
	cfg "Accomodation-reservation-Service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
