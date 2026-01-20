package middleware

import (
	"golang/internal/web"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginJWTMiddlewareBuilder struct {
}

func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		allowURL := []string{"/users/signup", "/users/login"}
		for _, url := range allowURL {
			if url == path {
				return // 对该请求地址不予校验
			}
		}

		// 按约定,JWT签发的TOKEN需要于请求头中的Authorization中带回(Bearer xxx)
		authCode := c.GetHeader("Authorization")
		if authCode == "" {
			// 不存在TOKEN信息,无登录态
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segments := strings.Split(authCode, " ")
		if len(segments) != 2 {
			// TOKEN信息无效,无登录态
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segments[1]
		var uc web.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			// 前述登录方法中用于签发JWT的Key值
			return "secret", nil
		})
		if err != nil {
			// 伪造的TOKEN信息
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid { // token == nil || !token.Valid || expireTime.Before(time.now())
			// 解析TOKEN为非法形式,过期形式
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		expireTime := uc.ExpiresAt

		// ExpireTime - Now = Keep Refresh Duration
		if expireTime.Sub(time.Now()) < time.Minute*5 {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 30))
			tokenStr, err = token.SignedString([]byte("secret"))
			c.Header("x-jwt-token", tokenStr)
			if err != nil {
				log.Println()
			}
		}
		c.Set("uc", uc)
	}
}
