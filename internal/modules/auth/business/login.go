package business

//
//import (
//	"context"
//	"fmt"
//	"iam/common"
//	"iam/internal/modules/auth/model"
//	mhttp "iam/sdk/httpserver"
//)
//
//type ILoginStorage interface {
//	WithTx(ctx context.Context, fn func(c context.Context) error) error
//	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
//	GetTWLByAccountIDForUpdate(ctx context.Context, userAccountID int) (*model.TokenWhiteList, error)
//	CreateTokenWhileList(ctx context.Context, twl model.TokenWhiteList) error
//	UpdateTWL(ctx context.Context, twl model.TokenWhiteList) error
//}
//
//type loginBusiness struct {
//	appCtx  common.IAppContext
//	storage ILoginStorage
//}
//
//func NewLoginBusiness(appCtx common.IAppContext, storage ILoginStorage) *loginBusiness {
//	return &loginBusiness{
//		appCtx:  appCtx,
//		storage: storage,
//	}
//}
//
//func (biz *loginBusiness) Login(ctx context.Context, phoneNumber, password string) (*model.AuthorizedData, error) {
//	// Get user by phone number
//	ua, err := biz.storage.GetUserByPhone(ctx, phoneNumber)
//	if err != nil {
//		if err == common.ErrRecordNotFound {
//			return nil, mhttp.BadRequestErrorResponse(fmt.Errorf("record not found"), "account not found", "ACCOUNT_NOT_FOUND")
//		}
//		return nil, err
//	}
//
//	switch ua.UserStatus {
//	case model.UserActiveStatus:
//
//	case model.UserInactiveStatus:
//		return nil, mhttp.BadRequestErrorResponse(fmt.Errorf("account invalid"), "this account has banned", "ACCOUNT_BANNED")
//	case model.UserUnverifiedStatus:
//		return nil, mhttp.BadRequestErrorResponse(fmt.Errorf("record not found"), "account not found", "ACCOUNT_NOT_FOUND")
//	default:
//		return nil, mhttp.InternalErrorResponse(fmt.Errorf("unreconized user status: %s", ua.UserStatus), "something went wrong", "UNRECOGNIZED_STATUS")
//	}
//
//	if ok := VerifyPassword(ua.Password, password); ok != true {
//		return nil, mhttp.BadRequestErrorResponse(fmt.Errorf("password invalid for user: %s", ua.PhoneNumber), "username or password wrong", "USERNAME_PASSWORD_WRONG")
//	}
//
//	accessToken, err := GenerateToken(true, ua.ID, 24, AccessSecretKey)
//	if err != nil {
//		return nil, err
//	}
//	refreshToken, err := GenerateToken(true, ua.ID, 24*30, RefreshSecretKey)
//	if err != nil {
//		return nil, err
//	}
//
//	err = biz.storage.WithTx(ctx, func(txContext context.Context) error {
//		_, err := biz.storage.GetTWLByAccountIDForUpdate(txContext, ua.ID)
//		if err != nil && err != common.ErrRecordNotFound {
//			return err
//		}
//
//		newTWL := model.TokenWhiteList{
//			UserAccountID: ua.ID,
//			AccessToken:   accessToken,
//			RefreshToken:  refreshToken,
//		}
//
//		if err == common.ErrRecordNotFound {
//			// insert
//			fmt.Println("inserted")
//			return biz.storage.CreateTokenWhileList(txContext, newTWL)
//		} else {
//			// update
//			fmt.Println("updated")
//			return biz.storage.UpdateTWL(txContext, newTWL)
//		}
//	})
//	if err != nil {
//		return nil, err
//
//	}
//
//	return &model.AuthorizedData{
//		AccessToken:  accessToken,
//		RefreshToken: refreshToken,
//	}, nil
//}
