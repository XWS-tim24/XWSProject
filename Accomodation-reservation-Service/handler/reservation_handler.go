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
	println("accomodation id in handler %s", reservationRequest.AccomodationId)
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
	pendingRequestsDtos, err := handler.ReservationRequestService.GetAllPendingForUser(userId)
	if err != nil {
		return &pb.GetAllPendingForUserResponse{}, err
	}

	response := &pb.GetAllPendingForUserResponse{
		ReservationRequest: []*pb.GetAllPendingForUserDTO{},
	}
	///PROVERI OVE POKAZIVACE
	for _, reservationRequestDto := range *pendingRequestsDtos {
		response.ReservationRequest = append(response.ReservationRequest, &reservationRequestDto)
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

func (handler *ReservationHandler) DenyReservationRequest(ctx context.Context, request *pb.GetByIdRequest) (*pb.ReservationRequestResponse, error) {
	id := request.Id
	log.Printf("Reservation request with id %s", id)
	err := handler.ReservationRequestService.Deny(id)
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

func (handler *ReservationHandler) CancelReservation(ctx context.Context, request *pb.GetByIdAndUserIdRequest) (*pb.ReservationResponse, error) {
	id := request.Id
	userId := request.UserIdDto.UserId
	log.Printf("Canceling request with id %s and userId %s ", id, userId)
	err := handler.ReservationService.Cancel(id, userId)
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

func (handler *ReservationHandler) AlreadyReservedForDate(ctx context.Context, request *pb.AlreadyReservedForDateRequest) (*pb.AlreadyReservedForDateResponse, error) {
	print("Stigao handler")
	alreadyReserved := handler.ReservationRequestService.AlreadyReservedForDate(request.DateAndAccomodationDTO.AccommodationId, request.DateAndAccomodationDTO.StartDate.AsTime(), request.DateAndAccomodationDTO.EndDate.AsTime())

	response := &pb.AlreadyReservedForDateResponse{
		AlreadyReserved: alreadyReserved,
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
func (handler *ReservationHandler) GetAllAcceptedReservationsForUser(ctx context.Context, request *pb.GetAllAcceptedReservationsForUserRequest) (*pb.GetAllReservationsResponse, error) {
	reservations, err := handler.ReservationService.GetAllAcceptedReservationsForUser(request.UserId)
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
