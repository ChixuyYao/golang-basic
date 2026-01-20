package service

import (
	"context"
	"golang/internal/domain"
	"golang/internal/repository"
)

type LanguagesService interface {
	GetLists(ctx context.Context) ([]domain.Languages, error)
	CreateLanguage(ctx context.Context, language domain.Languages) error
	ModifyLanguageById(ctx context.Context, id string, language domain.Languages) error
	RemoveLanguageById(ctx context.Context, id string) error
}

func NewLanguagesService(repo repository.LanguageRepository) LanguagesService {
	return &languageService{
		repo: repo,
	}
}

type languageService struct {
	repo repository.LanguageRepository
}

func (svc *languageService) GetLists(ctx context.Context) ([]domain.Languages, error) {
	languages, err := svc.repo.FindAll(ctx)
	if err != nil {
		return []domain.Languages{}, err
	}
	return languages, nil
}

func (svc *languageService) CreateLanguage(ctx context.Context, language domain.Languages) error {
	return svc.repo.Create(ctx, language)
}

func (svc *languageService) ModifyLanguageById(ctx context.Context, id string, language domain.Languages) error {
	return svc.repo.Change(ctx, id, language)
}

func (svc *languageService) RemoveLanguageById(ctx context.Context, id string) error {
	return svc.repo.Remove(ctx, id)
}
