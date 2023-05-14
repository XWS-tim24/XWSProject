package service

import (
	"Accomodation-Service/communication"
	"Accomodation-Service/domain"
	"Accomodation-Service/dto"
	"Accomodation-Service/repo"
	"context"
	"fmt"
	pb "github.com/XWS-tim24/Common/common/proto/accommodation_reservation_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AvailableDateService struct {
	AvailableDateRepository   *repo.AvailableDateRepository
	AccommodationRepository   *repo.AccommodationRepository
	ReservationServiceAddress string
}

func (service *AvailableDateService) Create(availableDate *domain.AvailableDate) (*domain.AvailableDate, error) {
	if availableDate.StartDate.After(availableDate.EndDate) {
		return nil, fmt.Errorf("start date cannot be after end date")
	}
	if !service.AccommodationRepository.ExistsById(availableDate.AccommodationId) {
		return nil, fmt.Errorf(fmt.Sprintf("Accommodation with id %s not found", availableDate.AccommodationId))
	}
	if !service.AvailableDateRepository.TimeSlotFree(availableDate.AccommodationId, availableDate.StartDate, availableDate.EndDate) {
		return nil, fmt.Errorf("time slot already taken")
	}
	resp, err := service.AvailableDateRepository.Create(availableDate)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (service *AvailableDateService) GetById(id string) (*domain.AvailableDate, error) {
	availableDate, err := service.AvailableDateRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("reservation with id %s not found", id))
	}
	return availableDate, nil
}

func (service *AvailableDateService) Update(id string, availableDateDto *dto.AvailableDateDTO) error {
	availableDate, err := service.GetById(id)
	if err != nil {
		return err
	}
	reservationServiceClient := communication.NewReservationClient(service.ReservationServiceAddress)

	req := &pb.AlreadyReservedForDateRequest{}
	req.DateAndAccomodationDTO = &pb.DateAndAccomodationDTO{AccommodationId: availableDate.AccommodationId, StartDate: timestamppb.New(availableDateDto.StartDate), EndDate: timestamppb.New(availableDateDto.EndDate)}
	response, err1 := reservationServiceClient.AlreadyReservedForDate(context.TODO(), req)
	if err1 != nil {
		return err1
	}
	if response.AlreadyReserved {
		return fmt.Errorf("cannot update you have reservation")
	}
	availableDate.StartDate = availableDateDto.StartDate
	availableDate.EndDate = availableDateDto.EndDate
	availableDate.Price = availableDateDto.Price

	availableDate.PricingType = availableDateDto.PricingType

	err = service.AvailableDateRepository.Update(id, availableDate)
	if err != nil {
		print("error during deleting reservation request logicaly, request id %s", id)
		return err
	}
	return nil
}

func (service *AvailableDateService) TimeSlotAvailableForAccommodation(slotDTO *dto.AvailableTimeSlotDTO) (bool, error) {
	id := slotDTO.AccommodationId
	if !service.AccommodationRepository.ExistsById(id) {
		return false, fmt.Errorf(fmt.Sprintf("Accommodation with id %s not found", id))
	}
	timeSlotAvailable := service.AvailableDateRepository.TimeSlotAvailableForAccommodation(id, slotDTO.StartDate, slotDTO.EndDate)
	return timeSlotAvailable, nil
}
