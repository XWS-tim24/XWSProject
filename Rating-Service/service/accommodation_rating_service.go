package service

import (
	"Rating-Service/domain"
	"Rating-Service/repo"
	"fmt"
)

type AccommodationRatingService struct {
	AccommodationRatingRepo *repo.AccommodationRatingRepository
}

func (service *AccommodationRatingService) GetAll() (*[]domain.AccommodationRating, error) {
	return service.AccommodationRatingRepo.GetAll()
}

func (service *AccommodationRatingService) GetById(id string) (*domain.AccommodationRating, error) {
	rating, err := service.AccommodationRatingRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("host rating with id %s not found", id))
	}
	return &rating, nil
}

func (service *AccommodationRatingService) Create(rating *domain.HostRating) error {
	err := service.AccommodationRatingRepo.Create(rating)
	if err != nil {
		return err
	}
	return nil
}
