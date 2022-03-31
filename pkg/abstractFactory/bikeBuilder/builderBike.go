package bikeBuilder

import (
	"goproject/pkg/constants"
	"goproject/pkg/models"
)

type IBuilderBike interface {
	GetBike() models.CustomBike
}

func GetBuilder(builderType constants.Category) IBuilderBike {
	if builderType == constants.Sport {

		return &SportBuilderBike{}
	}

	// if builderType == constants.Electric {
	// return &ElectricBuilderCar{}
	// }
	return nil
}
