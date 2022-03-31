package bikeBuilder

import "goproject/pkg/models"

type SportBuilderBike struct {
	Vehicle      models.Vehicle
	User         models.User
	TotalPrice   float64
	Engine       models.Engine
	Transmission models.Transmission
	Tyre         models.Tyre
	Color        models.Color
	Accessories  []models.Accessory
}

func (sbc *SportBuilderBike) GetBike() models.CustomBike {
	return models.CustomBike{

		CustomVehicle: models.CustomVehicle{
			Vehicle: sbc.Vehicle,
		},
	}
}

func (sbc *SportBuilderBike) SetVehicle(vehicle models.Vehicle) {

	sbc.Vehicle = vehicle
}

func (sbc *SportBuilderBike) SetUser(user models.User) {

	sbc.User = user
}

func (sbc *SportBuilderBike) SetEngine(engine models.Engine) {

	sbc.Engine = engine
}

func (sbc *SportBuilderBike) SetTransmission(trans models.Transmission) {

	sbc.Transmission = trans
}

func (sbc *SportBuilderBike) SetAccessories(accessories []models.Accessory) {

	sbc.Accessories = accessories
}

func (sbc *SportBuilderBike) SetColor(color models.Color) {

	sbc.Color = color
}

func (sbc *SportBuilderBike) SetTyre(tyre models.Tyre) {

	sbc.Tyre = tyre
}
