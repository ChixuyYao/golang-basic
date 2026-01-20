package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	successCode     = http.StatusOK
	badRequestCode  = http.StatusBadRequest
	serverErrorCode = http.StatusInternalServerError
)
const (
	successMsg     = "处理成功"
	badRequestMsg  = "处理失败,请重试"
	serverErrorMsg = "服务器繁忙,请稍后再试"
)

// Response 接口返回响应体规范
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

// Success 成功响应,默认返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    successCode,
		Message: successMsg,
		Data:    data,
	})
}

// SuccessWithMsg 成功响应,自定义返回内容
func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    successCode,
		Message: msg,
		Data:    data,
	})
}

// Failed 失败响应,默认返回
func Failed(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code:    badRequestCode,
		Message: badRequestMsg,
	})
}

// FailedWithMsg 失败响应,自定义返回内容
func FailedWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &Response{
		Code:    badRequestCode,
		Message: msg,
	})
}
