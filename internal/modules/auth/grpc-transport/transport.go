package grpc_transport

import (
	"context"
	"iam/common"
	"iam/internal/modules/auth/business"
	"iam/internal/modules/auth/grpc-transport/protoc"
	"iam/internal/modules/auth/storage"
)

type grpcTransport struct {
	protoc.UnimplementedAuthServer
	AppCtx common.IAppContext
}

func New(appCtx common.IAppContext) *grpcTransport {
	return &grpcTransport{
		AppCtx: appCtx,
	}
}

func (t *grpcTransport) Register(ctx context.Context, in *protoc.RegisterRequest) (*protoc.RegisterResponse, error) {

	st := storage.NewMysqlStorage(t.AppCtx.GetDB())
	biz := business.NewRegisterBusiness(t.AppCtx, st)

	err := biz.Register(ctx, in.GetClientId(), in.GetPhoneNumber(), in.GetPassword())
	if err != nil {
		return &protoc.RegisterResponse{}, err
	}
	return &protoc.RegisterResponse{Code: 200, Message: "success"}, nil
}

func (t *grpcTransport) VerifyRegister(ctx context.Context, in *protoc.VerifyRegisterRequest) (*protoc.VerifyRegisterResponse, error) {

	return &protoc.VerifyRegisterResponse{
		Code:    0,
		Message: "",
	}, nil
}

func (t *grpcTransport) Login(ctx context.Context, in *protoc.LoginRequest) (*protoc.LoginResponse, error) {
	return &protoc.LoginResponse{
		IdToken:      "",
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}
