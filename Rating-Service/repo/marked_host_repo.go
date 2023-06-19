package repo

import (
	"Rating-Service/domain"

	"gorm.io/gorm"
)

type MarkedHostRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *MarkedHostRepository) GetByUserId(id string) (*domain.MarkedHost, error) {
	markedHost := domain.MarkedHost{}
	dbResult := repo.DatabaseConnection.First(&markedHost, "hostId = ?", id)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &markedHost, nil
}

func (repo *MarkedHostRepository) Create(markedHost *domain.MarkedHost) error {
	dbResult := repo.DatabaseConnection.Create(markedHost)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
