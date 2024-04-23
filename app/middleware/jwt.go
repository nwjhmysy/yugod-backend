package middleware

import (
	"log"
	"time"
	"yugod-backend/app/lib/response"
	"yugod-backend/app/model"
	"yugod-backend/app/openapi"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func initAdminAuth() {
	// JWT 认证中间件配置
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   "id",
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		// 登陆验证：
		// 返回的结构体作为参数传入 Authorizator 和 PayloadFunc
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals model.LoginParam
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			// 验证用户，检查用户名和密码是否正确
			if email == "634365439@qq.com" && password == "123456" {
				return &model.User{
					Email: "634365439@qq.com",
					Name:  "test",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// 权限验证（可选）
		// 根据 token 验证后的数据，自定义权限的验证
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// 授权逻辑，可以根据需求自定义
			if v, ok := data.(*model.User); ok && v.Name == "admin" {
				return true
			}

			return true
		},
		// 将信息添加到 JWT 令牌中
		// 将参数中需要的数据添加到 JWT 令牌中，会返回一个 map[string]interface{} 类型的字典
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"email": v.Email,
					"token": v.Password,
				}
			}
			return jwt.MapClaims{}
		},
		// 认证失败
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"code": code, "message": message})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			resp := response.Gin{Ctx: c}
			response := openapi.CommonResponse{
				Message: "Logout success",
				Status:  openapi.RESPONSESTATUS_SUCCESS,
			}

			resp.Success(response)
		},
	})

	if err != nil {
		log.Fatal("Admin JWT Error:" + err.Error())
	}
	AuthMiddleware = authMiddleware
}

func init() {
	initAdminAuth()
}
