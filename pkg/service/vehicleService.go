package service

import (
	"database/sql"
	"goproject/pkg/constants"
	"goproject/pkg/db"
	"goproject/pkg/models"
)

func GetCars() []models.Vehicle {

	return GetVehiclesByType(constants.Car)
}

func GetBikes() []models.Vehicle {

	return GetVehiclesByType(constants.Bike)
}

// func GetCarById(id int) models.Vehicle {

// 	return GetVehicleById(id)
// }

// func GetBikeById(id int) models.Vehicle {

// 	return GetVehicleById(id)
// }

func GetVehiclesByType(vehicleType constants.VehicleType) []models.Vehicle {

	// return loadCars()

	query := `SELECT * FROM Vehicle WHERE type_of_vehicle = ?`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, vehicleType)

	checkErr(err)

	var vehicles []models.Vehicle

	for rows.Next() {

		var vehicle models.Vehicle

		rows.Scan(&vehicle.VehicleId, &vehicle.Make,
			&vehicle.Model, &vehicle.Submodel,
			&vehicle.Price, &vehicle.ModelType,
			&vehicle.Category, &vehicle.VehicleType,
			&vehicle.ImageUrl, &vehicle.YearOfMfg)

		// rows.Scan(vehicle)

		vehicles = append(vehicles, vehicle)
	}

	return vehicles
}

func GetVehicleById(id int) models.Vehicle {

	query := `SELECT * FROM Vehicle WHERE vehicle_id = ?`

	con, err := db.GetConnection()

	checkErr(err)

	var vehicle models.Vehicle

	con.QueryRow(query, id).Scan(&vehicle.VehicleId, &vehicle.Make,
		&vehicle.Model, &vehicle.Submodel,
		&vehicle.Price, &vehicle.ModelType,
		&vehicle.Category, &vehicle.VehicleType,
		&vehicle.ImageUrl, &vehicle.YearOfMfg)

	return vehicle
}

func GetCarsByFilter(params ...interface{}) []models.Vehicle {

	params = append(params, constants.Car)

	return GetVehiclesByFilter(params...)
}

func GetBikesByFilter(params ...interface{}) []models.Vehicle {

	params = append(params, constants.Bike)

	return GetVehiclesByFilter(params...)
}

func GetVehiclesByFilter(params ...interface{}) []models.Vehicle {

	var vehicles []models.Vehicle

	var query string

	con, err := db.GetConnection()

	checkErr(err)

	var rows *sql.Rows

	switch len(params) {

	case 2:
		query = `SELECT * FROM Vehicle 
							WHERE   category = ? AND
									type_of_vehicle = ?`
		rows, err = con.Query(query, params[0], params[1])
		checkErr(err)

	case 3:
		query = `SELECT * FROM Vehicle 
							WHERE 	category = ? AND 
									make = ? AND
									type_of_vehicle = ?`

		rows, err = con.Query(query, params[0], params[1], params[2])

		checkErr(err)

	case 4:
		query = `SELECT * FROM Vehicle 
							WHERE 	category = ? AND 
									make = ? AND
									year_of_mfg = ? AND
									type_of_vehicle = ?`

		rows, err = con.Query(query, params[0], params[1], params[2], params[3])
		checkErr(err)
	}

	for rows.Next() {

		var vehicle models.Vehicle

		rows.Scan(&vehicle.VehicleId, &vehicle.Make,
			&vehicle.Model, &vehicle.Submodel,
			&vehicle.Price, &vehicle.ModelType,
			&vehicle.Category, &vehicle.VehicleType,
			&vehicle.ImageUrl, &vehicle.YearOfMfg)

		vehicles = append(vehicles, vehicle)
	}

	return vehicles
}

func GetVehicleDetailsPageData(id int) map[string]interface{} {

	dataMap := make(map[string]interface{})

	vehicle := GetVehicleById(id)

	dataMap["vehicle"] = vehicle

	dataMap["colors"] = GetColorsByMakeModelVehicleType(vehicle.Make, vehicle.Model, vehicle.VehicleType)

	dataMap["engines"] = GetEnginesByMakeVehicleType(vehicle.Make, vehicle.VehicleType)

	if vehicle.VehicleType == constants.Car {

		dataMap["rooftops"] = GetRooftopsByMakeModelVehicleType(vehicle.Make, vehicle.Model, vehicle.VehicleType)
	}

	dataMap["transmission"] = GetTransmissionsByMakeModelSubmodelVehicleType(vehicle.Make, vehicle.Model, vehicle.Submodel, vehicle.VehicleType)

	dataMap["tyres"] = GetTyresByMakeModelVehicleType(vehicle.Make, vehicle.Model, vehicle.VehicleType)

	dataMap["accessories"] = GetAccessoriesByMakeModelVehicleType(vehicle.Make, vehicle.Model, vehicle.VehicleType)

	return dataMap
}
