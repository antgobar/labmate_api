package store

import (
	"context"
	"database/sql"
)

type PostsStore struct {
	df *sql.DB
}

func (s *PostsStore) Create(ctx context.Context) error {
	return nil
}
