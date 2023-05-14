package service

import (
	"Accomodation-Service/domain"
	"Accomodation-Service/dto"
	"Accomodation-Service/repo"
	"fmt"
)

type AccommodationService struct {
	AccommodationRepository *repo.AccommodationRepository
}

func (service *AccommodationService) Create(accommodation *domain.Accommodation) (*domain.Accommodation, error) {
	acc, err := service.AccommodationRepository.Create(accommodation)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (service *AccommodationService) GetById(id string) (*domain.Accommodation, error) {
	accommodation, err := service.AccommodationRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("reservation with id %s not found", id))
	}
	return accommodation, nil
}

func (service *AccommodationService) Search(accommodationSearchDTO *dto.AccommodationSearchDTO) *[]domain.Accommodation {
	return service.AccommodationRepository.Search(accommodationSearchDTO)

}

func (service *AccommodationService) GetAutomaticAcceptById(id string) (bool, error) {
	accommodation, err := service.AccommodationRepository.GetById(id)
	if err != nil {
		return false, fmt.Errorf(fmt.Sprintf("reservation with id %s not found", id))
	}
	return accommodation.AutomaticAccept, nil
}
