package service

import "rest-api-service/internal/domain"

//go:generate mockery --name articleRepository --output ./mocks
type articleRepository interface {
	Create(article *domain.Article) (int, error)
}