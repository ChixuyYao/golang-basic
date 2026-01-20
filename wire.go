//go:build wireinject

package main

import (
	"golang/internal/repository"
	"golang/internal/repository/dao"
	"golang/internal/service"
	"golang/internal/web"
	"golang/ioc"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖初始化
		ioc.InitDB,
		// Dao层初始化
		dao.NewLanguageDao,
		// Repository层初始化
		repository.NewLanguageRepository,
		// Service层初始化
		service.NewLanguagesService,
		// Handler服务(路由挂载)
		web.NewLanguageHandler,

		// 服务器本身依赖
		ioc.InitGinMiddleware,
		ioc.InitWebServer,
	)
	return gin.Default()
}
