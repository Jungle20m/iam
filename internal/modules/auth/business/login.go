package business

import (
	"context"
	"iam/common"
	"iam/internal/modules/auth/model"
)

type ILoginStorage interface {
}

type loginBusiness struct {
	appCtx  common.IAppContext
	storage ILoginStorage
}

func NewLoginBusiness(appCtx common.IAppContext, storage ILoginStorage) *loginBusiness {
	return &loginBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, phoneNumber, password string) (*model.AuthorizedData, error) {

	return nil, nil
}
