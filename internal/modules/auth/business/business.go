package business

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"iam/common"
	"iam/internal/modules/auth/model"
)

type IUserStorage interface {
	WithTx(ctx context.Context, fn func(c context.Context) error) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
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

func generatePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (biz *userBusiness) Register(ctx context.Context, phoneNumber, password string) error {
	err := biz.storage.WithTx(ctx, func(txContext context.Context) error {
		userAccount, err := biz.storage.GetUserByPhone(txContext, phoneNumber)
		if err != nil {
			return err
		}

		if userAccount.UserStatus != model.UserUnverifiedStatus {
			return fmt.Errorf("account has existed")
		}

		// generate hash password
		hashPassword, err := generatePassword(password)
		if err != nil {
			return err
		}
		fmt.Printf("hashed password: %v\n", hashPassword)

		return nil
	})
	return err
}
