package db

import (
	"database/sql"
	"fmt"

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

func (r *Repository) GetProducts() ([]Product, error) {
	query := fmt.Sprintf(`SELECT name, brand, category, rating,
							price, available FROM %s;`, productsTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var (
		name      string
		brand     string
		category  string
		price     uint32
		rating    float64
		available bool
		products  []Product
	)

	for rows.Next() {
		rows.Scan(&name, &brand, &category, &rating,
			&price, &available)
		products = append(
			products,
			NewProduct(
				category, name, brand, price,
				available, rating,
			),
		)
	}

	return products, nil
}
