package business

import (
	"context"
	"iam/common"
)

type IUserStorage interface {
}

type userBusiness struct {
	appCtx  common.IAppContext
	storage IUserStorage
}

func NewUserBusiness(appCtx common.IAppContext, storage IUserStorage) *userBusiness {
	return &userBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func (biz *userBusiness) Register(ctx context.Context) error {
	// Kiểm tra user đó đã tồn tại hay chưa
	return nil
}
