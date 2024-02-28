package controller

import (
	"io"
	"os"
	"yugod-backend/app/lib/response"
	"yugod-backend/app/openapi"

	"github.com/gin-gonic/gin"
)

type TestAPI struct {
	openapi.TestAPI
}

func (api *TestAPI) GetMarkDownByKey(ctx *gin.Context) {
	resp := response.Gin{Ctx: ctx}

	// 打开本地的 Markdown 文件
	file, err := os.Open("mds/ssh_config.md")
	if err != nil {
		resp.ClientError("获取失败！")
		return
	}
	defer file.Close()

	ctx.Header("Content-Type", "text/markdown; charset=utf-8")

	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		resp.ClientError("获取失败！")
		return
	}
}

var TestApi TestAPI
