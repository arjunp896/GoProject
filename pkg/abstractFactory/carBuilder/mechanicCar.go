package carBuilder

import "goproject/pkg/models"

type MechanicCar struct {
	builderCar IBuilderCar
}

func NewMechanicCar(b IBuilderCar) *MechanicCar {

	return &MechanicCar{
		builderCar: b,
	}
}

func (m *MechanicCar) SetBuilder(b IBuilderCar) {
	m.builderCar = b
}

func (m *MechanicCar) MakeCar(data map[string]interface{}) models.CustomCar {

	m.builderCar.SetVehicle(data["vehicle"].(models.Vehicle))
	m.builderCar.SetAccessories(data["accessories"].([]models.Accessory))
	m.builderCar.SetEngine(data["engine"].(models.Engine))
	m.builderCar.SetColor(data["color"].(models.Color))
	m.builderCar.SetRooftop(data["rooftop"].(models.Rooftop))
	m.builderCar.SetTransmission(data["transmission"].(models.Transmission))
	m.builderCar.SetUser(data["user"].(models.User))

	return m.builderCar.GetCar()
}
