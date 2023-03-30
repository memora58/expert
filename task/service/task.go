package service

import (
	"context"
	"task/common/response"
	"task/model"
	service "task/proto"
)

type TaskService struct{}

func (*TaskService) TaskCreate(_ context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	resp = &service.CommonResponse{Code: response.SUCCESS}

	err = (&model.Task{}).Create(req)
	if err != nil {
		resp.Code = response.ERROR
		resp.Data = err.Error()
	}

	resp.Msg = response.GetMsg(uint(resp.Code))
	return resp, nil
}

func (*TaskService) TaskShow(_ context.Context, req *service.TaskRequest) (resp *service.TasksDetailResponse, err error) {
	resp = &service.TasksDetailResponse{Code: response.SUCCESS}

	taskResp, err := (&model.Task{}).Show(req)
	if err != nil {
		resp.Code = response.ERROR
	}

	resp.TaskDetail = model.BuildTasks(taskResp)
	return resp, nil
}

func (*TaskService) TaskUpdate(_ context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	resp = &service.CommonResponse{Code: response.SUCCESS}

	err = (&model.Task{}).Update(req)
	if err != nil {
		resp.Code = response.ERROR
		resp.Data = err.Error()
	}

	resp.Msg = response.GetMsg(uint(resp.Code))
	return resp, nil
}

func (*TaskService) TaskDelete(_ context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	resp = &service.CommonResponse{Code: response.SUCCESS}

	err = (&model.Task{}).Delete(req)
	if err != nil {
		resp.Code = response.ERROR
		resp.Data = err.Error()
	}

	resp.Msg = response.GetMsg(uint(resp.Code))
	return resp, nil
}
