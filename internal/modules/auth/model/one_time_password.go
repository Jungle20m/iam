package model

import "time"

type OneTimePassword struct {
	ID          int        `gorm:"column:id"`
	UserID      int        `gorm:"column:user_id"`
	ClientID    string     `gorm:"column:client_id"`
	PhoneNumber string     `gorm:"column:phone_number"`
	OTP         string     `gorm:"column:token"`
	Expired     int64      `gorm:"column:expired"`
	MessageBody string     `gorm:"column:message_body"`
	CreateTime  *time.Time `gorm:"column:create_time; autoCreateTime"`
	UpdateTime  *time.Time `gorm:"column:update_time; autoUpdateTime"`
}

func (OneTimePassword) TableName() string {
	return "one_time_password"
}
