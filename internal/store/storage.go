package store

import (
	"context"
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("resouce not found")

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStor{db},
	}
}
