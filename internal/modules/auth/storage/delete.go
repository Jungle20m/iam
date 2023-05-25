package storage

import (
	"context"
	"iam/internal/modules/auth/model"
)

func (s *Storage) DeleteToken(ctx context.Context, userID int) error {
	db := s.getConnection(ctx)
	return db.Where("user_account_id = ?", userID).Delete(&model.TokenWhiteList{}).Error
}
