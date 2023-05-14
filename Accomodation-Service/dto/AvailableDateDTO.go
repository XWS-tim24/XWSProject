package dto

import (
	"time"

	"Accomodation-Service/domain"
)

type AvailableDateDTO struct {
	Id          string             `json:"id"`
	StartDate   time.Time          `json:"startDate"`
	EndDate     time.Time          `json:"endDate" `
	Price       uint16             `json:"price" `
	PricingType domain.PricingType `json:"pricingType" `
}
