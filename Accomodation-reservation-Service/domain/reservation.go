package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reservation struct {
	Id        uuid.UUID `gorm:"index:idx_name,unique"`
	RequestId string
	Status    ReservationStatus //ili bool
}

// enum
type ReservationStatus int8

const (
	Active ReservationStatus = iota
	Canceled
)

func (status ReservationStatus) String() string {
	switch status {
	case Active:
		return "Active"
	case Canceled:
		return "Canceled"
	}
	return "Unknown"
}

// id
func (reservation *Reservation) BeforeCreate(scope *gorm.DB) error {
	reservation.Id = uuid.New()
	return nil
}
