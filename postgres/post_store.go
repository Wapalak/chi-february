package postgres

import (
	"chi_test_second"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (chi_test_second.Post, error) {
	var p chi_test_second.Post
	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return chi_test_second.Post{}, fmt.Errorf("error getting post: %w", err)
	}
	return p, nil
}

func (s *PostStore) PostsByThread(threadID uuid.UUID) ([]chi_test_second.Post, error) {
	var pp []chi_test_second.Post
	if err := s.Select(&pp, `SELECT * FROM posts WHERE thread_id = $1, threadID`); err != nil {
		return []chi_test_second.Post{}, fmt.Errorf("error getting posts: %w", err)
	}
	return pp, nil
}
func (s PostStore) CreatePost(p *chi_test_second.Post) error {
	if err := s.Get(p, `INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		p.ID,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}

func (s PostStore) UpdatePost(p *chi_test_second.Post) error {
	if err := s.Get(p, `UPDATE posts SET thread_id = $1, title = $2, content = $3, votes = $4 WHERE id = $5 RETURNING *`,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes,
		p.ID); err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}
func (s PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
		return nil
	}
	return nil
}
