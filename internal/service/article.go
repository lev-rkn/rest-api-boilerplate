package service

import (
	"fmt"
	"rest-api-service/internal/domain"
)

//go:generate mockery --name ArticleRepoInterface --output ./mocks
type ArticleRepoInterface interface {
	Create(article *domain.Article) (int, error)
}

type articleService struct {
	articleStorage ArticleRepoInterface
}

func (s *articleService) CreateArticle(article *domain.Article) (int, error) {
	id, err := s.articleStorage.Create(article)
	if err != nil {
		return -1, fmt.Errorf("repository.Article.Create: %w", err)
	}

	return id, nil
}
