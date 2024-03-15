package controller

import (
	"yugod-backend/app/dao"
	"yugod-backend/app/lib/response"
	"yugod-backend/app/openapi"

	"github.com/gin-gonic/gin"
)

// TODO

type iinfoAPI interface {
	GetClickVolume(c *gin.Context)
	SetClickVolumeByTag(c *gin.Context)
}

var InfoAPI iinfoAPI = &openapi.InfoAPI{}

func GetClickVolume(c *gin.Context) {
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
		Data:    data,
	}

	resp.Success(response)
}

func SetClickVolumeByTag(c *gin.Context) {

}
