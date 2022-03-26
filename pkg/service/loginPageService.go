package service

import (
	"fmt"
	"goproject/pkg/db"
)

func ValidateUser(username string, password string) bool {

	query := fmt.Sprintf(`SELECT user_id 
							FROM Users
							WHERE 	username=%s AND 
									password = %s;`,
		username, password)

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query)

	checkErr(err)

	var userid int = -1

	for rows.Next() {
		err = rows.Scan(&userid)
		checkErr(err)
		fmt.Println(userid)

		return userid != -1

	}

	return false
}

func CheckIsNewUserName(username string) bool {

	query := fmt.Sprintf(`SELECT username FROM Users 
							WHERE username = %s`, username)

	con, err := db.GetConnection()

	checkErr(err)

	row, err := con.Query(query)

	checkErr(err)

	var uname string

	if row.Next() {
		row.Scan(&uname)
		fmt.Println("username :" + uname)

		return false
	}

	return true
}

func CreateUser(username string, password string, name string) bool {

	query := `INSERT INTO Users (username, password, name)
						VALUES (?, ?, ?);`

	con, err := db.GetConnection()

	checkErr(err)

	stmt, err := con.Prepare(query)

	checkErr(err)

	result, err := stmt.Exec(username, password, name)

	checkErr(err)

	noRow, err := result.RowsAffected()

	checkErr(err)

	return noRow > 0
}

func checkErr(err error) {

	if err != nil {
		println("\nError from service")
		panic(err)
	}
}