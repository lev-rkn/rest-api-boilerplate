package storage

import (
	"context"
	"rest-api-service/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type articleRepo struct {
	ctx  context.Context
	conn *pgxpool.Pool
}

func NewArticleRepo(ctx context.Context, conn *pgxpool.Pool) *articleRepo {
	return &articleRepo{
		ctx:  ctx,
		conn: conn,
	}
}

func (s *articleRepo) Create(article *domain.Article) (int, error) {
	var id int
	err := s.conn.QueryRow(s.ctx,
		`INSERT INTO articles (title, description, photos, user_id) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id;`,
		article.Title,
		article.Description,
		article.Photos,
		article.UserId).Scan(&id)

	return id, err
}
