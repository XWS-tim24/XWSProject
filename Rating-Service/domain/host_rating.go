package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HostRating struct {
	Id     uuid.UUID `json:"id"`
	HostId string    `json:"hostId" gorm:"not null"`
	UserId string    `json:"userId" gorm:"not null"`
	Value  uint32    `json:"value" gorm:"not null"`
	Date   time.Time `json:"date" gorm:"not null"`
}

func (hostRating *HostRating) BeforeCreate(scope *gorm.DB) error {
	hostRating.Id = uuid.New()
	return nil
}
