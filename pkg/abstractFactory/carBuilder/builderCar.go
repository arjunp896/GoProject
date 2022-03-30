package carBuilder

import (
	"goproject/pkg/constants"
)

type IBuilderCar interface {
	getCar() Car
}

func GetBuilder(builderType constants.Category) IBuilderCar {
	if builderType == constants.Sport {

		return &SportBuilderCar{}
	}

	if builderType == "hawaiian" {
		// return &HawaiianBuilderPizza{}
	}
	return nil
}
