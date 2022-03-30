package service

import (
	"fmt"
	"goproject/pkg/abstractFactory"
	"goproject/pkg/constants"
	"goproject/pkg/db"
	"goproject/pkg/models"
	"strings"
)

func CreateNewCustomVehicle() {
	category := constants.Sport

	factory := abstractFactory.GetVehicleFactory(category)

	vehicleType := constants.Car

	vehicleId := 1

	accessories := []int{1, 2, 3}

	engineId := 1

	m := make(map[string]interface{})

	m["engine"] = engineId

	m["rooftop"] = "1"

	m["transmission"] = "1"

	m["tyres"] = "1"

	m["accessories"] = GetAccessoriesByIds(accessories)

	vehicle := GetVehicleById(vehicleId)

	m["vehicle"] = vehicle

	if vehicleType == constants.Car {

		car := factory.MakeCar()

		fmt.Println(car)

	} else {
		car := factory.MakeCar()

		car.BuildCar()
	}
}

func GetAccessoriesByIds(accessorieIds []int) []models.Accessory {

	query := `SELECT * FROM Accessories WHERE accessories_id IN (?)`

	con, err := db.GetConnection()

	checkErr(err)

	s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(accessorieIds)), ","), "[]")

	fmt.Println("\n", s)

	rows, err := con.Query(query, s)

	checkErr(err)

	var accessories []models.Accessory

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
