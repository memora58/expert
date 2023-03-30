package controller

import (
	"context"
	"task/proto"
	"task/service"
)

type TaskController struct {
	proto.UnimplementedTaskServiceServer
}

var taskService service.TaskService

func (controller *TaskController) TaskCreate(ctx context.Context, request *proto.TaskRequest) (*proto.CommonResponse, error) {
	return taskService.TaskCreate(ctx, request)
}

func (controller *TaskController) TaskShow(ctx context.Context, request *proto.TaskRequest) (*proto.TasksDetailResponse, error) {
	return taskService.TaskShow(ctx, request)
}

func (controller *TaskController) TaskUpdate(ctx context.Context, request *proto.TaskRequest) (*proto.CommonResponse, error) {
	return taskService.TaskUpdate(ctx, request)
}

func (controller *TaskController) TaskDelete(ctx context.Context, request *proto.TaskRequest) (*proto.CommonResponse, error) {
	return taskService.TaskDelete(ctx, request)
}
