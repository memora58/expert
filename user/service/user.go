package service

import (
	"context"
	"user/common/response"
	"user/model"
	service "user/proto"
)

type UserService struct{}

func (*UserService) UserLogin(c context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	resp = &service.UserDetailResponse{Code: response.SUCCESS}

	user := &model.User{}
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = response.ERROR
	}

	resp.UserDetail = model.BuildUser(*user)
	return resp, nil
}

func (*UserService) UserRegister(c context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	resp = &service.UserDetailResponse{Code: response.SUCCESS}

	user := &model.User{}
	err = user.Create(req)
	if err != nil {
		resp.Code = response.ERROR
	}
	resp.UserDetail = model.BuildUser(*user)
	return resp, nil
}

func (*UserService) UserLogout(c context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	resp = new(service.UserDetailResponse)
	resp.Code = response.SUCCESS
	return resp, nil
}
