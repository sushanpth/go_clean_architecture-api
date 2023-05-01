package models

import (
	"clean-architecture-api/lib"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessCategory struct {
	lib.ModelBase
	CategoryID string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
}

// BeforeCreate run this before creating user
func (t *BusinessCategory) BeforeCreate(_ *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = lib.BinaryUUID(id)
	return err
}
