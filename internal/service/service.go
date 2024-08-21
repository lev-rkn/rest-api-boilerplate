package service

import (
	"rest-api-service/internal/domain"
	"rest-api-service/internal/storage"
)

type Service struct {
	Article ArticleServiceInterface
}

func NewService(
	storage *storage.Storage,
) *Service {
	return &Service{
		Article: &articleService{storage: storage},
	}
}


//go:generate mockery --name ArticleServiceInterface --output ./mocks
type ArticleServiceInterface interface {
	CreateArticle(article *domain.Article) (int, error)
}
