package handler

import (
	"Accomodation-reservation-Service/service"
	"context"
	"log"

	pb "github.com/XWS-tim24/Common/common/proto/accommodation_reservation_service"
)

type ReservationHandler struct {
	pb.UnimplementedAccommodationReservationServiceServer
	ReservationService        *service.ReservationService
	ReservationRequestService *service.ReservationRequestService
}

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

func (handler *ReservationHandler) CreateReservation(ctx context.Context, request *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	reservation := mapToReservation(request.Reservation)
	err := handler.ReservationService.Create(reservation)
	if err != nil {
		return nil, err
	}
	return &pb.CreateReservationResponse{
		Reservation: mapToReservationPb(reservation),
	}, nil
}

func (handler *ReservationHandler) GetNumberOfCanceled(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetIntResponse, error) {
	userId := request.UserId
	log.Printf("Reservation for user with id %s", userId)
	number_of_canceled := handler.ReservationService.GetNumberOfCanceled(userId)

	response := &pb.GetIntResponse{
		Num: int32(number_of_canceled),
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
	err := handler.ReservationRequestService.Create(reservationRequest)
	if err != nil {
		return nil, err
	}
	return &pb.CreateReservationRequestResponse{
		ReservationRequest: mapToReservationRequestPb(reservationRequest),
	}, nil
}

func (handler *ReservationHandler) GetAllPendingForUser(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetAllPendingForUserResponse, error) {
	userId := request.UserId
	pending_requests := handler.ReservationRequestService.GetAllPendingForUser(userId)

	response := &pb.GetAllPendingForUserResponse{
		ReservationRequest: []*pb.ReservationRequest{},
	}
	///PROVERI OVE POKAZIVACE
	for _, reservationRequest := range *pending_requests {
		current := mapToReservationRequestPb(&reservationRequest)
		response.ReservationRequest = append(response.ReservationRequest, current)
	}
	return response, nil

}

func (handler *ReservationHandler) GetAllPendingForAccomodation(ctx context.Context, request *pb.GetAllPendingForAccRequest) (*pb.GetAllPendingForAccResponse, error) {
	accomodationId := request.AccomodationId
	pending_requests := handler.ReservationRequestService.GetAllPendingForAccomodation(accomodationId)
	response := &pb.GetAllPendingForAccResponse{
		ReservationRequest: []*pb.ReservationRequest{},
	}
	///PROVERI OVE POKAZIVACE
	for _, reservationRequest := range *pending_requests {
		current := mapToReservationRequestPb(&reservationRequest)
		response.ReservationRequest = append(response.ReservationRequest, current)
	}
	return response, nil
}

func (handler *ReservationHandler) DeleteReservationRequest(ctx context.Context, request *pb.GetByIdRequest) (*pb.ReservationRequestResponse, error) {
	id := request.Id
	log.Printf("Reservation request with id %s", id)
	err := handler.ReservationRequestService.Delete(id)
	if err != nil {
		return nil, err
	}
	reservationRequest, err2 := handler.ReservationRequestService.GetById(id)
	if err2 != nil {
		return nil, err
	}
	ReservationRequestPb := mapToReservationRequestPb(reservationRequest)
	response := &pb.ReservationRequestResponse{
		ReservationRequest: ReservationRequestPb,
	}
	return response, nil
}

func (handler *ReservationHandler) AcceptReservationRequest(ctx context.Context, request *pb.GetByIdRequest) (*pb.ReservationRequestResponse, error) {
	id := request.Id
	log.Printf("Reservation request with id %s", id)
	err := handler.ReservationRequestService.Accept(id)
	if err != nil {
		return nil, err
	}
	reservationRequest, err2 := handler.ReservationRequestService.GetById(id)
	if err2 != nil {
		return nil, err
	}
	ReservationRequestPb := mapToReservationRequestPb(reservationRequest)
	response := &pb.ReservationRequestResponse{
		ReservationRequest: ReservationRequestPb,
	}
	return response, nil
}

func (handler *ReservationHandler) CancelReservation(ctx context.Context, request *pb.GetByIdRequest) (*pb.ReservationResponse, error) {
	id := request.Id
	log.Printf("Reservation request with id %s", id)
	err := handler.ReservationService.Cancel(id, "user2")
	if err != nil {
		return nil, err
	}
	reservation, err2 := handler.ReservationService.GetById(id)
	if err2 != nil {
		return nil, err
	}
	ReservationPb := mapToReservationPb(reservation)
	response := &pb.ReservationResponse{
		Reservation: ReservationPb,
	}
	return response, nil
}

/*rpc DeleteReservationRequest(GetByIdRequest) returns(ReservationRequestResponse) {
        	option (google.api.http) = {
			put: "/reservationRequests/delete/{id}"
		};
        }

        rpc AcceptReservationRequest(GetByIdRequest) returns(ReservationRequestResponse) {
        	option (google.api.http) = {
			put: "/reservationRequests/accept/{id}"
		};
        }

        rpc CancelReservation(GetByIdRequest) returns(ReservationRequestResponse) {
        	option (google.api.http) = {
			put: "/reservation/cancel/{id}"
		};
        }*/
