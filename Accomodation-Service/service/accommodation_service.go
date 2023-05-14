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

func (service *AccommodationService) Create(accommodation *domain.Accommodation) error {
	err := service.AccommodationRepository.Create(accommodation)
	if err != nil {
		return err
	}
	return nil
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
