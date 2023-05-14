package startup

import (
	"Accomodation-reservation-Service/domain"
	"Accomodation-reservation-Service/handler"
	"Accomodation-reservation-Service/repo"
	"Accomodation-reservation-Service/service"
	"Accomodation-reservation-Service/startup/config"
	"fmt"
	"log"
	"net"

	reservationGw "github.com/XWS-tim24/Common/common/proto/accommodation_reservation_service"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) getClient(host, user, password, dbname, port string) (*gorm.DB, error) {
	//host = "localhost"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func (server *Server) initPostgresClient() *gorm.DB {
	client, err := server.getClient(
		server.config.DBHost, server.config.DBUser,
		server.config.DBPass, server.config.DBName,
		server.config.DBPort)
	if err != nil {
		log.Fatal(err)
	}
	err = client.AutoMigrate(&domain.Reservation{})
	if err != nil {
		print(err)
		return nil
	}

	err = client.AutoMigrate(&domain.ReservationRequest{})
	if err != nil {
		print(err)
		return nil
	}

	return client
}

func (server *Server) Start() {
	database := server.initPostgresClient()

	resRepo := &repo.ReservationRepository{DatabaseConnection: database}
	reqRepo := &repo.ReservationRequestRepository{DatabaseConnection: database}
	accommodationServiceAddress := fmt.Sprintf("%s:%s", server.config.AccommodationServiceHost, server.config.AccommodationServicePort)
	resService := &service.ReservationService{ReservationRepo: resRepo, ReservationRequestRepo: reqRepo, AccommodationServiceAddres: accommodationServiceAddress}
	reqService := &service.ReservationRequestService{ReservationRequestRepo: reqRepo, ReservationService: resService, AccommodationServiceAddres: accommodationServiceAddress}

	resHandler := &handler.ReservationHandler{ReservationService: resService, ReservationRequestService: reqService}

	server.startGrpcServer(resHandler)
	fmt.Println("Finished")
}

func (server *Server) startGrpcServer(reservationHandler *handler.ReservationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reservationGw.RegisterAccommodationReservationServiceServer(grpcServer, reservationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	/*	accommodationGw.RegisterAccommodationServiceServer(grpcServer, availableDateHandler)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}*/

}
