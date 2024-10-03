package service

import (
	"fmt"
	"rest-api-service/internal/domain"
)

type articleService struct {
	articleStorage articleRepository
}

func (s *articleService) CreateArticle(article *domain.Article) (int, error) {
	id, err := s.articleStorage.Create(article)
	if err != nil {
		return -1, fmt.Errorf("repository.Article.Create: %w", err)
	}

	return id, nil
}
