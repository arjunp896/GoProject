package abstractFactory

import (
	"goproject/pkg/constants"
	"log"
)

type IAbstractFactory interface {
	MakeCar() ICar
	MakeBike() IBike
}

var mapFactory = make(map[constants.Category]IAbstractFactory)

func GetFactory(category constants.Category) IAbstractFactory {

	if _, ok := mapFactory[category]; ok {

		return mapFactory[category]
	} else {

		var factory IAbstractFactory = nil

		switch category {

		case constants.Sport:
			factory = &Sport{}

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
