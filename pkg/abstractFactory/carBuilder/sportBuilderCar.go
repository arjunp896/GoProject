package carBuilder

import "goproject/pkg/models"

type SportBuilderCar struct {
	vehicle models.Vehicle
}

func (sbc *SportBuilderCar) getCar() Car {
	return Car{
		Vehicle: sbc.vehicle,
	}
}

func (b *SportBuilderCar) setVehicle(vehicle models.Vehicle) {
	b.vehicle = vehicle
}
