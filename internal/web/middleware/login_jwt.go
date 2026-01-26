package middleware

import (
	"golang/internal/web/ijwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginJWTMiddlewareBuilder struct {
	ijwt.Handler
}

func NewLoginJWTMiddlewareBuilder(hdl ijwt.Handler) *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{
		Handler: hdl,
	}
}

func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		allowURL := []string{
			"/api/v1/users/signup",
			"/api/v1/users/login",
		}
		for _, url := range allowURL {
			if url == path {
				return // 对该请求地址不予校验
			}
		}

		// 按约定,JWT签发的TOKEN需要于请求头中的Authorization中带回(Bearer xxx)
		tokenStr := m.ExtractToken(c)

		var uc ijwt.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			// 前述登录方法中用于签发JWT的Key值
			return ijwt.JwtKey, nil
		})
		if err != nil {
			// 伪造的TOKEN信息
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "用户身份认证无效,请尝试重新登录!"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid { // token == nil || !token.Valid || expireTime.Before(time.now())
			// 解析TOKEN为非法形式,过期形式
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "用户身份认证无效,请尝试重新登录!"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//expireTime := uc.ExpiresAt

		// ExpireTime - Now = Keep Refresh Duration
		//if expireTime.Sub(time.Now()) < time.Minute*5 {
		//	uc.ExpiresAt = ijwt.NewNumericDate(time.Now().Add(time.Minute * 30))
		//	tokenStr, err = token.SignedString([]byte("secret"))
		//	c.Header("x-ijwt-token", tokenStr)
		//	if err != nil {
		//		log.Println()
		//	}
		//}

		// 前述Token校验后,查阅Redis
		//cnd, err := m.cmd.Exists(c, fmt.Sprintf("users:ssid:%s", uc.Ssid)).Result()
		//if err != nil || cnd > 0 {
		//	c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "用户身份认证无效,请尝试重新登录!"})
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//}

		err = m.CheckSession(c, uc.Ssid)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("uc", uc)
	}
}
