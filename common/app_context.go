package common

import (
	"gorm.io/gorm"
	"iam/config"
)

type IAppContext interface {
	GetConfig() *config.Config
	GetDB() *gorm.DB
}

type appContext struct {
	Config *config.Config
	DB     *gorm.DB
}

func NewAppContext(conf *config.Config, db *gorm.DB) *appContext {
	return &appContext{
		Config: conf,
		DB:     db,
	}
}

func (ctx *appContext) GetConfig() *config.Config {
	return ctx.Config
}

func (ctx *appContext) GetDB() *gorm.DB {
	return ctx.DB
}
