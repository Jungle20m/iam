package model

import "time"

type TokenWhileList struct {
	ID           int        `gorm:"column:id"`
	AccessToken  string     `gorm:"column:access_token"`
	RefreshToken string     `gorm:"column:refresh_token"`
	CreateTime   *time.Time `gorm:"column:create_time; autoCreateTime"`
	UpdateTime   *time.Time `gorm:"column:update_time; autoUpdateTime"`
}

func (TokenWhileList) TableName() string {
	return "token_while_list"
}
