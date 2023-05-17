package dto

import (
	"Accomodation-Service/domain"
)

type AccommodationSearchDTOResponse struct {
	Id          string             `json:"id"`
	Name        string             `json:"name" `
	City        string             `json:"city" `
	Address     string             `json:"address" `
	Benefits    string             `json:"benefits" `
	MinGuests   uint32             `json:"minGuests" `
	MaxGuests   uint32             `json:"maxGuests" `
	Price       uint16             `json:"price" `
	TotalPrice  uint16             `json:"totalPrice"`
	PricingType domain.PricingType `json:"pricingType" `
}
