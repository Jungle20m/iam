package business

import (
	"context"
	"fmt"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"iam/common"
	"iam/internal/modules/auth/model"
	mhttp "iam/sdk/httpserver"
	"time"
)

const OtpPeriod int = 30

type IUserStorage interface {
	WithTx(ctx context.Context, fn func(c context.Context) error) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
	CreateUserVerification(ctx context.Context, uv model.UserVerification) (*model.UserVerification, error)
	CreateUserAccount(ctx context.Context, ua model.UserAccount) error
	UpdateUserAccount(ctx context.Context, ua model.UserAccount) error
}

type userBusiness struct {
	appCtx  common.IAppContext
	storage IUserStorage
}

func NewUserBusiness(appCtx common.IAppContext, storage IUserStorage) *userBusiness {
	return &userBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func generateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (biz *userBusiness) generateOTP(ctx context.Context, phoneNumber string) (*model.UserVerification, error) {
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
		Token:  otp,
		Period: OtpPeriod,
	}
	return biz.storage.CreateUserVerification(ctx, uv)
}

func (biz *userBusiness) Register(ctx context.Context, phoneNumber, password string) error {
	err := biz.storage.WithTx(ctx, func(txContext context.Context) error {
		// Generate hash password
		hashPassword, err := generateHashPassword(password)
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

func Add(a, b int) int {
	return a + b
}
