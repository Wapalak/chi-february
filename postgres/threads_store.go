package postgres

import (
	"chi_test_second"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id uuid.UUID) (chi_test_second.Thread, error) {
	var t chi_test_second.Thread
	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return chi_test_second.Thread{}, fmt.Errorf("Error getting thread: %w", err)
	}
	return t, nil
}

func (s *ThreadStore) Threads() ([]chi_test_second.Thread, error) {
	var tt []chi_test_second.Thread
	if err := s.Select(&tt, `SELECT * FROM threads`); err != nil {
		return []chi_test_second.Thread{}, fmt.Errorf("Error getting threads: %w", err)
	}
	return tt, nil
}

func (s *ThreadStore) CreateThread(t *chi_test_second.Thread) error {
	if err := s.Get(t, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`,
		t.ID,
		t.Title,
		t.Description); err != nil {
		return fmt.Errorf("Error creating thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) UpdateThread(t *chi_test_second.Thread) error {
	if err := s.Get(t, `UPDATE threads SET VALUES title = $1, description = $2, id = $3 RETURNING *`,
		t.Title,
		t.Description,
		t.ID); err != nil {
		return fmt.Errorf("Error updating thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("Error deleting thread: %w", err)
	}
	return nil
}
