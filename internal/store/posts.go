package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostStor struct {
	db *sql.DB
}

func (s *PostStor) Create(ctx context.Context, post *Post) error {
	query := `
    INSERT INTO posts (title, content, category, tags)
    VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
  `

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.Category,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStor) GetById(ctx context.Context, id int64) (*Post, error) {
	query := `select id, title, content, category, created_at, updated_at, tags from posts where id = $1`

	var post Post

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Category,
		&post.CreatedAt,
		&post.UpdatedAt,
		pq.Array(&post.Tags),
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}
