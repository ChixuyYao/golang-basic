package ginx

import (
	"golang/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prometheus/client_golang/prometheus"
)

var L logger.LoggerV1 = logger.NewNoLogger() // 注入日志模块
var vector *prometheus.CounterVec

// WrapBodyWithClaim 包装请求信息,附加用户认证信息(User Claim)
func WrapBodyWithClaim[Request any, Claim jwt.Claims](bizFn func(ctx *gin.Context, request Request, uc Claim) (Result, error)) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request Request
		if err := context.Bind(&request); err != nil {
			L.Error("参数错误", logger.Error(err))
			return // 参数错误
		}
		L.Debug("请求参数:", logger.Field{Key: "Request Param", Val: request})

		val, ok := context.Get("user")
		if !ok {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		uc, ok := val.(Claim)
		if !ok {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 业务逻辑
		result, err := bizFn(context, request, uc)
		// Prometheus 提交数据
		vector.WithLabelValues(result.Code).Inc()

		if err != nil {
			L.Error("业务逻辑执行失败", logger.Error(err))
		}
		context.JSON(http.StatusOK, result)
	}
}

// WrapBody 包装请求信息
func WrapBody[Request any, Claim jwt.Claims](bizFn func(ctx *gin.Context, request Request) (Result, error)) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request Request
		if err := context.Bind(&request); err != nil {
			L.Error("参数错误", logger.Error(err))
			return // 参数错误
		}
		L.Debug("请求参数:", logger.Field{Key: "Request Param", Val: request})

		// 业务逻辑
		result, err := bizFn(context, request)
		// Prometheus 提交数据
		vector.WithLabelValues(result.Code).Inc()

		if err != nil {
			L.Error("业务逻辑执行失败", logger.Error(err))
		}
		context.JSON(http.StatusOK, result)
	}
}

// WrapClaim 包装认证信息
func WrapClaim[Claim any](bizFn func(ctx *gin.Context, uc Claim) (Result, error)) gin.HandlerFunc {
	return func(context *gin.Context) {
		val, ok := context.Get("user")
		if !ok {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		uc, ok := val.(Claim)
		if !ok {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 业务逻辑
		result, err := bizFn(context, uc)
		// Prometheus 提交数据
		vector.WithLabelValues(result.Code).Inc()

		if err != nil {
			L.Error("业务逻辑执行失败", logger.Error(err))
		}
		context.JSON(http.StatusOK, result)
	}
}

// InitCounter Prometheus监控接口返回码
func InitCounter(opt prometheus.CounterOpts) {
	vector = prometheus.NewCounterVec(opt, []string{"code"})
	prometheus.MustRegister(vector)
}
