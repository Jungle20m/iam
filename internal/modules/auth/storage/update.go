package storage

import (
	"context"
	"iam/internal/modules/auth/model"
)

func (s *Storage) UpdateUserAccount(ctx context.Context, ua model.UserAccount) error {
	db := s.getConnection(ctx)
	return db.Save(&ua).Error
}
