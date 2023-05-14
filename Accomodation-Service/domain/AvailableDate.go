package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AvailableDate struct {
	Id              uuid.UUID   `json:"id"`
	AccommodationId string      `json:"accommodation" gorm:"not null"`
	StartDate       time.Time   `json:"startDate" gorm:"not null"`
	EndDate         time.Time   `json:"endDate" gorm:"not null"`
	Price           uint16      `json:"price" gorm:"not null"`
	PricingType     PricingType `json:"pricingType" gorm:"not null"`
}

func (availableDate *AvailableDate) BeforeCreate(scope *gorm.DB) error {
	availableDate.Id = uuid.New()
	return nil
}
