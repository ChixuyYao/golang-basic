package repository

import (
	"context"
	"database/sql"
	"golang/internal/domain"
	"golang/internal/repository/dao"
	"time"

	"github.com/ecodeclub/ekit/slice"
)

func NewCategoriesRepository(dao dao.CategoriesDao) CategoriesRepository {
	return &categoriesRepository{
		dao: dao,
	}
}

type CategoriesRepository interface {
	FindByUid(ctx context.Context, uid string) ([]domain.Categories, error)
	Create(ctx context.Context, uid string, c domain.Categories) error
	Modify(ctx context.Context, id string, c domain.Categories) error
	Remove(ctx context.Context, id string) error
}

type categoriesRepository struct {
	dao dao.CategoriesDao
}

func (repo *categoriesRepository) FindByUid(ctx context.Context, uid string) ([]domain.Categories, error) {
	categories, err := repo.dao.SelectByUid(ctx, uid)
	if err != nil {
		return []domain.Categories{}, err
	}
	return slice.Map[dao.Categories, domain.Categories](categories, func(idx int, src dao.Categories) domain.Categories {
		return repo.toDomain(src)
	}), nil
}

func (repo *categoriesRepository) Create(ctx context.Context, uid string, c domain.Categories) error {
	return repo.dao.InsertWithUid(ctx, uid, repo.toEntity(c))
}

func (repo *categoriesRepository) Modify(ctx context.Context, id string, c domain.Categories) error {
	return repo.dao.Update(ctx, id, repo.toEntity(c))
}

func (repo *categoriesRepository) Remove(ctx context.Context, id string) error {
	return repo.dao.Delete(ctx, id)
}

// 数据库对象 -> 领域对象
func (repo *categoriesRepository) toDomain(entity dao.Categories) domain.Categories {
	return domain.Categories{
		Id:          entity.Id,
		Name:        entity.Name.String,
		Description: entity.Description,
		Color:       entity.Color,
		Icon:        entity.Icon,
		Type:        entity.Type,
		SortOrder:   entity.SortOrder,
		IsActive:    entity.IsActive,
		CreatedAt:   time.UnixMilli(entity.CreatedAt),
		UpdatedAt:   time.UnixMilli(entity.UpdatedAt),
	}
}

// 领域对象 -> 数据库对象
func (repo *categoriesRepository) toEntity(domain domain.Categories) dao.Categories {
	return dao.Categories{
		Name:        sql.NullString{String: domain.Name, Valid: domain.Name != ""},
		Description: domain.Description,
		Color:       domain.Color,
		Icon:        domain.Icon,
		Type:        domain.Type,
		SortOrder:   domain.SortOrder,
		IsActive:    domain.IsActive,
	}
}
