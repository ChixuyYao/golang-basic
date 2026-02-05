package middleware

import (
	"golang/internal/web/ijwt"
	"golang/pkg/ginx"
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
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
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
		tokenStr := m.ExtractToken(ctx)

		var uc ijwt.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			// 前述登录方法中用于签发JWT的Key值
			return ijwt.JwtKey, nil
		})
		if err != nil {
			// 伪造的TOKEN信息
			ctx.JSON(http.StatusOK, ginx.Result{Msg: "服务器繁忙..."})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid { // token == nil || !token.Valid || expireTime.Before(time.now())
			// 解析TOKEN为非法形式,过期形式
			ctx.JSON(http.StatusOK, ginx.Result{Msg: "服务器繁忙..."})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//err = m.CheckSession(ctx, uc.Ssid)
		//if err != nil {
		//	ctx.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		ctx.Set("uc", uc)
	}
}
