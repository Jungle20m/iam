package model

import "time"

type UserVerification struct {
	ID         int        `gorm:"column:id"`
	Token      string     `gorm:"column:token"`
	Period     int        `gorm:"column:period"`
	CreateTime *time.Time `gorm:"column:create_time; autoCreateTime"`
	UpdateTime *time.Time `gorm:"column:update_time; autoUpdateTime"`
}

func (UserVerification) TableName() string {
	return "user_verification"
}
