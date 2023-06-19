package startup

import (
	"Rating-Service/domain"
	"Rating-Service/handler"
	"Rating-Service/repo"
	"Rating-Service/startup/config"
	"fmt"
	"log"
	"net"

	//accommodationGw "github.com/XWS-tim24/Common/common/proto/accommodation_service"
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
	err = client.AutoMigrate(&domain.Accommodation{})
	if err != nil {
		print(err)
		return nil
	}

	err = client.AutoMigrate(&domain.AvailableDate{})
	if err != nil {
		print(err)
		return nil
	}

	return client
}

func (server *Server) Start() {
	postgresClient := server.initPostgresClient()

	accommodationRatingRepository := &repo.AccommodationRatingRepository{DatabaseConnection: postgresClient}
	//accommodationRepository := &repo.AccommodationRepository{DatabaseConnection: postgresClient}

	//reservationServiceAddress := fmt.Sprintf("%s:%s", server.config.ReservationServiceHost, server.config.ReservationServicePort)
	/*availableDateService := &service.AvailableDateService{AvailableDateRepository: availableDateRepository, AccommodationRepository: accommodationRepository, ReservationServiceAddress: reservationServiceAddress}
	accommodationService := &service.AccommodationService{AccommodationRepository: accommodationRepository, AvailableRepository: availableDateRepository}

	accommodationHandler := &handler.AccommodationHandler{AccommodationService: accommodationService, AvailableDateService: availableDateService}
	server.startGrpcServer(accommodationHandler)
	fmt.Println("Finished")*/
}

func (server *Server) startGrpcServer(accommodationHandler *handler.AccommodationHandler) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	//accommodationGw.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
