package business

import (
	"context"
	"fmt"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"iam/common"
	"iam/internal/modules/auth/model"
	"time"
)

const OtpPeriod int = 30

type IUserStorage interface {
	WithTx(ctx context.Context, fn func(c context.Context) error) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
	CreateUserVerification(ctx context.Context, uv model.UserVerification) (*model.UserVerification, error)
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
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
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
			return err
		}
		fmt.Printf("hashed password: %v\n", hashPassword)

		// Create OTP
		uv, err := biz.generateOTP(txContext, phoneNumber)
		if err != nil {
			return err
		}
		fmt.Printf("user_verification: %v\n", uv)

		//userAccount, err := biz.storage.GetUserByPhone(txContext, phoneNumber)
		//if err != nil {
		//	return err
		//}
		//
		//if userAccount.UserStatus != model.UserUnverifiedStatus {
		//	return fmt.Errorf("account has existed")
		//}

		return nil
	})
	return err
}
