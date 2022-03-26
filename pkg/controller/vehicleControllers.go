package controller

import (
	"encoding/json"
	"goproject/pkg/service"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {

	vehicles := service.GetCars()

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(vehicles)

}

func GetCar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	car := service.GetCar(params["id"])

	json.NewEncoder(w).Encode(car)

}
