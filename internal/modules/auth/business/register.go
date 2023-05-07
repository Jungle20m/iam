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

const (
	AccessSecretKey  = "ACCESS_SECRET_KEY"
	RefreshSecretKey = "REFRESH_SECRET_KEY"
)

type IRegisterStorage interface {
	WithTx(ctx context.Context, fn func(c context.Context) error) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
	CreateUserVerification(ctx context.Context, uv model.UserVerification) (*model.UserVerification, error)
	CreateUserAccount(ctx context.Context, ua model.UserAccount) error
	UpdateUserAccount(ctx context.Context, ua model.UserAccount) error
	GetUserVerificationByID(ctx context.Context, id int) (*model.UserVerification, error)
	CreateTokenWhileList(ctx context.Context, twl model.TokenWhiteList) error
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

func (biz *registerBusiness) generateOTP(ctx context.Context, phoneNumber string) (*model.UserVerification, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "IAM",
		AccountName: phoneNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("error generating OTP key: %v", err)
	}
	otp, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return nil, fmt.Errorf("error generating TOTP code: %v", err)
	}
	uv := model.UserVerification{
		Token:   otp,
		Expired: time.Now().Add(time.Second * time.Duration(OtpPeriod)).Unix(),
	}
	return biz.storage.CreateUserVerification(ctx, uv)
}

func (biz *registerBusiness) Register(ctx context.Context, phoneNumber, password string) error {
	err := biz.storage.WithTx(ctx, func(txContext context.Context) error {
		// Generate hash password
		hashPassword, err := GenerateHashPassword(password)
		if err != nil {
			return mhttp.InternalErrorResponse(
				fmt.Errorf("error to generate hash password: %v", err),
				"Sorry, you cannot register at this time",
				"INTERNAL_SERVER_ERROR")
		}

		// Create OTP
		uv, err := biz.generateOTP(txContext, phoneNumber)
		if err != nil {
			return mhttp.InternalErrorResponse(
				fmt.Errorf("error to create OTP: %v", err),
				"Sorry, you cannot register at this time",
				"INTERNAL_SERVER_ERROR")
		}

		ua, err := biz.storage.GetUserByPhone(txContext, phoneNumber)

		// If the user has never registered before, create new
		if err == common.ErrRecordNotFound {
			newUA := model.UserAccount{
				PhoneNumber:        phoneNumber,
				Password:           hashPassword,
				UserStatus:         model.UserUnverifiedStatus,
				UserVerificationID: uv.ID,
			}
			return biz.storage.CreateUserAccount(txContext, newUA)
		}
		// If an error occurs
		if err != nil {
			return err
		}
		// If user has registered
		if ua.UserStatus != model.UserUnverifiedStatus {
			return mhttp.BadRequestErrorResponse(nil, "account already exists", "ACCOUNT_ALREADY_EXISTS")
		}
		// If the user has registered but has not verified
		// generate new otp and hash password
		ua.UserVerificationID = uv.ID
		ua.Password = hashPassword
		return biz.storage.UpdateUserAccount(txContext, *ua)
	})
	return err
}

func (biz *registerBusiness) VerifyRegister(ctx context.Context, phoneNumber, otp string) (*model.AuthorizedData, error) {
	// Get user and verify otp
	ua, err := biz.storage.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		if err == common.ErrRecordNotFound {
			return nil, fmt.Errorf("record not found")
		}
		return nil, err
	}
	// Check if user has existed?
	if ua.UserStatus != model.UserUnverifiedStatus {
		return nil, mhttp.BadRequestErrorResponse(nil, "user has existed", "USER_HAS_EXISTED")
	}

	uv, err := biz.storage.GetUserVerificationByID(ctx, ua.UserVerificationID)
	if err != nil {
		return nil, err
	}
	if uv.Token != otp || uv.Expired < time.Now().Unix() {
		return nil, mhttp.BadRequestErrorResponse(nil, "incorrect otp or otp has expired", "OTP_INVALID")
	}

	accessToken, err := GenerateToken(true, ua.ID, 24, AccessSecretKey)
	if err != nil {
		return nil, err
	}
	refreshToken, err := GenerateToken(true, ua.ID, 24*30, RefreshSecretKey)
	if err != nil {
		return nil, err
	}

	// Save token and activate user
	err = biz.storage.WithTx(ctx, func(txContext context.Context) error {
		// Activate user
		now := time.Now()
		ua.UserStatus = model.UserActiveStatus
		ua.RegistrationTime = &now
		if err := biz.storage.UpdateUserAccount(txContext, *ua); err != nil {
			return err
		}
		// Add token to while list
		twl := model.TokenWhiteList{
			UserAccountID: ua.ID,
			AccessToken:   accessToken,
			RefreshToken:  refreshToken,
		}
		if err := biz.storage.CreateTokenWhileList(txContext, twl); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &model.AuthorizedData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
