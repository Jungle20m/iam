package storage

import (
	"context"
	"gorm.io/gorm"
)

type Storage struct {
	db    *gorm.DB
	other string
}

func NewMysqlStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) getConnection(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(TransactionContextKey).(*gorm.DB)
	if ok {
		return tx
	}
	return s.db
}
