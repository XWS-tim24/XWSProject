package communication

import (
	pb "github.com/XWS-tim24/Common/common/proto/accommodation_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewAccommodationClient(address string) pb.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return pb.NewAccommodationServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
