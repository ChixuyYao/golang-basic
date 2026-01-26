package web

import (
	"golang/internal/domain"
	"golang/internal/service"
	res "golang/pkg"

	"github.com/gin-gonic/gin"
)

type CategoriesHandler struct {
	svc service.CategoriesService
}

func NewCategoriesHandler(svc service.CategoriesService) *CategoriesHandler {
	return &CategoriesHandler{svc: svc}
}

func (h *CategoriesHandler) RegistryRoutes(server *gin.Engine) {
	ug := server.Group("/api/v1/categories")
	ug.GET("/all", h.FindCategoriesByUid)
	ug.POST("/add", h.AddCategory)
	//ug.POST("/change", h.ChangeCategory)
	//ug.POST("/rm", h.RemoveCategoryById)
}

func (h *CategoriesHandler) FindCategoriesByUid(ctx *gin.Context) {
	uid := ctx.MustGet("uid").(UserClaims)
	categories, err := h.svc.FindCategoriesByUid(ctx, uid.Uid)
	switch err {
	case nil:
		res.Success(ctx, categories)
	default:
		res.Failed(ctx)
	}
}

func (h *CategoriesHandler) AddCategory(ctx *gin.Context) {
	uid := ctx.MustGet("uid").(UserClaims)
	var request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Color       string `json:"color"`
		Icon        string `json:"icon"`
		Type        int    `json:"type"`
		SortOrder   int    `json:"sort_order"`
	}
	if err := ctx.Bind(&request); err != nil {
		res.Failed(ctx)
		return
	}

	err := h.svc.AddCategory(ctx, uid.Uid, domain.Categories{
		Name:        request.Name,
		Description: request.Description,
		Color:       request.Color,
		Icon:        request.Icon,
		Type:        request.Type,
		SortOrder:   request.SortOrder,
		IsActive:    1,
	})

	switch err {
	case nil:
		res.SuccessWithMsg(ctx, "数据修改成功", struct{}{})
	default:
		res.Failed(ctx)
	}
}
