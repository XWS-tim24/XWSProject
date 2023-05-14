package dto

import (
	"time"
)

type AvailableTimeSlotDTO struct {
	AccommodationId string    `json:"accommodationId"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate" `
}
