package repository

import (
	"context"
	"golang/internal/languages/domain"
	"golang/internal/languages/repository/dao"
	"time"
)

type LanguageRepository interface {
	FindAll(ctx context.Context) ([]domain.Languages, error)
	Create(ctx context.Context, data domain.Languages) error
	Change(ctx context.Context, id string, data domain.Languages) error
	Remove(ctx context.Context, id string) error
}

func NewLanguageRepository(dao dao.LanguagesDao) LanguageRepository {
	return &languageRepository{
		dao: dao,
	}
}

type languageRepository struct {
	dao dao.LanguagesDao
}

func (repo *languageRepository) FindAll(ctx context.Context) ([]domain.Languages, error) {
	languages, err := repo.dao.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	var languagesList []domain.Languages
	for _, language := range languages {
		languagesList = append(languagesList, repo.toDomain(language))
	}
	return languagesList, nil
}

func (repo *languageRepository) Create(ctx context.Context, data domain.Languages) error {
	return repo.dao.Insert(ctx, repo.toDao(data))
}

func (repo *languageRepository) Change(ctx context.Context, id string, data domain.Languages) error {
	return repo.dao.Update(ctx, id, repo.toDao(data))
}

func (repo *languageRepository) Remove(ctx context.Context, id string) error {
	return repo.dao.Delete(ctx, id)
}

// 转换领域对象
func (repo *languageRepository) toDomain(entity dao.Languages) domain.Languages {
	return domain.Languages{
		Id:        entity.Id,
		ISO:       entity.ISO,
		AliasCn:   entity.AliasCn,
		AliasEn:   entity.AliasEn,
		Name:      entity.Name,
		Script:    entity.Script,
		Direction: entity.Direction,
		Family:    entity.Family,
		IsActive:  entity.IsActive,
		CreatedAt: time.UnixMilli(entity.CreatedAt),
		UpdatedAt: time.UnixMilli(entity.UpdatedAt),
	}
}

// 转换实体对象
func (repo *languageRepository) toDao(domain domain.Languages) dao.Languages {
	return dao.Languages{
		ISO:       domain.ISO,
		AliasCn:   domain.AliasCn,
		AliasEn:   domain.AliasEn,
		Name:      domain.Name,
		Script:    domain.Script,
		Direction: domain.Direction,
		Family:    domain.Family,
		IsActive:  domain.IsActive,
	}
}
