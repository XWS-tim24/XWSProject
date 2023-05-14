package startup

import (
	cfg "api-gateway/startup/config"
	"context"
	"fmt"
	"log"
	"net/http"

	reservationGw "github.com/XWS-tim24/Common/common/proto/accommodation_reservation_service"
	accommodationGw "github.com/XWS-tim24/Common/common/proto/accommodation_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationReservationHost, server.config.AccommodationReservationPort)
	err := reservationGw.RegisterAccommodationReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
	if err != nil {
		panic(err)
	}
	accommodationEmdpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err = accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEmdpoint, opts)
	if err != nil {
		panic(err)
	}

}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
