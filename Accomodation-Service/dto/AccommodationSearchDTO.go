package dto

import "time"

type AccommodationSearchDTO struct {
	City      string    `json:"city"`
	GuestNum  uint16    `json:"guestNum"`
	StartDate time.Time `json:"startDate" `
	EndDate   time.Time `json:"endDate"`
}
