package business

import (
	"context"
	"fmt"
	"iam/common"
	"iam/internal/modules/auth/model"
	mhttp "iam/sdk/httpserver"
	"time"
)

type IPasswordStorage interface {
	GetUserByPhone(ctx context.Context, phoneNumber string) (*model.UserAccount, error)
	CreateOneTimePassword(ctx context.Context, otp *model.OneTimePassword) error
	GetLastOneTimePasswordByUserID(ctx context.Context, userID int, clientID string) (*model.OneTimePassword, error)
	UpdateUserAccount(ctx context.Context, ua model.UserAccount) error
}

type passwordBusiness struct {
	appCtx  common.IAppContext
	storage IPasswordStorage
}

func NewPasswordBusiness(appCtx common.IAppContext, storage IPasswordStorage) *passwordBusiness {
	return &passwordBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func (biz *passwordBusiness) Recover(ctx context.Context, clientID, phoneNumber string) error {
	// Check if account is available
	ua, err := biz.storage.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		if err == common.ErrRecordNotFound {
			return fmt.Errorf("record not found")
		}
		return err
	}
	if ua.UserStatus != model.UserActiveStatus {
		return mhttp.BadRequestErrorResponse(nil, "user not found", "")
	}

	// Generate OTP
	otp, err := GenerateOTP(ctx, clientID, phoneNumber, ua.ID)
	if err != nil {
		return mhttp.InternalErrorResponse(err, "")
	}
	return biz.storage.CreateOneTimePassword(ctx, otp)
}

func (biz *passwordBusiness) Verify(ctx context.Context, clientID, phoneNumber, password, otpCode string) error {
	// Check if account is available
	ua, err := biz.storage.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		if err == common.ErrRecordNotFound {
			return fmt.Errorf("record not found")
		}
		return err
	}
	if ua.UserStatus != model.UserActiveStatus {
		return mhttp.BadRequestErrorResponse(nil, "user not found", "")
	}

	// Validate the last otp
	otp, err := biz.storage.GetLastOneTimePasswordByUserID(ctx, ua.ID, clientID)
	if err != nil {
		return err
	}
	if otp.OTP != otpCode || otp.Expired < time.Now().Unix() {
		return mhttp.BadRequestErrorResponse(nil, "incorrect otp or otp has expired", "OTP_INVALID")
	}

	// Generate new password
	newPassword, err := GenerateHashPassword(password)
	if err != nil {
		return err
	}
	ua.Password = newPassword

	fmt.Println(newPassword)

	return biz.storage.UpdateUserAccount(ctx, *ua)
}
