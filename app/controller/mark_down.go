package controller

import (
	"os"
	"strings"
	"yugod-backend/app/config"
	"yugod-backend/app/lib/response"
	"yugod-backend/app/openapi"
	"yugod-backend/app/util"

	"github.com/gin-gonic/gin"
)

type mdAPI struct {
	openapi.MdAPI
}

type MdAPIParam struct {
	MdPath string `form:"md_path"`
}
type DownloadMdAPIParam struct {
	MdPath       string `form:"md_path"`
	DownloadCode uint   `form:"download_code"`
}

func (md *mdAPI) GetMarkDownByPath(ctx *gin.Context) {
	resp := response.Gin{Ctx: ctx}
	param := MdAPIParam{}
	// 获取 md 文件的路径参数
	if err := ctx.ShouldBindQuery(&param); err != nil {
		if err != nil {
			resp.ClientError("参数获取失败！")
			return
		}
	}

	// 打开本地的 Markdown 文件
	file, err := os.Open("mds/" + param.MdPath + ".md")
	if err != nil {
		resp.ClientError("获取失败！")
		return
	}
	defer file.Close()

	content, err := util.EncodeFile(file)

	// _, err = io.Copy(ctx.Writer, file)

	if err != nil {
		resp.ClientError("获取失败！")
		return
	}
	response := openapi.GetMdbyPathResponse{
		Message: "获取成功！",
		Status:  openapi.RESPONSESTATUS_SUCCESS,
		Data:    content,
	}

	resp.Success(response)
}

func (md *mdAPI) DownloadMDByCode(ctx *gin.Context) {
	resp := response.Gin{Ctx: ctx}
	param := DownloadMdAPIParam{}

	// 获取参数 md_path 和 download_code
	if err := ctx.ShouldBindQuery(&param); err != nil {
		resp.ClientError("参数获取失败！")
		return
	}
	// 判断 code 码
	if param.DownloadCode != config.App.DownloadCode {
		resp.ClientError("下载码错误！")
		return
	}

	// 打开本地的 Markdown 文件
	filePath := "mds/" + param.MdPath + ".md"
	_, err := os.Stat(filePath)
	if err != nil {
		resp.ClientError("没找到文件！")
		return
	}

	// 设置响应头，告诉浏览器这是一个文件下载
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+strings.Replace(filePath, "/", "_", -1))
	ctx.Header("Expires", "0")
	ctx.Header("Cache-Control", "must-revalidate")
	ctx.Header("Pragma", "public")
	ctx.File(filePath)
}

var MdAPI = mdAPI{}
