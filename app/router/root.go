package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRootRouter(engine *gin.Engine) {
	rootRouter := engine.Group("/")
	rootRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running.")
	})
}
