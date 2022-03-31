package carBuilder

import (
	"goproject/pkg/constants"
	"goproject/pkg/models"
)

type IBuilderCar interface {
	GetCar() models.CustomCar
	SetVehicle(models.Vehicle)
	SetUser(models.User)
	SetEngine(models.Engine)
	SetTransmission(models.Transmission)
	SetAccessories([]models.Accessory)
	SetColor(models.Color)
	SetTyre(models.Tyre)
	SetRooftop(models.Rooftop)
}

func GetBuilder(builderType constants.Category) IBuilderCar {
	if builderType == constants.Sport {

		return &SportBuilderCar{}
	}

	if builderType == constants.Electric {
		// return &ElectricBuilderCar{}
	}
	return nil
}
