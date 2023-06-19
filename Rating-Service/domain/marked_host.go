package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MarkedHost struct {
	Id     uuid.UUID `json:"id"`
	HostId string    `json:"hostId" gorm:"not null"`
}

func (markedHost *MarkedHost) BeforeCreate(scope *gorm.DB) error {
	markedHost.Id = uuid.New()
	return nil
}
