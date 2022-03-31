package carBuilder

import "goproject/pkg/models"

type SportBuilderCar struct {
	// vehicle models.Vehicle
	Vehicle      models.Vehicle
	User         models.User
	TotalPrice   float64
	Engine       models.Engine
	Transmission models.Transmission
	Tyre         models.Tyre
	Color        models.Color
	Accessories  []models.Accessory
	Rooftop      models.Rooftop
}

func (sbc *SportBuilderCar) GetCar() models.CustomCar {
	return models.CustomCar{

		CustomVehicle: models.CustomVehicle{
			Vehicle: sbc.Vehicle,
		},
	}
}

func (sbc *SportBuilderCar) SetVehicle(vehicle models.Vehicle) {

	sbc.Vehicle = vehicle
}

func (sbc *SportBuilderCar) SetUser(user models.User) {

	sbc.User = user
}

func (sbc *SportBuilderCar) SetEngine(engine models.Engine) {

	sbc.Engine = engine
}

func (sbc *SportBuilderCar) SetTransmission(trans models.Transmission) {

	sbc.Transmission = trans
}

func (sbc *SportBuilderCar) SetAccessories(accessories []models.Accessory) {

	sbc.Accessories = accessories
}

func (sbc *SportBuilderCar) SetColor(color models.Color) {

	sbc.Color = color
}

func (sbc *SportBuilderCar) SetTyre(tyre models.Tyre) {

	sbc.Tyre = tyre
}

func (sbc *SportBuilderCar) SetRooftop(rooftop models.Rooftop) {

	sbc.Rooftop = rooftop
}
