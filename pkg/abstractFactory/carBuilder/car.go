package carBuilder

import (
	"fmt"
	"goproject/pkg/models"
)

type Car struct {
	models.Vehicle
	Doors   int    `json:"dorrs"`
	Rooftop string `json:"rooftop"`
}

type ICarBuilder interface {
	BuildCar()
}

func (c *Car) BuildCar() {

	fmt.Print("This is new car")
}
