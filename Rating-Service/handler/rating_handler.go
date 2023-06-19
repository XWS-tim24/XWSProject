package handler

import (
	"Rating-Service/service"
	"context"

	pb "github.com/XWS-tim24/Common/common/proto/rating_service"
)

type RatingHandler struct {
	pb.UnimplementedRatingServiceServer
	HostRatingService          *service.HostRatingService
	AccommodationRatingService *service.AccommodationRatingService
}

func (handler *RatingHandler) CreateHostRating(ctx context.Context, request *pb.CreateHostRatingRequest) (*pb.CreateHostRatingResponse, error) {
	rating := mapToNewHostRating(request.HostRating)
	err := handler.HostRatingService.Create(rating)
	if err != nil {
		return nil, err
	}
	return &pb.CreateHostRatingResponse{
		HostRating: mapToHostRatingPb(rating),
	}, nil
}

/*
  rpc CreateAccommodationRating(CreateAccommodationRatingRequest) returns(CreateAccommodationRatingResponse) {
    option (google.api.http) = {
      post: "/accommodationRating"
      body: "accommodationRating"
    };
  }

*/

func (handler *RatingHandler) CreateAccommodationRating(ctx context.Context, request *pb.CreateAccommodationRatingRequest) (*pb.CreateAccommodationRatingResponse, error) {
	rating := mapToNewAccommodationRating(request.AccommodationRating)
	err := handler.AccommodationRatingService.Create(rating)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccommodationRatingResponse{
		AccommodationRating: mapToAccommodationRatingPb(rating),
	}, nil
}

/*
func (handler *ReservationHandler) GetReservationById(ctx context.Context, request *pb.GetByIdRequest) (*pb.GetReservationByIdResponse, error) {
	id := request.Id
	log.Printf("Reservation with id %s", id)
	reservation, err := handler.ReservationService.GetById(id)
	if err != nil {
		return nil, err
	}
	ReservationPb := mapToReservationPb(reservation)
	response := &pb.GetReservationByIdResponse{
		Reservation: ReservationPb,
	}
	return response, nil
}



// REQUESTS
func (handler *ReservationHandler) GetRequestById(ctx context.Context, request *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	id := request.Id
	log.Printf("Reservation with id %s", id)
	reservationRequest, err := handler.ReservationRequestService.GetById(id)
	if err != nil {
		return nil, err
	}
	ReservationRequestPb := mapToReservationRequestPb(reservationRequest)
	response := &pb.GetByIdResponse{
		ReservationRequest: ReservationRequestPb,
	}
	return response, nil
}

func (handler *ReservationHandler) CreateRequest(ctx context.Context, request *pb.CreateReservationRequestRequest) (*pb.CreateReservationRequestResponse, error) {
	reservationRequest := mapToReservationRequest(request.ReservationRequest)
	println("accomodation id in handler %s", reservationRequest.AccomodationId)
	err := handler.ReservationRequestService.Create(reservationRequest)
	if err != nil {
		return nil, err
	}
	return &pb.CreateReservationRequestResponse{
		ReservationRequest: mapToReservationRequestPb(reservationRequest),
	}, nil
}

func (handler *ReservationHandler) DeleteReservationRequest(ctx context.Context, request *pb.GetByIdRequest) (*pb.ReservationRequestResponse, error) {
	id := request.Id
	log.Printf("Reservation request with id %s", id)
	err := handler.ReservationRequestService.Delete(id)
	if err != nil {
		return nil, err
	}
	reservationRequest, err2 := handler.ReservationRequestService.GetById(id)
	log.Printf("Get request with id %s", id)
	if err2 != nil {
		return nil, err
	}
	ReservationRequestPb := mapToReservationRequestPb(reservationRequest)
	response := &pb.ReservationRequestResponse{
		ReservationRequest: ReservationRequestPb,
	}
	return response, nil
}

func (handler *ReservationHandler) GetAllReservations(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllReservationsResponse, error) {
	reservations, err := handler.ReservationService.GetAll()
	if err != nil || *reservations == nil {
		return nil, err
	}
	response := &pb.GetAllReservationsResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, reservation := range *reservations {
		current := mapToReservationPb(&reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler *ReservationHandler) GetAllRequests(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllRequestsResponse, error) {
	reservationRequests, err := handler.ReservationRequestService.GetAll()
	if err != nil || *reservationRequests == nil {
		return nil, err
	}
	response := &pb.GetAllRequestsResponse{
		ReservationRequests: []*pb.ReservationRequest{},
	}
	for _, reservationRequest := range *reservationRequests {
		current := mapToReservationRequestPb(&reservationRequest)
		response.ReservationRequests = append(response.ReservationRequests, current)
	}
	return response, nil
}
*/
