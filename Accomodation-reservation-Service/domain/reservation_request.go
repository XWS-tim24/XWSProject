package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationRequest struct {
	Id             uuid.UUID `gorm:"index:idx_name,unique"`
	UserId         string
	AccomodationId string
	StartDate      time.Time
	EndDate        time.Time
	NumberOfGuests uint64
	Status         ReservationRequestStatus
	Deleted        bool
}

// enum
type ReservationRequestStatus int8

const (
	Pending ReservationRequestStatus = iota
	Accepted
	Denied
)

func (status ReservationRequestStatus) String() string {
	switch status {
	case Pending:
		return "Pending"
	case Accepted:
		return "Accepted"
	case Denied:
		return "Denied"
	}
	return "Unknown"
}

// id
func (reservationRequest *ReservationRequest) BeforeCreate(scope *gorm.DB) error {
	reservationRequest.Id = uuid.New()
	return nil
}
