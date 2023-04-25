package model

import "time"

type UserVerification struct {
	ID         int        `gorm:"column:id"`
	Token      string     `gorm:"column:token"`
	Expired    int64      `gorm:"column:expired"`
	CreateTime *time.Time `gorm:"column:create_time; autoCreateTime"`
	UpdateTime *time.Time `gorm:"column:update_time; autoUpdateTime"`
}

func (UserVerification) TableName() string {
	return "user_verification"
}
