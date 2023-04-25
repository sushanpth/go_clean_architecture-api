package repository

import (
	"clean-architecture-api/infrastructure"
	"clean-architecture-api/lib"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	infrastructure.Database
	logger lib.Logger
}

func NewCategoryRepository(db infrastructure.Database, logger lib.Logger) CategoryRepository {
	return CategoryRepository{db, logger}
}

// WithTrx delegate transaction from user repository
func (r CategoryRepository) WithTrx(trxHandle *gorm.DB) CategoryRepository {
	if trxHandle != nil {
		r.logger.Debug("using WithTrx as trxHandle is not nil")
		r.Database.DB = trxHandle
	}
	return r
}
