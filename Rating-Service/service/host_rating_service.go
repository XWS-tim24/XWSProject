package service

import (
	"Rating-Service/domain"
	"Rating-Service/repo"
	"fmt"
)

type HostRatingService struct {
	HostRatingRepo *repo.HostRatingRepository
}

func (service *HostRatingService) GetAll() (*[]domain.HostRating, error) {
	return service.HostRatingRepo.GetAll()
}

func (service *HostRatingService) GetById(id string) (*domain.HostRating, error) {
	rating, err := service.HostRatingRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("host rating with id %s not found", id))
	}
	return &rating, nil
}

func (service *HostRatingService) Create(rating *domain.HostRating) error {
	err := service.HostRatingRepo.Create(rating)
	if err != nil {
		return err
	}
	return nil
}
