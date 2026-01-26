package ijwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisJwtHandler struct {
	client        redis.Cmdable
	signingMethod jwt.SigningMethod
	rcExpiration  time.Duration

	//refreshKey []byte
	//jwtKey     []byte
}

func NewRedisJwtHandler(client redis.Cmdable) *RedisJwtHandler {
	return &RedisJwtHandler{
		client:        client,
		signingMethod: jwt.SigningMethodHS512,
		rcExpiration:  time.Hour * 24 * 7,
		//refreshKey:    []byte("k6CswdUm77WKcbM683jfuxVsHSpTCwgK"),
		//jwtKey:        []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK"),
	}
}

// ClearToken 清空Token方法
func (h *RedisJwtHandler) ClearToken(ctx *gin.Context) error {
	ctx.Header("x-ijwt-token", "")
	ctx.Header("x-refresh-token", "")
	uc := ctx.MustGet("uc").(UserClaims)
	return h.client.
		Set(ctx, fmt.Sprintf("users:ssid:%s", uc.Ssid), "", h.rcExpiration).
		Err()
}

// SetJWTToken 设置短TOKEN(JWT Token)
func (h *RedisJwtHandler) SetJWTToken(ctx *gin.Context, uid string, ssid string) error {
	userClaim := UserClaims{
		Uid:       uid,
		Ssid:      ssid,
		UserAgent: ctx.GetHeader("User-Agent"),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}
	token := jwt.NewWithClaims(h.signingMethod, userClaim)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenString)
	return nil
}

// SetRefreshToken 设置长TOKEN(Refresh Token)
func (h *RedisJwtHandler) SetRefreshToken(ctx *gin.Context, uid string, ssid string) error {
	refreshClaim := RefreshClaims{
		Uid:  uid,
		Ssid: ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.rcExpiration)),
		},
	}
	token := jwt.NewWithClaims(h.signingMethod, refreshClaim)
	tokenString, err := token.SignedString(RefreshKey)
	if err != nil {
		return err
	}
	ctx.Header("x-refresh-token", tokenString)
	return nil
}

// CheckSession 检查用户ssid有效性
func (h *RedisJwtHandler) CheckSession(ctx *gin.Context, ssid string) error {
	cnt, err := h.client.Exists(ctx, fmt.Sprintf("users:ssid:%s", ssid)).Result()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("Token信息无效")
	}
	return nil
}

// SetLoginToken 登录时,同时签发JwtToken,RefreshToken
func (h *RedisJwtHandler) SetLoginToken(ctx *gin.Context, uid string) error {
	ssid := uuid.New().String()
	err := h.SetRefreshToken(ctx, uid, ssid)
	if err != nil {
		//ctx.JSON(http.StatusOK, gin.H{"msg": "系统错误"})
		return err
	}
	return h.SetJWTToken(ctx, uid, ssid)
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	Uid  string
	Ssid string
}

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       string
	Ssid      string
	UserAgent string
}

var (
	RefreshKey = []byte("k6CswdUm77WKcbM683jfuxVsHSpTCwgK")
	JwtKey     = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")
)
