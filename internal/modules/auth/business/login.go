package business

import (
	"context"
	"fmt"
	"iam/common"
	"iam/internal/modules/auth/model"
	mhttp "iam/sdk/httpserver"
)

type ILoginStorage interface {
	WithTx(ctx context.Context, fn func(c context.Context) error) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
	GetUserTokenForUpdate(ctx context.Context, userID int, clientID string) (*model.UserToken, error)
	CreateUserToken(ctx context.Context, ut model.UserToken) error
	UpdateUserToken(ctx context.Context, ut model.UserToken) error
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

func (biz *loginBusiness) Login(ctx context.Context, clientID, phoneNumber, password string) (*model.AuthorizedData, error) {
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
	refreshToken, err := GenerateToken(true, ua.ID, ua.UserName, ua.Email, RefreshSecretKey, RefreshTokenExpired)
	if err != nil {
		return nil, err
	}
	idToken, err := GenerateToken(true, ua.ID, ua.UserName, ua.Email, IdTokenSecretKey, IdTokenExpired)
	if err != nil {
		return nil, err
	}

	err = biz.storage.WithTx(ctx, func(txContext context.Context) error {
		_, err := biz.storage.GetUserTokenForUpdate(txContext, ua.ID, clientID)
		if err != nil && err != common.ErrRecordNotFound {
			return err
		}

		ut := model.UserToken{
			UserID:       ua.ID,
			ClientID:     clientID,
			IDToken:      idToken,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		if err == common.ErrRecordNotFound {
			return biz.storage.CreateUserToken(txContext, ut)
		} else {
			return biz.storage.UpdateUserToken(txContext, ut)
		}
	})
	if err != nil {
		return nil, err

	}

	return &model.AuthorizedData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IdToken:      idToken,
	}, nil
}
