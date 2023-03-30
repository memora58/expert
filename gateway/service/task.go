package service

import (
	"gateway/common/response"
	"gateway/common/util"
	service "gateway/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskService struct{}

func (*TaskService) GetTaskList(c *gin.Context) {
	var tReq service.TaskRequest

	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserID)
	taskClient := c.Keys["task"].(service.TaskServiceClient)

	TaskResp, err := taskClient.TaskShow(c, &tReq)
	PanicIfTaskError(err)

	r := response.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    response.GetMsg(uint(TaskResp.Code)),
	}

	c.JSON(http.StatusOK, r)
}

func (*TaskService) CreateTask(c *gin.Context, params *service.TaskRequest) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	params.UserID = uint32(claim.UserID)

	taskClient := c.Keys["task"].(service.TaskServiceClient)

	TaskResp, err := taskClient.TaskCreate(c, params)
	PanicIfTaskError(err)
	r := response.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    response.GetMsg(uint(TaskResp.Code)),
	}
	c.JSON(http.StatusOK, r)
}
