package service

import (
	"goproject/pkg/constants"
	"goproject/pkg/db"
	"goproject/pkg/models"
)

func GetCategories(vehicle string) []string {

	query := `SELECT DISTINCT category FROM Vehicle WHERE type_of_vehicle = ? ORDER BY category`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, vehicle)

	checkErr(err)

	var categories []string

	var cat string

	defer rows.Close()

	for rows.Next() {

		rows.Scan(&cat)
		categories = append(categories, cat)
	}

	return categories
}

func GetMakesByCategory(category string, vehicle string) []string {

	query := `SELECT DISTINCT make FROM Vehicle WHERE CATEGORY = ? AND type_of_vehicle = ? ORDER BY make`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, category, vehicle)

	checkErr(err)

	var makes []string

	var make string

	defer rows.Close()

	for rows.Next() {

		rows.Scan(&make)
		makes = append(makes, make)
	}

	// makes = append(makes, "Audi")
	// makes = append(makes, "Honda")
	// makes = append(makes, "Hundai")

	return makes
}

func GetYearByCategoryAndMake(category string, vehicleMake string, vehicle string) []string {

	query := `SELECT DISTINCT year_of_mfg FROM Vehicle WHERE category = ? 
				AND make = ? AND type_of_vehicle = ? ORDER BY year_of_mfg`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, category, vehicleMake, vehicle)

	checkErr(err)

	var years []string

	var year string

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&year)
		years = append(years, year)
	}

	return years
}

func GetAccessoriesByMakeModelVehicleType(make string, model string, vehicleType constants.VehicleType) []models.Accessory {

	query := `SELECT * FROM Accessories WHERE make = ? AND model = ? AND type_of_vehicle = ?`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, make, model, vehicleType)

	checkErr(err)

	var accessories []models.Accessory

	defer rows.Close()

	for rows.Next() {

		var accessory models.Accessory

		rows.Scan(&accessory.AccessoryId, &accessory.AccessoryName,
			&accessory.Description, &accessory.Make,
			&accessory.Model, &accessory.Price,
			&accessory.Vehicletype)

		accessories = append(accessories, accessory)
	}

	return accessories

}

func GetTyresByMakeModelVehicleType(make string, model string, vehicleType constants.VehicleType) []models.Tyre {

	query := `SELECT * FROM Tyre WHERE make = ? AND model = ? AND type_of_vehicle = ?`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, make, model, vehicleType)

	checkErr(err)

	var tyres []models.Tyre

	defer rows.Close()
	for rows.Next() {

		var tyre models.Tyre

		rows.Scan(&tyre.Tyreid, &tyre.TyreType,
			&tyre.Size, &tyre.Description, &tyre.Make,
			&tyre.Model, &tyre.Price,
			&tyre.Vehicletype)

		tyres = append(tyres, tyre)
	}

	return tyres

}

func GetTransmissionsByMakeModelSubmodelVehicleType(make string, model string, submodel string, vehicleType constants.VehicleType) []models.Transmission {

	query := `SELECT * FROM Transmission WHERE make = ? AND model = ? AND submodel = ? AND type_of_vehicle = ?`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, make, model, submodel, vehicleType)

	checkErr(err)

	var transmissions []models.Transmission

	defer rows.Close()
	for rows.Next() {

		var transmission models.Transmission

		rows.Scan(&transmission.TransmissionId, &transmission.TransmissionType,
			&transmission.Description, &transmission.Make,
			&transmission.Model, &transmission.Submodel, &transmission.Price,
			&transmission.Vehicletype)

		transmissions = append(transmissions, transmission)
	}

	return transmissions

}

func GetRooftopsByMakeModelVehicleType(make string, model string, vehicleType constants.VehicleType) []models.Rooftop {

	query := `SELECT * FROM Rooftop WHERE make = ? AND model = ? AND type_of_vehicle = ?`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, make, model, vehicleType)

	checkErr(err)

	var rooftops []models.Rooftop

	defer rows.Close()
	for rows.Next() {

		var rooftop models.Rooftop

		rows.Scan(&rooftop.RooftopId, &rooftop.RooftopType,
			&rooftop.Description, &rooftop.Make,
			&rooftop.Model, &rooftop.Price,
			&rooftop.Vehicletype)

		rooftops = append(rooftops, rooftop)
	}

	return rooftops

}

func GetEnginesByMakeVehicleType(make string, vehicleType constants.VehicleType) []models.Engine {

	query := `SELECT * FROM Engine WHERE make = ? AND type_of_vehicle = ?`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, make, vehicleType)

	checkErr(err)

	var engines []models.Engine

	defer rows.Close()
	for rows.Next() {

		var engine models.Engine

		rows.Scan(&engine.EngineId, &engine.EngineType,
			&engine.Power, &engine.Description,
			&engine.Make, &engine.Price,
			&engine.Vehicletype)

		engines = append(engines, engine)
	}

	return engines

}

func GetColorsByMakeModelVehicleType(make string, model string, vehicleType constants.VehicleType) []models.Color {

	query := `SELECT * FROM Color WHERE make = ? AND model = ? AND type_of_vehicle = ?`

	con, err := db.GetConnection()

	checkErr(err)

	rows, err := con.Query(query, make, model, vehicleType)

	checkErr(err)

	var colors []models.Color

	defer rows.Close()
	for rows.Next() {

		var color models.Color

		rows.Scan(&color.ColorId, &color.ColorName,
			&color.Make, &color.Model,
			&color.Price, &color.Vehicletype)

		colors = append(colors, color)
	}

	return colors

}

func GetEngineById(id int) models.Engine {

	query := `SELECT * FROM Engine WHERE engine_id = ? `

	con, err := db.GetConnection()

	checkErr(err)
	var engine models.Engine

	con.QueryRow(query, id).Scan(&engine.EngineId, &engine.EngineType,
		&engine.Power, &engine.Description,
		&engine.Make, &engine.Price,
		&engine.Vehicletype)

	return engine
}

func GetColorById(id int) models.Color {

	query := `SELECT * FROM Color WHERE color_id = ?`

	con, err := db.GetConnection()

	checkErr(err)

	var color models.Color

	con.QueryRow(query, id).Scan(&color.ColorId, &color.ColorName,
		&color.Make, &color.Model,
		&color.Price, &color.Vehicletype)

	return color

}

func GetRooftopById(id int) models.Rooftop {

	query := `SELECT * FROM Rooftop WHERE rooftop_id = ?`

	con, err := db.GetConnection()

	checkErr(err)

	var rooftop models.Rooftop

	con.QueryRow(query, id).Scan(&rooftop.RooftopId, &rooftop.RooftopType,
		&rooftop.Description, &rooftop.Make,
		&rooftop.Model, &rooftop.Price,
		&rooftop.Vehicletype)

	return rooftop
}

func GetTransmissionById(id int) models.Transmission {

	query := `SELECT * FROM Transmission WHERE transmission_id = ?`

	con, err := db.GetConnection()

	checkErr(err)

	var transmission models.Transmission

	con.QueryRow(query, id).Scan(&transmission.TransmissionId, &transmission.TransmissionType,
		&transmission.Description, &transmission.Make,
		&transmission.Model, &transmission.Submodel, &transmission.Price,
		&transmission.Vehicletype)

	return transmission

}

func GetTyreById(id int) models.Tyre {

	query := `SELECT * FROM Tyre WHERE tyre_id = ?`

	con, err := db.GetConnection()

	checkErr(err)

	var tyre models.Tyre

	con.QueryRow(query, id).Scan(&tyre.Tyreid, &tyre.TyreType,
		&tyre.Size, &tyre.Description, &tyre.Make,
		&tyre.Model, &tyre.Price,
		&tyre.Vehicletype)

	return tyre

}
