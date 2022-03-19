package controllers

import (
	"encoding/json"
	"fmt"
	"goproject/pkg/services"
	"net/http"

	"github.com/gorilla/mux"
)

func Index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Jai Swaminarayan from goproject")
}

func GetCars(w http.ResponseWriter, r *http.Request) {

	// Mock data @todo - implement DB

	cars := services.GetCars()

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(cars)

}

func GetCar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	car := services.GetCar(params["id"])

	json.NewEncoder(w).Encode(car)

}
