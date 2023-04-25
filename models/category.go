package models

import (
	"clean-architecture-api/lib"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	lib.ModelBase
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

// BeforeCreate run this before creating user
func (t *Category) BeforeCreate(_ *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = lib.BinaryUUID(id)
	return err
}
