package handler

import (
	"Accomodation-reservation-Service/domain"

	pb "github.com/XWS-tim24/Common/common/proto/accommodation_reservation_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapToReservationPb(reservation *domain.Reservation) *pb.Reservation {
	reservationPb := &pb.Reservation{
		Id:        reservation.Id.String(),
		RequestId: reservation.RequestId,
		Status:    mapToStatusPb(reservation.Status),
	}
	return reservationPb
}

func mapToReservationRequestPb(reservationRequest *domain.ReservationRequest) *pb.ReservationRequest {
	reservationRequestPb := &pb.ReservationRequest{
		Id:             reservationRequest.Id.String(),
		UserId:         reservationRequest.UserId,
		StartDate:      timestamppb.New(reservationRequest.StartDate),
		EndDate:        timestamppb.New(reservationRequest.EndDate),
		NumberOfGuests: uint32(reservationRequest.NumberOfGuests),
		Status:         mapToReservationRequestStatusPb(reservationRequest.Status),
		Deleted:        reservationRequest.Deleted,
	}
	return reservationRequestPb
}

func mapToReservation(reservationPb *pb.Reservation) *domain.Reservation {
	reservation := &domain.Reservation{
		RequestId: reservationPb.RequestId,
		Status:    mapToStatus(reservationPb.Status),
	}
	return reservation
}

func mapToReservationRequest(reservationRequestPb *pb.ReservationRequest) *domain.ReservationRequest {
	reservationRequest := &domain.ReservationRequest{
		UserId:         reservationRequestPb.UserId,
		StartDate:      reservationRequestPb.StartDate.AsTime(),
		EndDate:        reservationRequestPb.EndDate.AsTime(),
		NumberOfGuests: uint64(reservationRequestPb.NumberOfGuests),
		Status:         mapToReservationRequestStatus(reservationRequestPb.Status),
		Deleted:        reservationRequestPb.Deleted,
	}
	return reservationRequest
}

func mapToStatus(statusPb pb.Reservation_ReservationStatus) domain.ReservationStatus {
	switch statusPb {
	case pb.Reservation_Active:
		return domain.Active
	}
	return domain.Canceled
}

func mapToStatusPb(status domain.ReservationStatus) pb.Reservation_ReservationStatus {
	switch status {
	case domain.Active:
		return pb.Reservation_Active
	}
	return pb.Reservation_Canceled
}

func mapToReservationRequestStatus(status pb.ReservationRequest_ReservationRequestStatus) domain.ReservationRequestStatus {
	switch status {
	case pb.ReservationRequest_Pending:
		return domain.Pending
	case pb.ReservationRequest_Accepted:
		return domain.Accepted
	}
	return domain.Denied
}

func mapToReservationRequestStatusPb(status domain.ReservationRequestStatus) pb.ReservationRequest_ReservationRequestStatus {
	switch status {
	case domain.Pending:
		return pb.ReservationRequest_Pending
	case domain.Accepted:
		return pb.ReservationRequest_Accepted
	}
	return pb.ReservationRequest_Denied
}
