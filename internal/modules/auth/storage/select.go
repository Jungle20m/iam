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
			SELECT id, user_name, phone_number, email, password, password_salt, password_hash_algorithms, user_status, user_verification_id, registration_time, create_time, update_time
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

//func (s *Storage) GetUserVerificationByID(ctx context.Context, id int) (*model.UserVerification, error) {
//	db := s.getConnection(ctx)
//	var record model.UserVerification
//	sql := `
//			SELECT id, token, expired, create_time, update_time
//			FROM user_verification
//			WHERE id=?
//		   `
//	err := db.Raw(sql, id).First(&record).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, common.ErrRecordNotFound
//		}
//		return nil, err
//	}
//	return &record, nil
//}
//
//func (s *Storage) GetTWLByAccountIDForUpdate(ctx context.Context, userAccountID int) (*model.TokenWhiteList, error) {
//	db := s.getConnection(ctx)
//	var record model.TokenWhiteList
//	sql := `
//			SELECT id, user_account_id, access_token, refresh_token
//			FROM token_white_list
//			WHERE user_account_id=?
//			FOR UPDATE;
//		   `
//	err := db.Raw(sql, userAccountID).First(&record).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, common.ErrRecordNotFound
//		}
//		return nil, err
//	}
//	return &record, nil
//}
