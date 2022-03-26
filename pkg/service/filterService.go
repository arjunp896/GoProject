package service

import "goproject/pkg/db"

func GetCategories() []string {

	query := `SELECT DISTINCT category FROM Vehicle ORDER BY category`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query)

	checkErr(err)

	var categories []string

	var cat string

	for rows.Next() {

		rows.Scan(&cat)
		categories = append(categories, cat)
	}

	return categories
}

func GetMakesByCategory(category string) []string {

	query := `SELECT DISTINCT make FROM Vehicle WHERE CATEGORY = ? ORDER BY make`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, category)

	checkErr(err)

	var makes []string

	var make string

	for rows.Next() {

		rows.Scan(&make)
		makes = append(makes, make)
	}

	makes = append(makes, "Audi")
	makes = append(makes, "Honda")
	makes = append(makes, "Hundai")

	return makes
}

func GetYearByCategoryAndMake(category string, vehicleMake string) []string {

	query := `SELECT DISTINCT year_of_mfg FROM Vehicle WHERE category = ? 
				AND make = ? ORDER BY year_of_mfg`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, category, vehicleMake)

	checkErr(err)

	var years []string

	var year string

	for rows.Next() {
		rows.Scan(&year)
		years = append(years, year)
	}

	return years
}
