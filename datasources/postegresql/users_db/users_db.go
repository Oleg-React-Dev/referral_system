package users_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	PGSQL_USERS_USER     = "PGSQL_USERS_USER"
	PGSQL_USERS_PASSWORD = "PGSQL_USERS_PASSWORD"
	PGSQL_USERS_HOST     = "PGSQL_USERS_HOST"
	PGSQL_USERS_PORT     = "PGSQL_USERS_PORT"
	PGSQL_USERS_DB_NAME  = "PGSQL_USERS_DB_NAME"
)

var (
	Db *sql.DB
)

func InitDB() error {
	user := os.Getenv(PGSQL_USERS_USER)
	password := os.Getenv(PGSQL_USERS_PASSWORD)
	host := os.Getenv(PGSQL_USERS_HOST)
	port := os.Getenv(PGSQL_USERS_PORT)
	dbname := os.Getenv(PGSQL_USERS_DB_NAME)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	return Db.Ping()
}
