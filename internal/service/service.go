package service

import (
	"rest-api-service/internal/storage"
)

type Service struct {
	Article *articleService
}

func NewService(
	storage *storage.Storage,
) *Service {
	return &Service{
		Article: &articleService{articleStorage: storage.Article},
	}
}


