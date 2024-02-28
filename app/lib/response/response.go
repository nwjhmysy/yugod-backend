package response

import (
	"net/http"
	"yugod-backend/app/openapi"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Gin struct {
	Ctx *gin.Context
}

func (r *Gin) SendJSON(code int, obj interface{}) {
	r.Ctx.Header("Content-Type", "application/json; charset=utf-8")
	r.Ctx.Header("Cache-Control", "no-cache")
	r.Ctx.Header("Pragma", "no-cache")
	r.Ctx.Header("Expires", "0")
	r.Ctx.Header("X-Content-Type-Options", "nosniff")
	r.Ctx.JSON(code, obj)
}

// 200
func (r *Gin) Success(response interface{}) {
	r.SendJSON(http.StatusOK, response)
}

// 400
// 请求参数错误
func (r *Gin) ValidationError(translator ut.Translator, error error, nonValidationErrorMessage string) {
	if len(nonValidationErrorMessage) < 1 {
		nonValidationErrorMessage = "Bad request."
	}
	switch error.(type) {
	case validator.ValidationErrors:
		errors := error.(validator.ValidationErrors)
		if len(errors) == 0 {
			r.ClientError(nonValidationErrorMessage)
		} else {
			r.ClientError(errors[0].Translate(translator))
		}
	default:
		r.ClientError(nonValidationErrorMessage)
	}
}

// 400
// 请求失败
func (r *Gin) ClientError(message string) {
	if len(message) < 1 {
		message = "Bad Request."
	}
	response := openapi.CommonResponse{Status: openapi.RESPONSESTATUS_FAIL, Message: message}
	r.SendJSON(http.StatusBadRequest, response)
}

// 413
// 有效载荷过大
func (r *Gin) RequestEntityTooLargeError(message string) {
	if len(message) < 1 {
		message = "Payload Too Large"
	}
	response := openapi.CommonResponse{Status: openapi.RESPONSESTATUS_FAIL, Message: message}
	r.SendJSON(http.StatusRequestEntityTooLarge, response)
}

// 404
// 页面不存在（服务端渲染）
func (r *Gin) NotFound(message string) {
	if len(message) < 1 {
		message = "Not Found"
	}
	response := openapi.CommonResponse{Status: openapi.RESPONSESTATUS_FAIL, Message: message}
	r.SendJSON(http.StatusNotFound, response)
}

// 401
// 权限/认证错误
func (r *Gin) Unauthorized(message string) {
	if len(message) < 1 {
		message = "Unauthorized."
	}
	response := openapi.CommonResponse{Status: openapi.RESPONSESTATUS_FAIL, Message: message}
	r.SendJSON(http.StatusUnauthorized, response)
}

// 403
// 角色权限错误
func (r *Gin) Forbidden(message string) {
	if len(message) < 1 {
		message = "Forbidden."
	}
	response := openapi.CommonResponse{Status: openapi.RESPONSESTATUS_FAIL, Message: message}
	r.SendJSON(http.StatusForbidden, response)
}

// 429
// 访问限制
func (r *Gin) RateLimited(message string) {
	if len(message) < 1 {
		message = "Rate limited."
	}
	response := openapi.CommonResponse{Status: openapi.RESPONSESTATUS_FAIL, Message: message}
	r.SendJSON(http.StatusTooManyRequests, response)
}

// 500
func (r *Gin) ServerError(message string) {
	if len(message) < 1 {
		message = "Server internal error."
	}
	response := openapi.CommonResponse{Status: openapi.RESPONSESTATUS_ERROR, Message: message}
	r.SendJSON(http.StatusInternalServerError, response)
}
