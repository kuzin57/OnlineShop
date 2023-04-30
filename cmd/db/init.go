package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/tanimutomo/sqlfile"
)

const (
	ddlScripts   = "./cmd/db/scripts/ddl.sql"
	dbConfigFile = "./cmd/config/databases.yaml"
)

var (
	Database *sql.DB
)

func ConnectToDB() error {
	dbConfig := getDBConfig(dbConfigFile)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Connection.Host, dbConfig.Connection.Port, dbConfig.Connection.User,
		dbConfig.Connection.Password, dbConfig.Connection.DBname)

	var err error
	Database, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err = createTables(Database); err != nil {
		return err
	}
	return nil
}

func createTables(db *sql.DB) error {
	s := sqlfile.New()

	if err := s.File(ddlScripts); err != nil {
		return err
	}

	_, err := s.Exec(db)
	if err != nil {
		return err
	}

	return nil
}
