package web

import (
	"golang/internal/domain"
	"golang/internal/service"
	"golang/internal/web/ijwt"
	"golang/pkg/ginx"
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UsersHandler struct {
	ijwt.Handler
	emailRegExp    *regexp.Regexp
	passwordRegExp *regexp.Regexp
	svc            service.UserService
}

func NewUserHandler(svc service.UserService, handler ijwt.Handler) *UsersHandler {
	return &UsersHandler{
		emailRegExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:            svc,
		Handler:        handler,
	}
}

func (h *UsersHandler) RegistryRoutes(server *gin.Engine) {
	ug := server.Group("/api/v1/users")
	ug.POST("/signup", h.UserSignup)
	ug.POST("/login", h.UserLogin)
}

// 正则表达式
const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

// UserSignup 用户注册接口
func (h *UsersHandler) UserSignup(ctx *gin.Context) {
	type Request struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var request Request
	// 请求体解析失败
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "服务器繁忙..."})
		return
	}

	// 邮箱校验
	isEmail, err := h.emailRegExp.MatchString(request.Email)
	if err != nil || !isEmail {
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "邮箱格式错误,请重试"})
		return
	}

	// 密码校验
	if request.Password != request.ConfirmPassword {
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "两次密码输入不一致,请重试"})
		return
	}
	isPassword, err := h.passwordRegExp.MatchString(request.Password)
	if err != nil || !isPassword {
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "密码格式错误,请重试"})
		return
	}

	err = h.svc.UserSignup(ctx, domain.Users{
		Email:    request.Email,
		Password: request.Password,
	})

	switch err {
	case nil:
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "注册成功"})
	case service.ErrDuplicateEmail:
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "邮箱冲突,请重试"})
	default:
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "服务器繁忙..."})

	}
}

// UserLogin 用户登录接口
func (h *UsersHandler) UserLogin(ctx *gin.Context) {
	type Request struct {
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	var request Request
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "服务器繁忙..."})
		return
	}

	user, err := h.svc.UserLogin(ctx, request.Email, request.Phone, request.Password)
	switch err {
	case nil:
		err = h.SetLoginToken(ctx, user.Id)
		if err != nil {
			ctx.JSON(http.StatusOK, ginx.Result{Msg: "服务器繁忙..."})
			return
		}
		type Response struct {
			Id       string `json:"id"`
			Username string `json:"username"`
		}

		ctx.JSON(http.StatusOK, ginx.Result{Msg: "登录成功", Data: Response{Id: user.Username, Username: user.Username}})
	case service.ErrInvalidUserOrPassword:
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "用户名,密码错误"})
	default:
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "服务器繁忙..."})
	}
}

func (h *UsersHandler) UserLogout(ctx *gin.Context) {
	err := h.ClearToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, ginx.Result{Msg: "服务器繁忙..."})
		return
	}
	ctx.JSON(http.StatusOK, ginx.Result{Msg: "退出登录成功"})
}

// RefreshToken 提供token重新签发
func (h *UsersHandler) RefreshToken(ctx *gin.Context) {
	// 按约定,JWT签发的TOKEN需要于请求头中的Authorization中带回(Bearer xxx)
	tokenStr := h.ExtractToken(ctx)

	var rc ijwt.RefreshClaims
	token, err := jwt.ParseWithClaims(tokenStr, &rc, func(token *jwt.Token) (interface{}, error) {
		return ijwt.RefreshKey, nil
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if !token.Valid { // token == nil || !token.Valid || expireTime.Before(time.now())
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//err = h.CheckSession(ctx, rc.Ssid)
	if err != nil { // Token 或 Redis存在异常
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = h.SetJWTToken(ctx, rc.Uid, rc.Ssid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, ginx.Result{Msg: "刷新成功"})
}
