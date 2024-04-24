package router

import (
	"yugod-backend/app/controller"
	"yugod-backend/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupApiRouter(engine *gin.Engine) {
	apiRouter := engine.Group("api")
	// login
	apiRouter.POST("/login", middleware.AuthMiddleware.LoginHandler)

	// auth
	authApiRouter := apiRouter.Group("auth", middleware.AuthMiddleware.MiddlewareFunc())
	authApiRouter.GET("/test", controller.GetAuthTest)
	authApiRouter.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
	// LogoutHandler can be used by clients to remove the jwt cookie (if set)
	// 退出登陆时只需清空前端的 token
	authApiRouter.GET("/logout", middleware.AuthMiddleware.LogoutHandler)

	// markdown 相关
	mdApiRouter := apiRouter.Group("md")
	mdApiRouter.GET("", controller.MdAPI.GetMarkDownByPath)
	mdApiRouter.GET("/download", controller.MdAPI.DownloadMDByCode)

	// info 相关
	infoApiRouter := apiRouter.Group("info")
	infoApiRouter.GET("", controller.InfoAPI.GetClickVolume)
	infoApiRouter.POST("", controller.InfoAPI.SetClickVolumeByTag)

	// user
	userApiRouter := apiRouter.Group("user")
	userApiRouter.POST("/create", controller.UserApi.CreateUser)
}
