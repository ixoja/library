package storage

import "database/sql"

type SQLite struct {
	Database *sql.DB
}

func New(db *sql.DB) *SQLite {
	return &SQLite{Database: db}
}