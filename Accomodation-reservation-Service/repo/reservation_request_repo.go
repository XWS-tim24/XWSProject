package repo

import (
	"Accomodation-reservation-Service/domain"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ReservationRequestRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *ReservationRequestRepository) GetById(id string) (domain.ReservationRequest, error) {
	reservationRequest := domain.ReservationRequest{}
	dbResult := repo.DatabaseConnection.First(&reservationRequest, "id = ?", id)
	if dbResult != nil {
		return reservationRequest, dbResult.Error
	}
	return reservationRequest, nil
}

func (repo *ReservationRequestRepository) Create(reservationRequest *domain.ReservationRequest) error {
	dbResult := repo.DatabaseConnection.Create(reservationRequest)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *ReservationRequestRepository) AcceptOrDeny(reservationRequest *domain.ReservationRequest) error {
	dbResult := repo.DatabaseConnection.Model(&domain.ReservationRequest{}).Where("id = ?", reservationRequest.Id).Update("status", reservationRequest.Status)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	fmt.Println(dbResult.RowsAffected)
	return nil
}

func (repo *ReservationRequestRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Model(&domain.ReservationRequest{}).Where("id = ? AND status = ?", id, domain.Pending).Update("deleted", true)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	fmt.Println(dbResult.RowsAffected)
	return nil
}

// ZA FRONT
func (repo *ReservationRequestRepository) GetAllPendingForUser(userId string) *[]domain.ReservationRequest {
	var pending_requests []domain.ReservationRequest
	query := `SELECT * FROM reservation_requests WHERE status = ? AND user_id = ?`
	stmt := repo.DatabaseConnection.Raw(query, domain.Pending, userId)
	stmt.Scan(&pending_requests)
	return &pending_requests
}

func (repo *ReservationRequestRepository) GetAllPendingForAccomodation(accomodationId string) *[]domain.ReservationRequest {
	var pending_requests []domain.ReservationRequest
	query := `SELECT * FROM reservation_requests WHERE status = ? AND accomodation_id = ?`
	stmt := repo.DatabaseConnection.Raw(query, domain.Pending, accomodationId)
	stmt.Scan(&pending_requests)
	return &pending_requests
}

// POMOCNE
func (repo *ReservationRequestRepository) AlreadyReservedForDate(accomodationId string, startDate time.Time, endDate time.Time) bool {
	var exists bool
	existsQuery := `SELECT EXISTS(
		SELECT 1 FROM reservations r, reservation_requests rr WHERE r.request_id = rr.id
		AND r.status = ?
		AND rr.accomodation_id = ?
		AND (? BETWEEN rr.start_date AND rr.end_date
		OR ? BETWEEN rr.start_date AND rr.end_date
		OR ? < rr.start_date AND ? > rr.end_date)
	 )`
	existsStmt := repo.DatabaseConnection.Raw(existsQuery,
		domain.Active,
		accomodationId,
		startDate,
		endDate,
		startDate, endDate)
	existsStmt.Scan(&exists)
	return exists
}

///////////////////////////////

func (repo *ReservationRequestRepository) DenyOthers(reservationRequest *domain.ReservationRequest) error {
	dbResult := repo.DatabaseConnection.Model(&domain.ReservationRequest{}).Where("status = ? AND accomodation_id = ? AND start_date < ? AND end_date > ?", domain.Pending, reservationRequest.AccomodationId, reservationRequest.EndDate, reservationRequest.StartDate).Update("status", domain.Denied)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	fmt.Println(dbResult.RowsAffected)
	return nil
}
