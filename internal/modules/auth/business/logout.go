package business

import (
	"context"
	"iam/common"
)

type ILogoutStorage interface {
	DeleteUserToken(ctx context.Context, userID int, clientID string) error
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

func (biz *logoutBusiness) Logout(ctx context.Context, userID int, clientID string) error {
	return biz.storage.DeleteUserToken(ctx, userID, clientID)
}
