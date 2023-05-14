package service

import (
	"Accomodation-Service/domain"
	"Accomodation-Service/dto"
	"Accomodation-Service/repo"
	"fmt"
)

type AvailableDateService struct {
	AvailableDateRepository *repo.AvailableDateRepository
	AccommodationRepository *repo.AccommodationRepository
}

func (service *AvailableDateService) Create(availableDate *domain.AvailableDate) error {
	if availableDate.StartDate.After(availableDate.EndDate) {
		return fmt.Errorf("start date cannot be after end date")
	}
	if !service.AccommodationRepository.ExistsById(availableDate.AccommodationId) {
		return fmt.Errorf(fmt.Sprintf("Accommodation with id %s not found", availableDate.AccommodationId))
	}
	if !service.AvailableDateRepository.TimeSlotFree(availableDate.AccommodationId, availableDate.StartDate, availableDate.EndDate) {
		return fmt.Errorf("time slot already taken")
	}
	err := service.AvailableDateRepository.Create(availableDate)
	if err != nil {
		return err
	}
	return nil
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
	availableDate.StartDate = availableDateDto.StartDate
	availableDate.EndDate = availableDateDto.EndDate
	availableDate.Price = availableDateDto.Price

	availableDate.PricingType = availableDateDto.PricingType

	service.AvailableDateRepository.Update(id, availableDate)
	if err != nil {
		print("error during deleting reservation request logicaly, request id %s", id)
		return err
	}
	return nil
}
