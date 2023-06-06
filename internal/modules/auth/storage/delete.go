package storage

import (
	"context"
	"iam/internal/modules/auth/model"
)

func (s *Storage) DeleteUserToken(ctx context.Context, userID int, clientID string) error {
	db := s.getConnection(ctx)
	return db.Where("user_id = ? AND client_id = ?", userID, clientID).Delete(&model.UserToken{}).Error
}
