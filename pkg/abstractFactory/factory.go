package abstractFactory

import (
	"goproject/pkg/abstractFactory/carBuilder"
	"goproject/pkg/constants"
	"log"
)

type IVehicleFactory interface {
	MakeCar() carBuilder.Car
	MakeBike() IBikeBuilder
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
