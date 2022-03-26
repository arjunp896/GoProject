package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var con, err = sql.Open("sqlite3", "./pkg/db/UrbanAuto.db")

func GetConnection() (*sql.DB, error) {

	return con, err

}

func CloseConnection() {

	con.Close()
}
