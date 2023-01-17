package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/OybekAbduvosiqov/tovar/config"
	"github.com/OybekAbduvosiqov/tovar/storage"
)

type Store struct {
	db    *pgxpool.Pool
	tovar *TovarRepo
	// category *CategoryRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))

	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConn

	pool, err := pgxpool.ConnectConfig(ctx, config)

	if err != nil {
		return nil, err
	}

	return &Store{
		db:    pool,
		tovar: NewTovarRepo(pool),
		// category: NewCategoryRepo(pool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Tovar() storage.TovarRepoI {

	if s.tovar == nil {
		s.tovar = NewTovarRepo(s.db)
	}

	return s.tovar
}

// func (s *Store) Category() storage.CategoryRepoI {

// 	if s.category == nil {
// 		s.category = NewCategoryRepo(s.db)
// 	}

// 	return s.category
// }
