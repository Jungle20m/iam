package grpc

import (
	"context"

	"iam/common"
	"iam/internal/modules/iam/business"
	protoc2 "iam/internal/modules/iam/handler/grpc/protoc"
	"iam/internal/modules/iam/repository"
)

type grpcTransport struct {
	protoc2.UnimplementedAuthServer
	AppCtx common.IAppContext
}

func New(appCtx common.IAppContext) *grpcTransport {
	return &grpcTransport{
		AppCtx: appCtx,
	}
}

func (t *grpcTransport) Register(ctx context.Context, in *protoc2.RegisterRequest) (*protoc2.RegisterResponse, error) {

	st := repository.NewMysqlStorage(t.AppCtx.GetDB())
	biz := business.NewRegisterBusiness(t.AppCtx, st)

	err := biz.Register(ctx, in.GetClientId(), in.GetPhoneNumber(), in.GetPassword())
	if err != nil {
		return &protoc2.RegisterResponse{}, err
	}
	return &protoc2.RegisterResponse{Code: 200, Message: "success"}, nil
}

func (t *grpcTransport) VerifyRegister(ctx context.Context, in *protoc2.VerifyRegisterRequest) (*protoc2.VerifyRegisterResponse, error) {

	return &protoc2.VerifyRegisterResponse{
		Code:    0,
		Message: "",
	}, nil
}

func (t *grpcTransport) Login(ctx context.Context, in *protoc2.LoginRequest) (*protoc2.LoginResponse, error) {
	return &protoc2.LoginResponse{
		IdToken:      "",
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}
