package service

import (
	"fmt"
	"goproject/pkg/db"
	"goproject/pkg/models"
)

func ValidateUser(username string, password string) int {

	query := fmt.Sprintf(`SELECT user_id 
							FROM Users
							WHERE 	username= '%s' AND 
									password = '%s';`,
		username, password)

	// println("query" + query)

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query)

	checkErr(err)

	var userid int = -1

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&userid)
		checkErr(err)
		// fmt.Println(userid)
	}

	return userid
}

func CheckIsNewUserName(username string) bool {

	query := `SELECT username FROM Users WHERE username = ?`

	con, err := db.GetConnection()

	checkErr(err)

	row, err := con.Query(query, username)

	checkErr(err)

	var uname string

	defer row.Close()

	if row.Next() {
		row.Scan(&uname)
		fmt.Println("username :" + uname)

		return false
	}

	return true
}

func GetUserById(id int) models.User {

	query := `SELECT * FROM Users WHERE user_id = ?`

	con, err := db.GetConnection()

	checkErr(err)

	var user models.User

	con.QueryRow(query, id).Scan(&user.UserId, &user.Username, &user.Password, &user.Name)

	return user
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
