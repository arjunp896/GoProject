package abstractFactory

import (
	"goproject/pkg/constants"
	"goproject/pkg/models"
	"log"
)

type IVehicleFactory interface {
	MakeCar(map[string]interface{}) models.CustomCar
	MakeBike() models.CustomBike
}

var mapFactory = make(map[constants.Category]IVehicleFactory)

func GetVehicleFactory(category constants.Category) IVehicleFactory {

	if _, ok := mapFactory[category]; ok {

		return mapFactory[category]
	} else {

		var factory IVehicleFactory = nil

		switch category {

		case constants.Sport:
			factory = &SportDealer{}

		case constants.Cruiser:

		case constants.Electric:

		case constants.Offroad:

		case constants.Racing:

		case constants.Touring:

		default:
			log.Fatal("unsupported car type")
		}

		mapFactory[category] = factory

		return factory
	}
}
