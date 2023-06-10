package storage

import (
	"context"
	"iam/internal/modules/auth/model"
	tracersdk "iam/sdk/tracer"
)

func (s *Storage) CreateOneTimePassword(ctx context.Context, otp *model.OneTimePassword) error {
	ctx, span := tracersdk.NewSpan(ctx)
	defer span.End()

	db := s.getConnection(ctx)
	return db.Create(&otp).Error
}

func (s *Storage) CreateUserAccount(ctx context.Context, ua *model.UserAccount) error {
	ctx, span := tracersdk.NewSpan(ctx)
	defer span.End()

	db := s.getConnection(ctx)
	return db.Create(&ua).Error
}

func (s *Storage) CreateUserToken(ctx context.Context, ut model.UserToken) error {
	db := s.getConnection(ctx)
	return db.Create(&ut).Error
}
