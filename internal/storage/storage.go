package storage

import (
	"context"
	"rest-api-service/internal/config"
	"rest-api-service/internal/logger"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Storage struct {
	Article *articleRepo
}

func NewStorage(ctx context.Context) *Storage {
	pool, err := pgxpool.New(ctx, config.Cfg.Postgres.PgUrl)
	if err != nil {
		logger.ErrorLog("Unable to connect to database", err)
	}
	pool.Config().MaxConns = config.Cfg.Postgres.MaxConnections

	var storage = &Storage{
		Article: NewArticleRepo(ctx, pool),
	}

	return storage
}
