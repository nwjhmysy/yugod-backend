package boot

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"yugod-backend/app/config"
	"yugod-backend/app/middleware"
	"yugod-backend/app/router"

	"github.com/gin-gonic/gin"
)

// to launch gin server
func GinServer() {
	// 创建gin实例
	engine := gin.New()

	engine.MaxMultipartMemory = config.App.MaximumUploadFileSize

	// CORS
	engine.Use(middleware.CORS)

	// Routers
	router.SetupRootRouter(engine)
	router.SetupApiRouter(engine)

	// 配置server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.App.Port),
		Handler:      engine,
		ReadTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(60) * time.Second,
	}

	// Start server
	go func() {
		log.Println("Server started.")
		log.Println("Port" + server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()
}

func init() {
	gin.SetMode(config.App.Mode)
}
