package business

import (
	"context"
	"fmt"

	"iam/common"
	"iam/internal/modules/iam/model"
	mhttp "iam/pkg/httpserver"
)

type ILoginStorage interface {
	WithTx(ctx context.Context, fn func(c context.Context) error) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
	CreateUserToken(ctx context.Context, ut model.UserToken) error
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

type AuthorizedData struct {
	AccessToken string `json:"access_token"`
}

func (biz *loginBusiness) Login(ctx context.Context, phoneNumber, password string) (*AuthorizedData, error) {
	// Get user by phone number
	ua, err := biz.storage.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		if err == common.ErrRecordNotFound {
			return nil, mhttp.BadRequestErrorResponse(fmt.Errorf("record not found"), "account not found", "ACCOUNT_NOT_FOUND")
		}
		return nil, err
	}

	switch ua.UserStatus {
	case model.UserActiveStatus:

	case model.UserInactiveStatus:
		return nil, mhttp.BadRequestErrorResponse(fmt.Errorf("account invalid"), "this account has banned", "ACCOUNT_BANNED")
	case model.UserUnverifiedStatus:
		return nil, mhttp.BadRequestErrorResponse(fmt.Errorf("record not found"), "account not found", "ACCOUNT_NOT_FOUND")
	default:
		return nil, mhttp.InternalErrorResponse(fmt.Errorf("unreconized user status: %s", ua.UserStatus), "something went wrong")
	}

	if ok := VerifyPassword(ua.Password, password); ok != true {
		return nil, mhttp.BadRequestErrorResponse(fmt.Errorf("password invalid for user: %s", ua.PhoneNumber), "username or password wrong", "USERNAME_PASSWORD_WRONG")
	}

	accessToken, err := GenerateToken(true, ua.ID, ua.UserName, ua.Email, AccessSecretKey, AccessTokenExpired)
	if err != nil {
		return nil, err
	}

	err = biz.storage.WithTx(ctx, func(txContext context.Context) error {
		// Create new access token for user login
		ut := model.UserToken{
			UserID:      ua.ID,
			AccessToken: accessToken,
		}
		return biz.storage.CreateUserToken(txContext, ut)
	})
	if err != nil {
		return nil, err
	}

	return &AuthorizedData{
		AccessToken: accessToken,
	}, nil
}
