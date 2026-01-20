package web

import (
	"golang/internal/domain"
	"golang/internal/service"
	res "golang/pkg"

	"github.com/gin-gonic/gin"
)

type LanguagesHandler struct {
	svc service.LanguagesService
}

func NewLanguageHandler(svc service.LanguagesService) *LanguagesHandler {
	return &LanguagesHandler{
		svc: svc,
	}
}

func (h *LanguagesHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/api/v1/languages")
	ug.GET("", h.GetLists)
	ug.POST("", h.CreateLanguage)
	ug.PUT("/:id", h.ModifyLanguageById)
	ug.DELETE("/:id", h.RemoveLanguageById)
}

// GetLists 获取语言列表
func (h *LanguagesHandler) GetLists(ctx *gin.Context) {
	languages, err := h.svc.GetLists(ctx)
	switch err {
	case nil:
		res.Success(ctx, languages)
	default:
		res.FailedWithMsg(ctx, "服务器忙碌中!")
	}
}

func (h *LanguagesHandler) CreateLanguage(ctx *gin.Context) {
	type Request struct {
	}
	var request Request
	if err := ctx.Bind(&request); err != nil {
		res.FailedWithMsg(ctx, "服务器忙碌中!")
		return
	}
	err := h.svc.CreateLanguage(ctx, domain.Languages{})
	switch err {
	case nil:
		res.Success(ctx, nil)
	default:
		res.FailedWithMsg(ctx, "服务器忙碌中!")
	}
}

func (h *LanguagesHandler) ModifyLanguageById(ctx *gin.Context) {
	res.Success(ctx, nil)
}

func (h *LanguagesHandler) RemoveLanguageById(ctx *gin.Context) {
	res.Success(ctx, nil)
}
