package db

import (
	"database/sql"

	"github.com/tanimutomo/sqlfile"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) executeQueriesFromFile(path string) error {
	s := sqlfile.New()

	if err := s.File(path); err != nil {
		return err
	}

	_, err := s.Exec(r.db)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CheckDataBaseAvailable() error {
	if err := r.db.Ping(); err != nil {
		return errDBUnavailable
	}
	return nil
}

func (r *Repository) MakeQueryRow(query string, args ...any) *sql.Row {
	return r.db.QueryRow(query, args...)
}
