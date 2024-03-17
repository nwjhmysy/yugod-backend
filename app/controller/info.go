package controller

import (
	"yugod-backend/app/dao"
	"yugod-backend/app/lib/response"
	"yugod-backend/app/openapi"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type iInfoAPI interface {
	GetClickVolume(c *gin.Context)
	SetClickVolumeByTag(c *gin.Context)
}
type infoAPI struct{}

var InfoAPI iInfoAPI = infoAPI{}

func (infoAPI) GetClickVolume(c *gin.Context) {
	resp := response.Gin{Ctx: c}

	// 从数据库拿到点击量
	data, err := dao.Info.GetClickVolume()
	if err != nil {
		resp.ClientError("获取失败！")
		return
	}

	response := openapi.GetClickVolumeResponse{
		Message: "获取成功！",
		Status:  openapi.RESPONSESTATUS_SUCCESS,
		Data:    *data,
	}

	resp.Success(response)
}

func (infoAPI) SetClickVolumeByTag(c *gin.Context) {
	resp := response.Gin{Ctx: c}
	param := openapi.SetClickVolumeByTagParam{}
	// 获取参数
	if err := c.ShouldBindWith(&param, binding.JSON); err != nil {
		resp.ClientError("参数获取失败！")
		return
	}

	if err := dao.Info.SetClickVolumeByTag(param.Tag); err != nil {
		resp.ClientError(param.Tag + "：数据库修改失败！")
		return
	}

	response := openapi.CommonResponse{
		Message: "修改成功！",
		Status:  openapi.RESPONSESTATUS_SUCCESS,
	}

	resp.Success(response)
}
