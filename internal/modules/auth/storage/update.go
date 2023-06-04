package storage

import (
	"context"
	"iam/internal/modules/auth/model"
)

func (s *Storage) UpdateUserAccount(ctx context.Context, ua model.UserAccount) error {
	db := s.getConnection(ctx)
	return db.Save(&ua).Error
}

//func (s *Storage) UpdateTWL(ctx context.Context, twl model.TokenWhiteList) error {
//	db := s.getConnection(ctx)
//	sql := `
//			UPDATE token_white_list
//			SET access_token=?, refresh_token=?
//			WHERE user_account_id=?
//		   `
//	return db.Exec(sql, twl.AccessToken, twl.RefreshToken, twl.UserAccountID).Error
//}
