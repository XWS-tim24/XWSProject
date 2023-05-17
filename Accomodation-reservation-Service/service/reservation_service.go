package service

import (
	"Accomodation-reservation-Service/communication"
	"Accomodation-reservation-Service/domain"
	"Accomodation-reservation-Service/mapper"
	"Accomodation-reservation-Service/repo"
	"context"
	"fmt"
	pbReservations "github.com/XWS-tim24/Common/common/proto/accommodation_reservation_service"
	pb "github.com/XWS-tim24/Common/common/proto/accommodation_service"
	"time"
)

type ReservationService struct {
	ReservationRepo            *repo.ReservationRepository
	ReservationRequestRepo     *repo.ReservationRequestRepository
	AccommodationServiceAddres string
}

func (service *ReservationService) GetAll() (*[]domain.Reservation, error) {
	return service.ReservationRepo.GetAll()
}

func (service *ReservationService) GetById(id string) (*domain.Reservation, error) {
	reservation, err := service.ReservationRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("reservation with id %s not found", id))
	}
	return &reservation, nil
}

func (service *ReservationService) Create(reservation *domain.Reservation) error {
	err := service.ReservationRepo.Create(reservation)
	if err != nil {
		return err
	}
	return nil
}

func (service *ReservationService) Cancel(id string) error {
	reservation, err := service.GetById(id)
	if err != nil {
		print("error during canceling reservation, reservation with id %s not found", id)
		return err
	}
	allowed, err2 := service.canCancel(reservation)
	if !allowed {
		return err2
	}
	err = service.ReservationRepo.Cancel(id)
	if err != nil {
		print("error during canceling reservation, reservation id %s", id)
		return err
	}
	return nil
}

// ZA FRONT
func (service *ReservationService) GetNumberOfCanceled(userId string) int64 {
	number_of_canceled := service.ReservationRepo.GetNumberOfCanceled(userId)
	return number_of_canceled
}

func (service *ReservationService) canCancel(reservation *domain.Reservation) (bool, error) {
	//number_of_canceled := service.ReservationRepo.GetNumberOfCanceled(userId)
	reservationRequest, err := service.ReservationRequestRepo.GetById(reservation.RequestId)
	if err != nil {
		print("error during canceling reservation, reservation request id not found")
		return false, err
	}

	if reservationRequest.StartDate.Before(time.Now()) {
		return false, fmt.Errorf(fmt.Sprintf("Not allowed to cancel reservation previous dates, reservation id %s", reservation.Id))
	}
	return true, nil
}

func (service *ReservationService) GetAllAcceptedReservationsForUser(userId string) ([]*pbReservations.ReservationDTO, error) {
	reservations, err := service.ReservationRepo.GetAllAcceptedReservationsForUser(userId)
	if err != nil {
		return nil, err
	}
	reservationDtos := []*pbReservations.ReservationDTO{}
	accommodationClient := communication.NewAccommodationClient(service.AccommodationServiceAddres)
	for _, reserv := range *reservations {
		req, err := service.ReservationRequestRepo.GetById(reserv.RequestId)
		if err != nil {
			return nil, err
		}
		pbRequest := pb.GetByIdRequest{Id: req.AccomodationId}
		acc, err1 := accommodationClient.GetAccommodationById(context.TODO(), &pbRequest)
		if err1 != nil {
			return nil, err1
		}

		reservationDTO := mapper.MapToReservationDTOPb(&reserv, &req, acc.Accommodation.Name)
		reservationDtos = append(reservationDtos, reservationDTO)
	}

	return reservationDtos, nil
}

func (service *ReservationService) GetAllAcceptedReservationsForAllAccommodations(accId string) ([]*pbReservations.ReservationDTO, error) {
	reservations, err := service.ReservationRepo.GetAllAcceptedReservationsForAllAccommodations(accId)
	if err != nil {
		return nil, err
	}
	reservationDtos := []*pbReservations.ReservationDTO{}
	accommodationClient := communication.NewAccommodationClient(service.AccommodationServiceAddres)
	for _, reserv := range *reservations {
		req, err := service.ReservationRequestRepo.GetById(reserv.RequestId)
		if err != nil {
			return nil, err
		}
		pbRequest := pb.GetByIdRequest{Id: req.AccomodationId}
		acc, err1 := accommodationClient.GetAccommodationById(context.TODO(), &pbRequest)
		if err1 != nil {
			return nil, err1
		}

		reservationDTO := mapper.MapToReservationDTOPb(&reserv, &req, acc.Accommodation.Name)
		reservationDtos = append(reservationDtos, reservationDTO)
	}

	return reservationDtos, nil
}
