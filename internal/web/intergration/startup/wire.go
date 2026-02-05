//go:build wireinject

package startup

import (
	"golang/internal/repository"
	"golang/internal/repository/dao"
	"golang/internal/service"
	"golang/internal/web"
	"golang/internal/web/ijwt"
	"golang/ioc"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖初始化
		InitRedis,
		InitDB,
		// Dao 层初始化
		dao.NewUserDao,

		// Repository 层初始化
		repository.NewUserRepository,

		// Service 层初始化
		service.NewUserService,

		// Handler 服务(路由挂载)
		ijwt.NewRedisJwtHandler,
		web.NewUserHandler,

		// 其他依赖
		ioc.InitGinMiddleware,
		ioc.InitWebServer,
	)
	return gin.Default()
}
