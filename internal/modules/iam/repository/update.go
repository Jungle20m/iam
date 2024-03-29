package repository

import (
	"context"

	"iam/internal/modules/iam/model"

	tracersdk "iam/pkg/tracer"
)

func (s *Storage) UpdateUserAccount(ctx context.Context, ua model.UserAccount) error {
	ctx, span := tracersdk.NewSpan(ctx)
	defer span.End()

	db := s.getConnection(ctx)
	return db.Save(&ua).Error
}

func (s *Storage) UpdateUserToken(ctx context.Context, ut model.UserToken) error {
	db := s.getConnection(ctx)
	sql := `
			UPDATE user_token
			SET id_token=?, access_token=?, refresh_token=?
			WHERE user_id=? AND client_id=? 
		   `
	return db.Exec(sql, ut.IDToken, ut.AccessToken, ut.RefreshToken, ut.UserID, ut.ClientID).Error
}
