package controller

import (
	"yugod-backend/app/lib/response"
	"yugod-backend/app/openapi"

	"github.com/gin-gonic/gin"
)

func GetAuthTest(c *gin.Context) {
	resp := response.Gin{Ctx: c}

	// 从token中获取 userId
	// userEmail := util.GetUserEmailFromClaims(c)
	// fmt.Println(userEmail)

	response := openapi.CommonResponse{
		Message: "测试成功！",
		Status:  openapi.RESPONSESTATUS_SUCCESS,
	}

	resp.Success(response)
}
