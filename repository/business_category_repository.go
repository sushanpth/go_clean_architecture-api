package repository

import (
	"clean-architecture-api/infrastructure"
	"clean-architecture-api/lib"

	"gorm.io/gorm"
)

type BusinessCategoryRepository struct {
	infrastructure.Database
	logger lib.Logger
}

func NewBusinessCategoryRepository(db infrastructure.Database, logger lib.Logger) BusinessCategoryRepository {
	return BusinessCategoryRepository{db, logger}
}

// WithTrx delegate transaction from user repository
func (r BusinessCategoryRepository) WithTrx(trxHandle *gorm.DB) BusinessCategoryRepository {
	if trxHandle != nil {
		r.logger.Debug("using WithTrx as trxHandle is not nil")
		r.Database.DB = trxHandle
	}
	return r
}
