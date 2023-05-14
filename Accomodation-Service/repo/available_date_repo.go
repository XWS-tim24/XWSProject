package repo

import (
	"Accomodation-Service/domain"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AvailableDateRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *AvailableDateRepository) Create(availableDate *domain.AvailableDate) error {
	dbResult := repo.DatabaseConnection.Create(availableDate)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *AvailableDateRepository) GetById(id string) (*domain.AvailableDate, error) {
	availableDate := &domain.AvailableDate{}
	dbResult := repo.DatabaseConnection.First(availableDate, "id = ?", id)
	if dbResult.Error != nil {
		println("error in available date repository")
		return availableDate, dbResult.Error
	}
	return availableDate, nil
}

func (repo *AvailableDateRepository) Update(id string, availableDate *domain.AvailableDate) error {

	var updateData map[string]interface{}
	updateData = map[string]interface{}{
		"start_date":   availableDate.StartDate,
		"end_date":     availableDate.EndDate,
		"price":        availableDate.Price,
		"pricing_type": availableDate.PricingType,
	}
	result := repo.DatabaseConnection.Model(&domain.AvailableDate{}).Where("id = ?", id).Updates(&updateData)

	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo *AvailableDateRepository) TimeSlotFree(accommodationId string, startTime time.Time, endTime time.Time) bool {
	result := repo.DatabaseConnection.First(&domain.AvailableDate{}, "accommodation_id = ? and ? < end_date and ? > start_date", accommodationId, startTime, endTime)
	return result.Error != nil
}
