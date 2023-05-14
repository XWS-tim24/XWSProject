package service

import (
	"Accomodation-reservation-Service/domain"
	"Accomodation-reservation-Service/repo"
	"fmt"
	"time"
)

type ReservationService struct {
	ReservationRepo        *repo.ReservationRepository
	ReservationRequestRepo *repo.ReservationRequestRepository
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

func (service *ReservationService) Cancel(id string, userId string) error {
	reservation, err := service.GetById(id)
	if err != nil {
		print("error during canceling reservation, reservation with id %s not found", id)
		return err
	}
	allowed, err2 := service.canCancel(reservation, userId)
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

func (service *ReservationService) canCancel(reservation *domain.Reservation, userId string) (bool, error) {
	//number_of_canceled := service.ReservationRepo.GetNumberOfCanceled(userId)
	reservationRequest, err := service.ReservationRequestRepo.GetById(reservation.RequestId)
	if err != nil {
		print("error during canceling reservation, reservation request id not found")
		return false, err
	}
	if reservationRequest.UserId != userId {
		return false, fmt.Errorf(fmt.Sprintf("Not allowed to cancel reservation for other users, reservation id %s", reservation.Id))
	}
	if reservationRequest.StartDate.Before(time.Now()) {
		return false, fmt.Errorf(fmt.Sprintf("Not allowed to cancel reservation previous dates, reservation id %s", reservation.Id))
	}
	return true, nil
}
