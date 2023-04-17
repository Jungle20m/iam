package model

import "time"

type User struct {
	ID         int        `gorm:"column:id"`
	Name       string     `gorm:"column:name"`
	Phone      string     `gorm:"column:phone"`
	Email      string     `gorm:"column:email"`
	Password   string     `gorm:"column:password"`
	IsActive   int        `gorm:"column:is_active"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
}

type RegisterBody struct {
	Phone    string `json:"phone"`
	PassWord string `json:"password"`
}
