package business

import (
	"context"
	"fmt"
	"iam/common"
	"iam/internal/modules/auth/model"
	mhttp "iam/sdk/httpserver"
)

type ILoginStorage interface {
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
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
	// check if user has registered
	ua, err := biz.storage.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		if err == common.ErrRecordNotFound {
			return nil, mhttp.BadRequestErrorResponse(err, "this account may be hasn't registered", "ACCOUNT_NOT_FOUND")
		}
		return nil, mhttp.InternalErrorResponse(err, "something went wrong", "INTERNAL_SERVER_ERROR")
	}
	if ua.UserStatus == model.UserUnverifiedStatus {
		return nil, mhttp.BadRequestErrorResponse(err, "this account may be hasn't registered", "ACCOUNT_NOT_FOUND")
	}
	if ua.UserStatus == model.UserInactiveStatus {
		return nil, mhttp.BadRequestErrorResponse(err, "this account has banned", "ACCOUNT_INACTIVE")
	}

	fmt.Printf("verify password: %v\n", VerifyPassword(ua.Password, password))

	return nil, nil
}
