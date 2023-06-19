package repo

import (
	"Rating-Service/domain"

	"gorm.io/gorm"
)

type HostRatingRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *HostRatingRepository) GetAll() (*[]domain.HostRating, error) {
	var rating []domain.HostRating
	result := repo.DatabaseConnection.Find(&rating)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rating, nil
}

func (repo *HostRatingRepository) GetById(id string) (domain.HostRating, error) {
	rating := domain.HostRating{}
	dbResult := repo.DatabaseConnection.First(&rating, "id = ?", id)
	if dbResult != nil {
		return rating, dbResult.Error
	}
	return rating, nil
}

func (repo *HostRatingRepository) Create(rating *domain.HostRating) error {
	dbResult := repo.DatabaseConnection.Create(rating)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
