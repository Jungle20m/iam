package repository

import (
	"context"

	"gorm.io/gorm"
)

const TransactionContextKey = "TransactionContextKey"

func (s *Storage) WithTx(ctx context.Context, fn func(txContext context.Context) error) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	c := context.WithValue(ctx, TransactionContextKey, tx)
	err := fn(c)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (s *Storage) GetTransactionFromCtx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(TransactionContextKey).(*gorm.DB)
	return tx, ok
}
