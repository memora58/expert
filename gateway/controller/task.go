package controller

import (
	"gateway/common/util"
	"gateway/proto"
	"gateway/service"
	"github.com/gin-gonic/gin"
)

type TaskController struct{}

func (*TaskController) GetTaskList(c *gin.Context) {
	(&service.TaskService{}).GetTaskList(c)
}

func (*TaskController) CreateTask(c *gin.Context) {
	params := &proto.TaskRequest{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		util.PanicIfUserError(err)
	}
	(&service.TaskService{}).CreateTask(c, params)
}
