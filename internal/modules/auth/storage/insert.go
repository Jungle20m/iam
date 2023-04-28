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

func (s *Storage) CreateUserAccount(ctx context.Context, ua model.UserAccount) error {
	db := s.getConnection(ctx)
	return db.Create(&ua).Error
}

func (s *Storage) CreateTokenWhileList(ctx context.Context, twl model.TokenWhiteList) error {
	db := s.getConnection(ctx)
	return db.Create(&twl).Error
}
