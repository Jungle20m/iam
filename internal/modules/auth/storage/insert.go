package storage

import (
	"context"
	"iam/internal/modules/auth/model"
)

func (s *Storage) CreateUserVerification(ctx context.Context, uv model.UserVerification) (*model.UserVerification, error) {
	db := s.getConnection(ctx)
	err := db.Create(&uv).Error
	return &uv, err
}
