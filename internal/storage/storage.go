package storage

import (
	"database/sql"
	"github.com/ixoja/library/internal/models"
	"github.com/pkg/errors"
)

type SQLite struct {
	Database *sql.DB
}

const (
	createTable = ""
)

func (s *SQLite) InitDB() error {
	stmt, err := s.Database.Prepare(createTable)
	if err != nil {
		return errors.Wrap(err, "failed to prepare statement to init table")
	}

	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "failed to init table")
	}
	return nil
}

func (s *SQLite) Create(book *models.Book) (*models.Book, error) {
	panic("implement me")
}

func (s *SQLite) Delete(id string) error {
	panic("implement me")
}

func (s *SQLite) Get(id string) (*models.Book, error) {
	panic("implement me")
}

func (s *SQLite) GetAll() ([]*models.Book, error) {
	panic("implement me")
}

func (s *SQLite) Rate(id, rate string) error {
	panic("implement me")
}

func (s *SQLite) CheckIn(id string) error {
	panic("implement me")
}

func (s *SQLite) CheckOut(id string) error {
	panic("implement me")
}

func New(db *sql.DB) *SQLite {
	return &SQLite{Database: db}
}