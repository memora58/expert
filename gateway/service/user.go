package service

import (
	"gateway/common/response"
	"gateway/common/util"
	service "gateway/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService struct{}

func (*UserService) Register(c *gin.Context, params *service.UserRequest) {
	userClient := c.Keys["user"].(service.UserServiceClient)
	userResp, err := userClient.UserRegister(c, params)
	util.PanicIfUserError(err)
	r := response.Response{
		Data:   userResp,
		Status: uint(userResp.Code),
		Msg:    response.GetMsg(uint(userResp.Code)),
	}

	c.JSON(http.StatusOK, r)
}

func (*UserService) Login(c *gin.Context, params *service.UserRequest) {
	userClient := c.Keys["user"].(service.UserServiceClient)
	userResp, err := userClient.UserLogin(c, params)
	util.PanicIfUserError(err)
	userId := userResp.UserDetail.UserID
	if userId == 0 {
		c.JSON(http.StatusNotFound, response.GetMsg(response.ERROR))
		return
	}
	token, err := util.GenerateToken(uint(userId))
	r := response.Response{
		Data:   response.TokenData{User: userResp.UserDetail, Token: token},
		Status: uint(userResp.Code),
		Msg:    response.GetMsg(uint(userResp.Code)),
	}
	c.JSON(http.StatusOK, r)
}
