package mapper

import (
	"Accomodation-reservation-Service/domain"
	pb "github.com/XWS-tim24/Common/common/proto/accommodation_reservation_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapToReservationRequestDTOPb(reservationRequest *domain.ReservationRequest) *pb.GetAllPendingForUserDTO {
	reservationRequestPb := &pb.GetAllPendingForUserDTO{
		Id:              reservationRequest.Id.String(),
		AccommodationId: reservationRequest.AccomodationId,
		UserId:          reservationRequest.UserId,
		StartDate:       timestamppb.New(reservationRequest.StartDate),
		EndDate:         timestamppb.New(reservationRequest.EndDate),
		NumberOfGuests:  uint32(reservationRequest.NumberOfGuests),
		Deleted:         reservationRequest.Deleted,
	}
	return reservationRequestPb
}

func MapToReservationDTOPb(reservation *domain.Reservation, reservationRequest *domain.ReservationRequest, accommodationName string) *pb.ReservationDTO {
	reservationRequestPb := &pb.ReservationDTO{
		Id:                reservation.Id.String(),
		RequestId:         reservation.RequestId,
		UserId:            reservationRequest.UserId,
		StartDate:         timestamppb.New(reservationRequest.StartDate),
		EndDate:           timestamppb.New(reservationRequest.EndDate),
		AccommodationName: accommodationName,
	}
	return reservationRequestPb
}
