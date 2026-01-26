package ioc

import (
	"golang/internal/web"
	"golang/internal/web/middleware"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func InitWebServer(
	middlewares []gin.HandlerFunc,
	languageHdl *web.LanguagesHandler,
	userHdl *web.UsersHandler,
	categoryHdl *web.CategoriesHandler,
) *gin.Engine {
	server := gin.Default()
	server.Use(middlewares...)
	languageHdl.RegisterRoutes(server)
	userHdl.RegistryRoutes(server)
	categoryHdl.RegistryRoutes(server)
	return server
}

func InitGinMiddleware(db *gorm.DB) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(context *gin.Context) {
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
		},
		NewGormHooks().InitHooks(db),
		middleware.NewLoginJWTMiddlewareBuilder().CheckLogin(),
	}
}
