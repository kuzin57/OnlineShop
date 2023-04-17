package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDB() error {
	_, err := sql.Open("postgres", "postgres:postgres@tcp(127.0.0.1:5432)/postgres")
	if err != nil {
		return err
	}

	return nil
}
