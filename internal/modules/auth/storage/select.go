package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"iam/common"
	"iam/internal/modules/auth/model"
)

func (s *Storage) GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error) {
	db := s.getConnection(ctx)
	var record model.UserAccount
	sql := `
			SELECT id, user_name, phone_number, email, password, password_salt, password_hash_algorithms, user_status, registration_time, create_time, update_time
			FROM user_account
			WHERE phone_number=?
   		   `
	err := db.Raw(sql, phoneNumber).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return &record, nil
}

func (s *Storage) GetLastOneTimePasswordByUserID(ctx context.Context, userID int, clientID string) (*model.OneTimePassword, error) {
	db := s.getConnection(ctx)
	var record model.OneTimePassword
	sql := `
			SELECT id, user_id, client_id, phone_number, otp, expired, message_body
			FROM one_time_password
			WHERE user_id=? AND client_id=?
			ORDER BY id DESC
			LIMIT 1
		   `
	err := db.Raw(sql, userID, clientID).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return &record, nil
}

func (s *Storage) GetUserTokenForUpdate(ctx context.Context, userID int, clientID string) (*model.UserToken, error) {
	db := s.getConnection(ctx)

	var record model.UserToken

	sql := `
			SELECT id, user_id, client_id, id_token, access_token, refresh_token
			FROM user_token
			WHERE user_id=? AND client_id=?
			FOR UPDATE;
		   `
	err := db.Raw(sql, userID, clientID).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return &record, nil
}
