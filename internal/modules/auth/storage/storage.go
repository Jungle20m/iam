package storage

import (
	"context"
	"fmt"
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
		fmt.Println("connection with transaction")
		return tx
	}
	fmt.Println("connection without transaction")
	return s.db
}
