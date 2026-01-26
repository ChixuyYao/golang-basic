package ijwt

import "github.com/gin-gonic/gin"

type Handler interface {
	CheckSession(ctx *gin.Context, ssid string) error
	ClearToken(ctx *gin.Context) error
	ExtractToken(ctx *gin.Context) string
	SetLoginToken(ctx *gin.Context, uid string) error
	SetJWTToken(ctx *gin.Context, uid string, ssid string) error
	//SetRefreshToken(ctx *gin.Context, uid string, ssid string) error
}
