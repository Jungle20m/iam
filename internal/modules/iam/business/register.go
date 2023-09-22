package business

import (
	"context"
	"errors"
	"fmt"

	"iam/common"
	"iam/internal/modules/iam/model"
	httpsdk "iam/pkg/httpserver"
)

const OtpPeriod int = 30

type IRegisterStorage interface {
	WithTx(ctx context.Context, fn func(c context.Context) error) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
	CreateOneTimePassword(ctx context.Context, otp *model.OneTimePassword) error
	CreateUserAccount(ctx context.Context, ua *model.UserAccount) error
	UpdateUserAccount(ctx context.Context, ua model.UserAccount) error
	GetLastOneTimePasswordByUserID(ctx context.Context, userID int, clientID string) (*model.OneTimePassword, error)
	CreateUserToken(ctx context.Context, ut model.UserToken) error
}

type registerBusiness struct {
	appCtx  common.IAppContext
	storage IRegisterStorage
}

func NewRegisterBusiness(appCtx common.IAppContext, storage IRegisterStorage) *registerBusiness {
	return &registerBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

// Register is func to registration account
func (biz *registerBusiness) Register(ctx context.Context, phoneNumber, password string) error {
	// ctx, span := tracersdk.NewSpan(ctx)
	// defer span.End()

	// Generate hash password
	hashPassword, err := GenerateHashPassword(password)
	if err != nil {
		return httpsdk.InternalErrorResponse(
			fmt.Errorf("error to generate hash password: %v", err),
			"Sorry, you cannot register at this time")
	}

	userAccount, err := biz.storage.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		// If we faced database error
		if err != common.ErrRecordNotFound {
			return httpsdk.InternalErrorResponse(err, "")
		}
		// If this the first time user registration, create new user account
		newUserAccount := model.UserAccount{
			PhoneNumber: phoneNumber,
			Password:    hashPassword,
			UserStatus:  model.UserActiveStatus,
		}
		// Create new user account
		if err := biz.storage.CreateUserAccount(ctx, &newUserAccount); err != nil {
			return httpsdk.InternalErrorResponse(err, "")
		}
		return nil
	}

	// Check if account exist
	if userAccount != nil {
		return errors.New("account has existed")
	}

	return nil
}

// // VerifyRegister is func to verify otp
// func (biz *registerBusiness) VerifyRegister(ctx context.Context, clientID, phoneNumber, otpCode string) (*model.AuthorizedData, error) {
// 	// Verify user, check if user has existed?
// 	ua, err := biz.storage.GetUserByPhone(ctx, phoneNumber)
// 	if err != nil {
// 		if err == common.ErrRecordNotFound {
// 			return nil, fmt.Errorf("record not found")
// 		}
// 		return nil, err
// 	}
// 	if ua.UserStatus != model.UserUnverifiedStatus {
// 		return nil, httpsdk.BadRequestErrorResponse(nil, "user has existed", "USER_HAS_EXISTED")
// 	}
//
// 	// Validate the last otp
// 	otp, err := biz.storage.GetLastOneTimePasswordByUserID(ctx, ua.ID, clientID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if otp.OTP != otpCode || otp.Expired < time.Now().Unix() {
// 		return nil, httpsdk.BadRequestErrorResponse(nil, "incorrect otp or otp has expired", "OTP_INVALID")
// 	}
//
// 	accessToken, err := GenerateToken(true, ua.ID, ua.UserName, ua.Email, AccessSecretKey, AccessTokenExpired)
// 	if err != nil {
// 		return nil, err
// 	}
// 	refreshToken, err := GenerateToken(true, ua.ID, ua.UserName, ua.Email, RefreshSecretKey, RefreshTokenExpired)
// 	if err != nil {
// 		return nil, err
// 	}
// 	idToken, err := GenerateToken(true, ua.ID, ua.UserName, ua.Email, IdTokenSecretKey, IdTokenExpired)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Save token and activate user
// 	err = biz.storage.WithTx(ctx, func(txContext context.Context) error {
// 		// Activate user
// 		now := time.Now()
// 		ua.UserStatus = model.UserActiveStatus
// 		ua.RegistrationTime = &now
// 		if err := biz.storage.UpdateUserAccount(txContext, *ua); err != nil {
// 			return err
// 		}
// 		// Add token to while list
// 		ut := model.UserToken{
// 			UserID:       ua.ID,
// 			ClientID:     clientID,
// 			IDToken:      idToken,
// 			AccessToken:  accessToken,
// 			RefreshToken: refreshToken,
// 		}
// 		if err := biz.storage.CreateUserToken(txContext, ut); err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &model.AuthorizedData{
// 		IdToken:      idToken,
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}, nil
// }
