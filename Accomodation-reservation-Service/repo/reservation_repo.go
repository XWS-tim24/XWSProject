package repo

import (
	"Accomodation-reservation-Service/domain"
	"fmt"

	"gorm.io/gorm"
)

type ReservationRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *ReservationRepository) GetAll() (*[]domain.Reservation, error) {
	var reservations []domain.Reservation
	result := repo.DatabaseConnection.Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reservations, nil
}

func (repo *ReservationRepository) GetById(id string) (domain.Reservation, error) {
	reservation := domain.Reservation{}
	dbResult := repo.DatabaseConnection.First(&reservation, "id = ?", id)
	if dbResult != nil {
		return reservation, dbResult.Error
	}
	return reservation, nil
}

func (repo *ReservationRepository) Create(reservation *domain.Reservation) error {
	dbResult := repo.DatabaseConnection.Create(reservation)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *ReservationRepository) Cancel(id string) error {
	dbResult := repo.DatabaseConnection.Model(&domain.Reservation{}).Where("id = ? AND status = ?", id, domain.Active).Update("status", domain.Canceled)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	fmt.Println(dbResult.RowsAffected)
	return nil
}

// ZA FRONT
func (repo *ReservationRepository) GetNumberOfCanceled(userId string) int64 {
	var number_of_canceled int64
	//query := `SELECT * FROM reservation_requests WHERE status = 0 AND user_id = ?`
	query := `SELECT COUNT(*) FROM reservations r, reservation_requests rr
		WHERE r.request_id = rr.id 
		AND rr.user_id = ? 
		AND r.status = 1`
	stmt := repo.DatabaseConnection.Raw(query, userId)
	stmt.Scan(&number_of_canceled)
	return number_of_canceled
}
func (repo *ReservationRepository) GetAllAcceptedReservationsForUser(userId string) (*[]domain.Reservation, error) {
	var reservations []domain.Reservation
	result := repo.DatabaseConnection.Model(&domain.Reservation{}).
		Joins("JOIN reservation_requests ON reservations.request_id = reservation_requests.id").
		Where(`reservation_requests.user_id = ? and reservations.status = ?`, userId, domain.Active).
		Find(&reservations)

	if result.Error != nil {
		return nil, result.Error
	}
	return &reservations, nil
}

func (repo *ReservationRepository) GetAllAcceptedReservationsForAccommodation(accId string) (*[]domain.Reservation, error) {
	var reservations []domain.Reservation
	result := repo.DatabaseConnection.Model(&domain.Reservation{}).
		Joins("JOIN reservation_requests ON reservations.request_id = reservation_requests.id").
		Where(`reservation_requests.accomodation_id = ? reservations.status = ?`, accId, domain.Active).
		Find(&reservations)

	if result.Error != nil {
		return nil, result.Error
	}
	return &reservations, nil
}

// POMOCNE
/*
func (repo *ReservationRepository) GetRequestForReservationId(reservationId string) (domain.ReservationRequest, error) {
	reservationRequest := domain.ReservationRequest{}
	dbResult := repo.DatabaseConnection.First(&reservationRequest, "id = ?", reservationId)
	if dbResult != nil {
		return reservationRequest, dbResult.Error
	}
	return reservationRequest, nil
}*/
