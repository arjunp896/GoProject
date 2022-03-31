package service

import (
	"fmt"
	"goproject/pkg/abstractFactory"
	"goproject/pkg/constants"
	"goproject/pkg/db"
	"goproject/pkg/models"
	"strconv"
	"strings"
)

func CreateNewCustomVehicle(data map[string]interface{}) {

	makeData := make(map[string]interface{})

	var intVar int

	intVar, _ = strconv.Atoi(data["vehicleid"].(string))

	vehicle := GetVehicleById(intVar)

	makeData["vehicle"] = vehicle

	intVar, _ = strconv.Atoi(data["engineid"].(string))

	makeData["engine"] = GetEngineById(intVar)

	var intArr = []int{}

	for _, i := range data["accessoriesid"].([]string) {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}

		intArr = append(intArr, j)
	}

	makeData["accessories"] = GetAccessoriesByIds(intArr)

	intVar, _ = strconv.Atoi(data["colorid"].(string))

	makeData["color"] = GetColorById(intVar)

	intVar, _ = strconv.Atoi(data["rooftopid"].(string))

	makeData["rooftop"] = GetRooftopById(intVar)

	intVar, _ = strconv.Atoi(data["transmissionid"].(string))

	makeData["transmission"] = GetTransmissionById(intVar)

	intVar, _ = strconv.Atoi(data["tyreid"].(string))

	makeData["tyre"] = GetTyreById(intVar)

	makeData["user"] = GetUserById(data["userid"].(int))

	category := vehicle.Category

	vehicleType := vehicle.VehicleType

	factory := abstractFactory.GetVehicleFactory(category)

	if vehicleType == constants.Car {

		car := factory.MakeCar(makeData)

		fmt.Println(car)

	}

	if vehicleType == constants.Bike {

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
