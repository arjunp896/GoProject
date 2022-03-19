package services

import (
	"goproject/pkg/models"
)

func GetCars() []models.Car {

	cars := loadCars()

	return cars

}

func GetCar(id string) *models.Car {

	cars := loadCars()

	for _, item := range cars {

		if item.ID == id {
			return &item
		}
	}
	return &models.Car{}
}

func loadCars() []models.Car {

	var cars []models.Car

	cars = append(cars, models.Car{ID: "1", Name: "car 1", Model: &models.Model{ID: "1", Name: "model 1"}})
	cars = append(cars, models.Car{ID: "2", Name: "car 2", Model: &models.Model{ID: "2", Name: "model 2"}})
	cars = append(cars, models.Car{ID: "3", Name: "car 3", Model: &models.Model{ID: "1", Name: "model 1"}})
	cars = append(cars, models.Car{ID: "4", Name: "car 4", Model: &models.Model{ID: "2", Name: "model 2"}})

	return cars
}
