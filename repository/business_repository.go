package repository

import (
	"clean-architecture-api/infrastructure"
	"clean-architecture-api/lib"

	"gorm.io/gorm"
)

type BusinessRepository struct {
	infrastructure.Database
	logger lib.Logger
}

func NewBusinessRepository(db infrastructure.Database, logger lib.Logger) BusinessRepository {
	return BusinessRepository{db, logger}
}

// WithTrx delegate transaction from user repository
func (r BusinessRepository) WithTrx(trxHandle *gorm.DB) BusinessRepository {
	if trxHandle != nil {
		r.logger.Debug("using WithTrx as trxHandle is not nil")
		r.Database.DB = trxHandle
	}
	return r
}
