package model

import "time"

type UserAccount struct {
	ID                     int        `gorm:"column:id"`
	UserName               string     `gorm:"column:name"`
	PhoneNumber            string     `gorm:"column:phone"`
	Email                  string     `gorm:"column:email"`
	Password               string     `gorm:"column:password"`
	PasswordSalt           string     `gorm:"column:password_salt"`
	PasswordHashAlgorithms string     `gorm:"column:password_hash_algorithms"`
	UserStatus             string     `gorm:"column:user_status"`
	UserVerificationID     int        `gorm:"column:user_verification_id"`
	RegistrationTime       *time.Time `gorm:"column:registration_time"`
	CreateTime             *time.Time `gorm:"column:create_time"`
	UpdateTime             *time.Time `gorm:"column:update_time"`
}

func (UserAccount) TableName() string {
	return "user_account"
}
