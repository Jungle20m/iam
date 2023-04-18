package model

import "time"

type UserVerification struct {
	ID         int        `gorm:"column:id"`
	Token      string     `gorm:"column:token"`
	ExpireTime *time.Time `gorm:"column:expire_time"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
}

func (UserVerification) TableName() string {
	return "user_verification"
}
