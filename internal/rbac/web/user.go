package web

import (
	"golang/internal/rbac/domain"
	"golang/internal/rbac/service"
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	emailRegExp    *regexp.Regexp
	passwordRegExp *regexp.Regexp
	svc            *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRegExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:            svc,
	}
}

func (h *UserHandler) RegistryRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", h.SignUp)
	ug.POST("/login", h.Login)
	ug.POST("/edit", h.Edit)
	ug.GET("/profile", h.Profile)
}

// 正则表达式
const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

// SignUp 用户注册接口
func (h *UserHandler) SignUp(ctx *gin.Context) {
	type Request struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var request Request
	if err := ctx.Bind(&request); err != nil {
		return
	}

	isEmail, err := h.emailRegExp.MatchString(request.Email)
	if err != nil || !isEmail {
		ctx.String(http.StatusOK, "邮箱格式错误")
		return
	}
	// 密码校验问题
	if request.Password != request.ConfirmPassword {
		ctx.String(http.StatusOK, "请保持两次密码一致")
		return
	}
	isPassword, err := h.passwordRegExp.MatchString(request.Password)
	if err != nil || !isPassword {
		ctx.String(http.StatusOK, "密码格式错误")
		return
	}

	// 领域层(Service)服务调用
	err = h.svc.SignUp(ctx, domain.User{
		Email:    request.Email,
		Password: request.Password,
	})

	switch err {
	case nil:
		ctx.String(http.StatusOK, "注册成功")
	case service.ErrDuplicateEmail:
		ctx.String(http.StatusOK, "邮箱冲突,请更换后重试")
	default:
		ctx.String(http.StatusOK, "系统异常")
	}
}

// Login 用户登录接口
func (h *UserHandler) Login(ctx *gin.Context) {

}

// Edit 用户编辑接口
func (h *UserHandler) Edit(ctx *gin.Context) {

}

// Profile 用户个人信息查阅接口
func (h *UserHandler) Profile(ctx *gin.Context) {

}
