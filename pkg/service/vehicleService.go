package service

import (
	"goproject/pkg/constants"
	"goproject/pkg/db"
	"goproject/pkg/models"
)

func GetCars() []models.Vehicle {

	// return loadCars()

	query := `SELECT * FROM Vehicle WHERE type_of_vehicle = ?`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, constants.Car)

	checkErr(err)

	var vehicles []models.Vehicle

	for rows.Next() {

		var vehicle models.Vehicle

		rows.Scan(&vehicle.VehicleId, &vehicle.Make,
			&vehicle.Model, &vehicle.Submodel,
			&vehicle.Price, &vehicle.ModelType,
			&vehicle.Category, &vehicle.VehicleType,
			&vehicle.ImageUrl, &vehicle.YearOfMfg)

		rows.Scan(vehicle)

		// fmt.Println(vehicle)

		vehicles = append(vehicles, vehicle)

	}
	return vehicles
}
