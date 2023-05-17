package service

import (
	"Accomodation-Service/domain"
	"Accomodation-Service/dto"
	"Accomodation-Service/mapper"
	"Accomodation-Service/repo"
	"fmt"
)

type AccommodationService struct {
	AccommodationRepository *repo.AccommodationRepository
	AvailableRepository     *repo.AvailableDateRepository
}

func (service *AccommodationService) Create(accommodation *domain.Accommodation) (*domain.Accommodation, error) {
	acc, err := service.AccommodationRepository.Create(accommodation)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
func (service *AccommodationService) GetAll() (*[]domain.Accommodation, error) {
	return service.AccommodationRepository.GetAll()
}
func (service *AccommodationService) GetById(id string) (*domain.Accommodation, error) {
	accommodation, err := service.AccommodationRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("reservation with id %s not found", id))
	}
	return accommodation, nil
}

func (service *AccommodationService) Search(accommodationSearchDTO *dto.AccommodationSearchDTO) (*[]dto.AccommodationSearchDTOResponse, error) {
	accommodations := service.AccommodationRepository.Search(accommodationSearchDTO)
	searchResponses := []dto.AccommodationSearchDTOResponse{}
	numberOfDays := uint16(accommodationSearchDTO.EndDate.Sub(accommodationSearchDTO.StartDate).Hours())
	for _, acc := range *accommodations {
		availableDate, err := service.AvailableRepository.GetAvailableDateForAccommodationAndTimeSlot(acc.Id.String(), accommodationSearchDTO.StartDate, accommodationSearchDTO.EndDate)
		if err != nil {
			return nil, err
		}
		searchResponse := &dto.AccommodationSearchDTOResponse{}
		if availableDate.PricingType == domain.PER_GUEST {
			searchResponse = mapper.MapToSearchResponse(&acc, availableDate.Price, numberOfDays*accommodationSearchDTO.GuestNum*availableDate.Price, domain.PER_GUEST)

		} else {
			searchResponse = mapper.MapToSearchResponse(&acc, availableDate.Price, numberOfDays*availableDate.Price, domain.PER_ACCOMMODATION)

		}

		searchResponses = append(searchResponses, *searchResponse)
	}
	return &searchResponses, nil

}

func (service *AccommodationService) GetAutomaticAcceptById(id string) (bool, error) {
	accommodation, err := service.AccommodationRepository.GetById(id)
	if err != nil {
		return false, fmt.Errorf(fmt.Sprintf("reservation with id %s not found", id))
	}
	return accommodation.AutomaticAccept, nil
}
