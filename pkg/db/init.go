package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	ddlScripts     = "./pkg/db/scripts/ddl.sql"
	insertsScripts = "./pkg/db/scripts/inserts.sql"
	viewsScripts   = "./pkg/db/scripts/views.sql"
	dbConfigFile   = "./cmd/config/databases.yaml"
)

func ConnectToDB() (*Repository, error) {
	dbConfig := getDBConfig(dbConfigFile)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Connection.Host, dbConfig.Connection.Port, dbConfig.Connection.User,
		dbConfig.Connection.Password, dbConfig.Connection.DBname)

	var err error
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	repos := NewRepository(database)
	if err = repos.executeQueriesFromFile(ddlScripts); err != nil {
		return nil, err
	}

	if err = repos.executeQueriesFromFile(insertsScripts); err != nil {
		return nil, err
	}

	// if err = repos.executeQueriesFromFile(viewsScripts); err != nil {
	// 	return nil, err
	// }

	return repos, nil
}
