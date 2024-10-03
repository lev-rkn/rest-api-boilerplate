package controller

import "rest-api-service/internal/domain"

//go:generate mockery --name articleService --output ./mocks
type articleService interface {
	CreateArticle(article *domain.Article) (int, error)
}