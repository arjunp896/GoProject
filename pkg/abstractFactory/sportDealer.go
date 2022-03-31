package abstractFactory

import (
	"goproject/pkg/abstractFactory/carBuilder"
	"goproject/pkg/constants"
	"goproject/pkg/models"
)

type SportDealer struct{}

func (s SportDealer) MakeCar(data map[string]interface{}) models.CustomCar {

	sportBuilder := carBuilder.GetBuilder(constants.Sport)

	mechanic := carBuilder.NewMechanicCar(sportBuilder)

	sportCar := mechanic.MakeCar(data)

	return sportCar
}

func (s SportDealer) MakeBike() models.CustomBike {

	// sportBikeBuilder :=

	// sportBuilder := carBuilder.GetBuilder(constants.Sport)

	// mechanic := carBuilder.NewMechanicBike(sportBuilder)

	return models.CustomBike{}
}
