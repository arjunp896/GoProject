package main

import (
	"database/sql"
	"goproject/pkg/db"
)

func main() {

	con, err := db.GetConnection()

	checkErr(err)

	createUserTable(con)

	db.CloseConnection()

}

func createUserTable(con *sql.DB) {

	stmt, err := con.Prepare(`CREATE TABLE IF NOT EXISTS Users (
								userid integer PRIMARY KEY,
								name text NOT NULL,
		   						username text NOT NULL,
								password text DEFAULT 1234 NOT NULL
							);`)

	checkErr(err)

	// 	INSERT INTO Users (name, username, password)
	// VALUES ('arjun', '123','123');

	stmt.Exec()
}

func checkErr(err error) {

	if err != nil {
		println("\nError from main 2")
		panic(err)
	}
}
