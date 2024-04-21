package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CommonFields ...
type CommonFields struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (c *CommonFields) BeforeCreate(_ *gorm.DB) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	c.ID = id

	return nil
}
