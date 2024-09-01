package store

import (
	"context"
	"database/sql"

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
