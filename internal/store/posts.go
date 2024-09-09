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

	ctx, cancal := context.WithTimeout(ctx, QueryTime)
	defer cancal()

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

	ctx, cancal := context.WithTimeout(ctx, QueryTime)
	defer cancal()

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

func (s *PostStor) Delete(ctx context.Context, id int64) error {
	query := `delete from posts where id = $1`

	ctx, cancal := context.WithTimeout(ctx, QueryTime)
	defer cancal()

	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostStor) Update(ctx context.Context, post *Post) error {
	query := `update posts set title = $1, content = $2, category = $3, tags = $4, updated_at = now() where id = $5`

	ctx, cancal := context.WithTimeout(ctx, QueryTime)
	defer cancal()

	_, err := s.db.ExecContext(ctx, query, post.Title, post.Content, post.Category, pq.Array(post.Tags), post.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStor) List(ctx context.Context) ([]Post, error) {
	query := `select id, title, content, category, tags, created_at, updated_at from posts`

	ctx, cancal := context.WithTimeout(ctx, QueryTime)
	defer cancal()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []Post{}

	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Category, pq.Array(&p.Tags), &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, nil
}
