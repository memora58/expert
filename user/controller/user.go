package controller

import (
	"context"
	"user/proto"
	"user/service"
)

type UserController struct {
	proto.UnimplementedUserServiceServer
}

var services service.UserService

func (UserController) UserLogin(ctx context.Context, req *proto.UserRequest) (*proto.UserDetailResponse, error) {
	return services.UserLogin(ctx, req)
}

func (UserController) UserLogout(ctx context.Context, req *proto.UserRequest) (*proto.UserDetailResponse, error) {
	return services.UserLogout(ctx, req)
}

func (UserController) UserRegister(ctx context.Context, req *proto.UserRequest) (*proto.UserDetailResponse, error) {
	return services.UserRegister(ctx, req)
}
