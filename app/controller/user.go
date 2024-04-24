package controller

import (
	"yugod-backend/app/lib/response"
	"yugod-backend/app/openapi"
	"yugod-backend/app/server"
	"yugod-backend/app/util"

	"github.com/gin-gonic/gin"
)

type uUserApi interface {
	CreateUser(c *gin.Context)
	GetUserInfo(c *gin.Context)
}

type userApi struct{}

var UserApi uUserApi = userApi{}

// 创建新用户
func (userApi) CreateUser(c *gin.Context) {
	resp := response.Gin{Ctx: c}
	userParam := openapi.CreateUserParam{}

	if err := c.ShouldBind(&userParam); err != nil {
		resp.ClientError("获取参数失败！")
		return
	}

	if err := server.CreateUser(&userParam); err != nil {
		resp.ClientError("创建失败！")
		return
	}

	response := openapi.CommonResponse{
		Message: "创建成功",
		Status:  openapi.RESPONSESTATUS_SUCCESS,
	}

	resp.Success(response)
}

// 获取当前用户信息
func (userApi) GetUserInfo(c *gin.Context) {
	resp := response.Gin{Ctx: c}
	userId := util.GetUserId(c)

	data, err := server.GetUserInfo(userId)

	if err != nil {
		resp.ClientError("获取用户信息失败！")
		return
	}

	response := openapi.GetUserInfoResponse{
		Message: "获取成功",
		Status:  openapi.RESPONSESTATUS_SUCCESS,
		Data:    *data,
	}

	resp.Success(response)
}
