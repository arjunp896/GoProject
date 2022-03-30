package carBuilder

type MechanicCar struct {
	builderCar IBuilderCar
}

func NewMechanicCar(b IBuilderCar) *MechanicCar {

	return &MechanicCar{
		builderCar: b,
	}
}

func (d *MechanicCar) SetBuilder(b IBuilderCar) {
	d.builderCar = b
}

func (d *MechanicCar) MakeCar() Car {

	return d.builderCar.getCar()
}
