package controller

import (
	"yugod-backend/app/lib/response"
	"yugod-backend/app/openapi"
	"yugod-backend/app/server"

	"github.com/gin-gonic/gin"
)

type uUserApi interface {
	CreateUser(c *gin.Context)
}

type userApi struct{}

var UserApi uUserApi = userApi{}

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
