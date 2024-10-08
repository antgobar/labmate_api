package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	df *sql.DB
}

func (s *UsersStore) Create(ctx context.Context) error {
	return nil
}
