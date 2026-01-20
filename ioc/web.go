package ioc

import (
	"golang/internal/web"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func InitWebServer(
	middlewares []gin.HandlerFunc,
	languageHdl *web.LanguagesHandler,
) *gin.Engine {
	server := gin.Default()
	server.Use(middlewares...)
	languageHdl.RegisterRoutes(server)
	return server
}

func InitGinMiddleware() []gin.HandlerFunc {
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
	}
}
