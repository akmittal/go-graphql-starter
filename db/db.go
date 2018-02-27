package db

import (
	"fmt"
	"go-graphql-starter/config"

	"github.com/jmoiron/sqlx"
)

type DB *sqlx.DB

var DBConnection DB

func Connect() (DB, error) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.User, config.Password, config.DBName)
	dbConn, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}
	schema := `CREATE TABLE IF NOT EXISTS USERS (first_name varchar(40), last_name varchar(40), email varchar(140) NOT NULL, username varchar(40) PRIMARY KEY, password varchar(60) NOT NULL)`
	_, err = dbConn.Exec(schema)
	if err != nil {
		return nil, err
	}
	DBConnection = dbConn
	return dbConn, nil
}

func GetConnection() *sqlx.DB {
	return DBConnection
}
