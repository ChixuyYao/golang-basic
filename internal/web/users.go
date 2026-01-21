package web

import (
	"golang/internal/domain"
	"golang/internal/service"
	res "golang/pkg"
	"time"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UsersHandler struct {
	emailRegExp    *regexp.Regexp
	passwordRegExp *regexp.Regexp
	svc            service.UserService
}

func NewUserHandler(svc service.UserService) *UsersHandler {
	return &UsersHandler{
		emailRegExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:            svc,
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
		res.Failed(ctx)
		return
	}

	// 邮箱校验
	isEmail, err := h.emailRegExp.MatchString(request.Email)
	if err != nil || !isEmail {
		res.FailedWithMsg(ctx, "邮箱格式错误,请重试")
		return
	}

	// 密码校验
	if request.Password != request.ConfirmPassword {
		res.FailedWithMsg(ctx, "两次密码输入不一致,请重试")
		return
	}
	isPassword, err := h.passwordRegExp.MatchString(request.Password)
	if err != nil || !isPassword {
		res.FailedWithMsg(ctx, "密码格式错误,请重试")
		return
	}

	err = h.svc.UserSignup(ctx, domain.Users{
		Email:    request.Email,
		Password: request.Password,
	})

	switch err {
	case nil:
		res.SuccessWithMsg(ctx, "账户注册成功", nil)
	case service.ErrDuplicateEmail:
		res.FailedWithMsg(ctx, "邮箱冲突,请更换后重试")
	default:
		res.FailedWithMsg(ctx, "服务错误")
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
		res.Failed(ctx)
		return
	}

	user, err := h.svc.UserLogin(ctx, request.Email, request.Phone, request.Password)
	switch err {
	case nil:
		err = h.SetJWTToken(ctx, user.Id)
		if err != nil {
			res.Failed(ctx)
			return
		}
		type Response struct {
			Id       string `json:"id"`
			Username string `json:"username"`
		}

		res.SuccessWithMsg(ctx, "登录成功", Response{
			Id:       user.Id,
			Username: user.Username,
		})
	case service.ErrInvalidUserOrPassword:
		res.FailedWithMsg(ctx, "账户,密码错误")
	default:
		res.FailedWithMsg(ctx, "服务错误")
	}
}

// -----JWT签发部分----------------------------------------------------------------------------------------------------

type UserClaims struct {
	jwt.RegisteredClaims
	Uid string
}

var JWTKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")

func (h *UsersHandler) SetJWTToken(ctx *gin.Context, uid string) error {
	uc := UserClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)), // JWT 有效期
		},
	}
	// 生成JWT TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
	// 签发JWT TOKEN
	JWT, err := token.SignedString(JWTKey)
	if err != nil {
		return err
	}
	// 写入响应体,同时需要使用ExposeHeaders暴露到前台
	ctx.Header("x-jwt-token", JWT)
	return nil
}
