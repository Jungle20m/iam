package business

import (
	"context"
	"iam/common"
)

type ILogoutStorage interface {
	DeleteToken(ctx context.Context, userID int) error
}

type logoutBusiness struct {
	appCtx  common.IAppContext
	storage ILogoutStorage
}

func NewLogoutBusiness(appCtx common.IAppContext, storage ILogoutStorage) *logoutBusiness {
	return &logoutBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func (biz *logoutBusiness) Logout(ctx context.Context, userID int) error {
	return biz.storage.DeleteToken(ctx, userID)
}
