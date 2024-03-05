package router

import (
	"yugod-backend/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupApiRouter(engine *gin.Engine) {
	apiRouter := engine.Group("api")

	// markdown 相关
	mdApiRouter := apiRouter.Group("/md")
	mdApiRouter.GET("", controller.MdAPI.GetMarkDownByPath)
	mdApiRouter.GET("/download", controller.MdAPI.DownloadMDByCode)
}
