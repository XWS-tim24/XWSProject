package repo

import (
	"Accomodation-Service/domain"
	"Accomodation-Service/dto"
	"strings"

	"gorm.io/gorm"
)

type AccommodationRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *AccommodationRepository) Create(accommodation *domain.Accommodation) (*domain.Accommodation, error) {
	dbResult := repo.DatabaseConnection.Create(accommodation)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return accommodation, nil
}

func (repo *AccommodationRepository) GetById(id string) (*domain.Accommodation, error) {
	accommodation := &domain.Accommodation{}
	dbResult := repo.DatabaseConnection.First(accommodation, "id = ?", id)
	if dbResult != nil {
		println("Gresdka pri get by id accommodation")
		return accommodation, dbResult.Error
	}
	return accommodation, nil
}

func (repo *AccommodationRepository) Search(searchDto *dto.AccommodationSearchDTO) *[]domain.Accommodation {
	var accommodations []domain.Accommodation
	city := strings.ToLower(searchDto.City)
	guestNum := searchDto.GuestNum
	startDate := searchDto.StartDate
	endDate := searchDto.EndDate
	repo.DatabaseConnection.Model(&domain.Accommodation{}).
		Joins("JOIN available_dates ON accommodations.id = available_dates.accommodation_id").
		Where(`lower(accommodations.city) like ? and 
	? BETWEEN accommodations.min_guests and accommodations.max_guests and 
	available_dates.start_date <= ? and 
	available_dates.end_date >= ?`, "%"+city+"%", guestNum, startDate, endDate).
		Find(&accommodations)

	return &accommodations
}

func (repo *AccommodationRepository) ExistsById(id string) bool {

	if err := repo.DatabaseConnection.First(&domain.Accommodation{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}
