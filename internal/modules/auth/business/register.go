package business

import (
	"context"
	"fmt"
	"github.com/pquerna/otp/totp"
	"iam/common"
	"iam/internal/modules/auth/model"
	mhttp "iam/sdk/httpserver"
	"time"
)

const OtpPeriod int = 30

type IRegisterStorage interface {
	WithTx(ctx context.Context, fn func(c context.Context) error) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
	CreateOneTimePassword(ctx context.Context, uv model.OneTimePassword) (*model.OneTimePassword, error)
	CreateUserAccount(ctx context.Context, ua model.UserAccount) error
	UpdateUserAccount(ctx context.Context, ua model.UserAccount) error
	//GetUserVerificationByID(ctx context.Context, id int) (*model.UserVerification, error)
	//CreateTokenWhileList(ctx context.Context, twl model.TokenWhiteList) error
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

func (biz *registerBusiness) generateOTP(ctx context.Context, clientID, phoneNumber string, userID int) (*model.OneTimePassword, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "IAM",
		AccountName: phoneNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("error generating OTP key: %v", err)
	}
	otpCode, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return nil, fmt.Errorf("error generating TOTP code: %v", err)
	}
	otp := model.OneTimePassword{
		UserID:      userID,
		ClientID:    clientID,
		PhoneNumber: phoneNumber,
		OTP:         otpCode,
		Expired:     time.Now().Add(time.Second * time.Duration(OtpPeriod)).Unix(),
		MessageBody: "",
	}
	return biz.storage.CreateOneTimePassword(ctx, otp)
}

func (biz *registerBusiness) Register(ctx context.Context, clientID, phoneNumber, password string) error {
	err := biz.storage.WithTx(ctx, func(txContext context.Context) error {
		// Generate hash password
		hashPassword, err := GenerateHashPassword(password)
		if err != nil {
			return mhttp.InternalErrorResponse(
				fmt.Errorf("error to generate hash password: %v", err),
				"Sorry, you cannot register at this time")
		}

		var userID int

		// Create or update user account
		ua, err := biz.storage.GetUserByPhone(txContext, phoneNumber)
		if err != nil {
			// If we faced database error
			if err != common.ErrRecordNotFound {
				return mhttp.InternalErrorResponse(err, "")
			}
			// If this the first time user registration, create new user account
			newUA := model.UserAccount{
				PhoneNumber: phoneNumber,
				Password:    hashPassword,
				UserStatus:  model.UserUnverifiedStatus,
			}
			// Create new user account
			if err := biz.storage.CreateUserAccount(txContext, newUA); err != nil {
				return mhttp.InternalErrorResponse(err, "")
			}
			userID = newUA.ID
		} else {
			// If account has existed
			if ua.UserStatus != model.UserUnverifiedStatus {
				return mhttp.BadRequestErrorResponse(nil, "account already exists", "ACCOUNT_ALREADY_EXISTS")
			}
			// If user has registered but hasn't verified
			ua.Password = hashPassword
			if err := biz.storage.UpdateUserAccount(txContext, *ua); err != nil {
				return mhttp.InternalErrorResponse(err, "")
			}
			userID = ua.ID
		}

		// Create OTP
		_, err = biz.generateOTP(txContext, clientID, phoneNumber, userID)
		if err != nil {
			return mhttp.InternalErrorResponse(err, "Sorry, you cannot register at this time")
		}
		return nil
	})
	return err
}

//func (biz *registerBusiness) VerifyRegister(ctx context.Context, phoneNumber, otp string) (*model.AuthorizedData, error) {
//	// Get user and verify otp
//	ua, err := biz.storage.GetUserByPhone(ctx, phoneNumber)
//	if err != nil {
//		if err == common.ErrRecordNotFound {
//			return nil, fmt.Errorf("record not found")
//		}
//		return nil, err
//	}
//	// Check if user has existed?
//	if ua.UserStatus != model.UserUnverifiedStatus {
//		return nil, mhttp.BadRequestErrorResponse(nil, "user has existed", "USER_HAS_EXISTED")
//	}
//
//	uv, err := biz.storage.GetUserVerificationByID(ctx, ua.UserVerificationID)
//	if err != nil {
//		return nil, err
//	}
//	if uv.Token != otp || uv.Expired < time.Now().Unix() {
//		return nil, mhttp.BadRequestErrorResponse(nil, "incorrect otp or otp has expired", "OTP_INVALID")
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
//	// Save token and activate user
//	err = biz.storage.WithTx(ctx, func(txContext context.Context) error {
//		// Activate user
//		now := time.Now()
//		ua.UserStatus = model.UserActiveStatus
//		ua.RegistrationTime = &now
//		if err := biz.storage.UpdateUserAccount(txContext, *ua); err != nil {
//			return err
//		}
//		// Add token to while list
//		twl := model.TokenWhiteList{
//			UserAccountID: ua.ID,
//			AccessToken:   accessToken,
//			RefreshToken:  refreshToken,
//		}
//		if err := biz.storage.CreateTokenWhileList(txContext, twl); err != nil {
//			return err
//		}
//		return nil
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &model.AuthorizedData{
//		AccessToken:  accessToken,
//		RefreshToken: refreshToken,
//	}, nil
//}
