package ijwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler interface {
	ExtractToken(ctx *gin.Context) string
	SetLoginToken(ctx *gin.Context, uid string) error
	//SetJWTToken(ctx *gin.Context, uid string) error
	SetJWTToken(ctx *gin.Context, uid string, ssid string) error
	ClearToken(ctx *gin.Context) error
	CheckSession(ctx *gin.Context, ssid string) error
}

var (
	RefreshKey = []byte("k6CswdUm77WKcbM683jfuxVsHSpTCwgK")
	JwtKey     = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")
)

type RefreshClaims struct {
	jwt.RegisteredClaims
	Uid string
	// Redis 服务使用
	Ssid string
}

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       string
	UserAgent string
	// Redis 服务使用
	Ssid string
}
