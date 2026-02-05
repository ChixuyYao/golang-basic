package web

import (
	"golang/internal/web/ijwt"
	"golang/pkg/ginx"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct{}

func (h *ArticleHandler) RegistryRoutes(ctx *gin.Engine) {
	g := ctx.Group("/articles")
	g.POST("edit", ginx.WrapBodyWithClaim(h.Edit))
}

type ArticleEditRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 接受Article输入,返回ID
func (h *ArticleHandler) Edit(ctx *gin.Context, request ArticleEditRequest, uc ijwt.UserClaims) (ginx.Result, error) {
	// 业务代码,调用Service层
	//id,err := h.Save(ctx, domain.Article{
	//	Id: request.Id
	//	// ...
	//})

	//if err != nil {
	//	return ginx.Result{
	//		Msg: "系统错误",
	//	},err
	//}
	return ginx.Result{Data: "id"}, nil
}
