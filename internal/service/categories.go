package service

import (
	"context"
	"golang/internal/domain"
	"golang/internal/repository"
)

func NewCategoriesService(repo repository.CategoriesRepository) CategoriesService {
	return &categoriesService{
		repo: repo,
	}
}

type CategoriesService interface {
	FindCategoriesByUid(ctx context.Context, uid string) ([]domain.Categories, error)
	AddCategory(ctx context.Context, uid string, category domain.Categories) error
	ChangeCategory(ctx context.Context, id string, category domain.Categories) error
	RemoveCategoryById(ctx context.Context, id string) error
}

type categoriesService struct {
	repo repository.CategoriesRepository
}

func (svc *categoriesService) FindCategoriesByUid(ctx context.Context, uid string) ([]domain.Categories, error) {
	categories, err := svc.repo.FindByUid(ctx, uid)
	if err != nil {
		return []domain.Categories{}, err
	}
	return categories, nil
}

func (svc *categoriesService) AddCategory(ctx context.Context, uid string, category domain.Categories) error {
	return svc.repo.Create(ctx, uid, category)
}

func (svc *categoriesService) ChangeCategory(ctx context.Context, id string, category domain.Categories) error {
	return svc.repo.Modify(ctx, id, category)
}

func (svc *categoriesService) RemoveCategoryById(ctx context.Context, id string) error {
	return svc.repo.Remove(ctx, id)
}
