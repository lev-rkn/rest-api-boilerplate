package service

import (
	"fmt"
	"rest-api-service/internal/domain"
	"rest-api-service/internal/storage"
)

type articleService struct {
	storage *storage.Storage
}

var _ ArticleServiceInterface = (*articleService)(nil)

func (s *articleService) CreateArticle(article *domain.Article) (int, error) {
	id, err := s.storage.Article.Create(article)
	if err != nil {
		return -1, fmt.Errorf("repository.Article.Create: %w", err)
	}

	return id, nil
}
