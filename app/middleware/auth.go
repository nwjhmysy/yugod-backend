package middleware

import (
	"strconv"
	"yugod-backend/app/lib/response"

	"github.com/gin-gonic/gin"
)

var ParamFilter gin.HandlerFunc

func init() {
	// 演示
	// 需求：只有携带了参数value，且 value >= 10 的GET请求才能通过

	ParamFilter = func(ctx *gin.Context) {
		resp := response.Gin{Ctx: ctx}

		// 只允许通过GET请求
		if resp.Ctx.Request.Method != "GET" {
			// 终止请求
			ctx.Abort()
			resp.ClientError("不是GET请求")
			return
		}

		param, err := strconv.Atoi(ctx.Query("value"))

		if err != nil {
			ctx.Abort()
			resp.ClientError("参数获取失败")
			return
		}

		if param < 10 {
			ctx.Abort()
			resp.ClientError("value小于10")
			return
		}
		// 通过请求
		ctx.Next()
	}
}
