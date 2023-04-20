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
