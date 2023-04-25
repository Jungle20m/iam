package model

import "time"

// UserStatus
// - unverified
// - active
// - inactive

const (
	UserUnverifiedStatus string = "unverified"
	UserActiveStatus     string = "active"
	UserInactiveStatus   string = "inactive"
)

type UserAccount struct {
	ID                     int        `gorm:"column:id"`
	UserName               string     `gorm:"column:user_name"`
	PhoneNumber            string     `gorm:"column:phone_number"`
	Email                  string     `gorm:"column:email"`
	Password               string     `gorm:"column:password"`
	PasswordSalt           string     `gorm:"column:password_salt"`
	PasswordHashAlgorithms string     `gorm:"column:password_hash_algorithms"`
	UserStatus             string     `gorm:"column:user_status"`
	UserVerificationID     int        `gorm:"column:user_verification_id"`
	RegistrationTime       *time.Time `gorm:"column:registration_time"`
	CreateTime             *time.Time `gorm:"column:create_time; autoCreateTime"`
	UpdateTime             *time.Time `gorm:"column:update_time; autoUpdateTime"`
}

func (UserAccount) TableName() string {
	return "user_account"
}

type AuthorizedData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refreshToken"`
}
