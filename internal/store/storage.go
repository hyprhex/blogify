package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStor{db},
	}
}
