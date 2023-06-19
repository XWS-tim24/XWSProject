package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccommodationRating struct {
	Id              uuid.UUID `json:"id"`
	AccommodationId string    `json:"accommodationId" gorm:"not null"`
	UserId          string    `json:"userId" gorm:"not null"`
	Value           uint32    `json:"value" gorm:"not null"`
	Date            time.Time `json:"date" gorm:"not null"`
}

func (accommodationRating *AccommodationRating) BeforeCreate(scope *gorm.DB) error {
	accommodationRating.Id = uuid.New()
	return nil
}
