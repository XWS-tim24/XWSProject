package repo

import (
	"Rating-Service/domain"

	"gorm.io/gorm"
)

type AccommodationRatingRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *AccommodationRatingRepository) GetAll() (*[]domain.AccommodationRating, error) {
	var rating []domain.AccommodationRating
	result := repo.DatabaseConnection.Find(&rating)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rating, nil
}

func (repo *AccommodationRatingRepository) GetById(id string) (domain.AccommodationRating, error) {
	rating := domain.AccommodationRating{}
	dbResult := repo.DatabaseConnection.First(&rating, "id = ?", id)
	if dbResult != nil {
		return rating, dbResult.Error
	}
	return rating, nil
}

func (repo *AccommodationRatingRepository) Create(rating *domain.AccommodationRating) error {
	dbResult := repo.DatabaseConnection.Create(rating)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
