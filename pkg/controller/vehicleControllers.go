package controller

import (
	"encoding/json"
	"goproject/pkg/models"
	"goproject/pkg/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetVehicles(w http.ResponseWriter, r *http.Request) {

	cars := service.GetCars()
	bikes := service.GetBikes()

	vehicles := append(cars, bikes...)

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(vehicles)

}

func GetCars(w http.ResponseWriter, r *http.Request) {

	cars := service.GetCars()

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(cars)

}

func GetCar(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	data := service.GetVehicleDetailsPageData(id)

	// json.NewEncoder(w).Encode(car)

	// carJsn := []byte(fmt.Sprintf("%v", car))

	renderTemplateWithData(w, string(http.Dir("./web/vehicle_details.html")), data)

}

func GetBikes(w http.ResponseWriter, r *http.Request) {

	bikes := service.GetBikes()

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(bikes)

}

func GetBike(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	data := service.GetVehicleDetailsPageData(id)

	// json.NewEncoder(w).Encode(car)

	// carJsn := []byte(fmt.Sprintf("%v", car))

	renderTemplateWithData(w, string(http.Dir("./web/vehicle_details.html")), data)

}

func GetCarsByFilter(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var cars []models.Vehicle

	switch len(params) {
	case 1:
		cars = service.GetCarsByFilter(params["category"])

	case 2:
		cars = service.GetCarsByFilter(params["category"], params["make"])

	case 3:
		cars = service.GetCarsByFilter(params["category"], params["make"], params["year"])

	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(cars)
}

func GetBikesByFilter(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var bikes []models.Vehicle

	switch len(params) {
	case 1:
		bikes = service.GetBikesByFilter(params["category"])

	case 2:
		bikes = service.GetBikesByFilter(params["category"], params["make"])

	case 3:
		bikes = service.GetBikesByFilter(params["category"], params["make"], params["year"])

	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(bikes)

}
