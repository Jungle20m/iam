package model

import "time"

type TokenWhiteList struct {
	ID           int        `gorm:"column:id"`
	AccessToken  string     `gorm:"column:access_token"`
	RefreshToken string     `gorm:"column:refresh_token"`
	CreateTime   *time.Time `gorm:"column:create_time; autoCreateTime"`
	UpdateTime   *time.Time `gorm:"column:update_time; autoUpdateTime"`
}

func (TokenWhiteList) TableName() string {
	return "token_white_list"
}
