package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Accommodation struct {
	Id              uuid.UUID `json:"id"`
	UserID          string    `json:"userId" gorm:"not null"`
	Name            string    `json:"name" gorm:"unique;not null"`
	City            string    `json:"city" gorm:"not null"`
	Address         string    `json:"address" gorm:"not null"`
	Benefits        string    `json:"benefits" `
	MinGuests       uint32    `json:"minGuests" gorm:"not null"`
	MaxGuests       uint32    `json:"maxGuests" gorm:"not null"`
	AutomaticAccept bool      `json:"automaticAccept" gorm:"not null"`
}

func (accommodation *Accommodation) BeforeCreate(scope *gorm.DB) error {
	accommodation.Id = uuid.New()
	return nil
}
