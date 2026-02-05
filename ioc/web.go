package ioc

import (
	"golang/internal/web"
	"golang/internal/web/ijwt"
	"golang/internal/web/middleware"
	"golang/pkg/ginx"
	"golang/pkg/ginx/prometheus"

	"strings"

	"github.com/gin-gonic/gin"
	prometheus2 "github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func InitWebServer(
	middlewares []gin.HandlerFunc,
	userHdl *web.UsersHandler,

) *gin.Engine {
	server := gin.Default()
	server.Use(middlewares...)

	userHdl.RegistryRoutes(server)

	return server
}

func InitGinMiddleware(db *gorm.DB, hdl ijwt.Handler) []gin.HandlerFunc {
	// Prometheus Http接口中间件构造
	pb := &prometheus.Builder{
		Namespace: "golang_basic",
		Subsystem: "user",
		Name:      "gin_http",
		Help:      "统计 Gin 服务的Http接口数据",
	}
	// Prometheus Http接口返回状态码监控
	ginx.InitCounter(prometheus2.CounterOpts{
		Namespace: "golang_basic",
		Subsystem: "user",
		Name:      "biz_code",
		Help:      "统计 Http 业务返回状态码",
	})

	// Cors 跨域中间件
	return []gin.HandlerFunc{
		func(context *gin.Context) {
			cors.New(cors.Options{
				AllowCredentials: true,
				AllowedHeaders:   []string{"Content-Type", "Authorization"},
				ExposedHeaders:   []string{"x-jwt-token", "x-refresh-token"},
				AllowOriginFunc: func(origin string) bool {
					if strings.HasPrefix(origin, "http://localhost") {
						return true
					}
					return false
				},
				MaxAge: 12 * 60 * 60,
			})
		},
		// Prometheus Http接口监控
		pb.BuildResponseTime(),
		// Prometheus 活跃请求监控
		pb.BuildActiveRequest(),
		// OTEL-openTelemetry接入
		//otelgin.Middleware("golang_basic"),
		//NewGormHooks().InitHooks(db),
		middleware.NewLoginJWTMiddlewareBuilder(hdl).CheckLogin(),
	}
}
