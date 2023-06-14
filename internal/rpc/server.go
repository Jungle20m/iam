package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"iam/common"
	"iam/internal/modules/auth/business"
	"iam/internal/modules/auth/grpc-transport/protoc"
	"iam/internal/modules/auth/storage"
	"log"
	"net"
)

type server struct {
	protoc.UnimplementedAuthServer
	appCtx common.IAppContext
}

func NewServer(appCtx common.IAppContext) *server {
	s := grpc.NewServer()
	protoc.RegisterAuthServer(s, &server{})

	return &server{
		appCtx: appCtx,
	}
}

func (s *server) Serve(host string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpc := grpc.NewServer()

	protoc.RegisterAuthServer(rpc, s)

	reflection.Register(rpc)

	if err := rpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Register(ctx context.Context, in *protoc.RegisterRequest) (*protoc.RegisterReply, error) {

	st := storage.NewMysqlStorage(s.appCtx.GetDB())
	biz := business.NewRegisterBusiness(s.appCtx, st)

	err := biz.Register(ctx, in.GetClientId(), in.GetPhoneNumber(), in.GetPassword())
	if err != nil {
		return &protoc.RegisterReply{}, err
	}
	return &protoc.RegisterReply{Code: 200, Message: "success"}, nil
}
