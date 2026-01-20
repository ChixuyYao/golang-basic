package main

import (
	"fmt"
	"golang/config"
	languageDao "golang/internal/languages/reoisitory/dao"
	"golang/internal/rbac/repository"
	rbacDao "golang/internal/rbac/repository/dao"
	"golang/internal/rbac/service"
	"golang/internal/rbac/web"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 数据库初始化
	rbacDB := initDatabase()
	langDB := initLanguageDatabase()
	fmt.Println(langDB.Name())
	// 初始化 WEB 服务,挂载中间件
	server := initWebServer()

	// 注册接口信息
	initUserHandler(rbacDB, server) // 注册 user 相关接口

	// 注册接口监听
	server.Run(":8080")
}

// initDatabase 初始化RBAC服务数据库
func initDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic(err) //数据库启动失败,终止进程
	}
	err = rbacDao.InitTables(db)
	if err != nil {
		panic(err) //数据库表创建失败,终止进程
	}
	return db

}

// initLanguageDatabase 初始化Languages服务数据库
func initLanguageDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.LanguageConfig.DB.DSN))
	if err != nil {
		panic(err) //数据库启动失败,终止进程
	}
	err = languageDao.InitTables(db)
	if err != nil {
		panic(err) //数据库表创建失败,终止进程
	}
	return db

}

func initWebServer() *gin.Engine {
	server := gin.Default()

	// 处理请求跨域
	server.Use(func(context *gin.Context) {
		cors.New(cors.Options{
			//AllowedOrigins: []string{"*"},
			//AllowedMethods: []string{"GET", "POST"},
			AllowCredentials: true,
			AllowedHeaders:   []string{"Content-Type"},
			ExposedHeaders:   []string{},
			AllowOriginFunc: func(origin string) bool {
				if strings.HasPrefix(origin, "http://localhost") {
					return true
				}
				return false
			},
			MaxAge: 12 * 60 * 60,
		})
	})

	return server
}

func initUserHandler(db *gorm.DB, server *gin.Engine) {
	ud := rbacDao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	uh := web.NewUserHandler(us)
	uh.RegistryRoutes(server)
}
