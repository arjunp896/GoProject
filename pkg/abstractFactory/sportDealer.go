package abstractFactory

import (
	"goproject/pkg/abstractFactory/carBuilder"
	"goproject/pkg/constants"
)

type SportDealer struct{}

func (s SportDealer) MakeCar() carBuilder.Car {

	sportBuilder := carBuilder.GetBuilder(constants.Sport)

	mechanic := carBuilder.NewMechanicCar(sportBuilder)

	sportCar := mechanic.MakeCar()

	return sportCar
}

func (s SportDealer) MakeBike() IBikeBuilder {

	// sportBikeBuilder :=

	return &Bike{}
}
