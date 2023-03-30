package controller

import (
	"gateway/common/util"
	"gateway/proto"
	"gateway/service"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (*UserController) Register(c *gin.Context) {
	params := &proto.UserRequest{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		util.PanicIfUserError(err)
	}

	(&service.UserService{}).Register(c, params)
}

func (*UserController) Login(c *gin.Context) {
	params := &proto.UserRequest{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		util.PanicIfUserError(err)
	}

	(&service.UserService{}).Login(c, params)
}
