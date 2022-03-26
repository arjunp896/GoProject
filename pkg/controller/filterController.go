package controller

import (
	"encoding/json"
	"goproject/pkg/service"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {

	result := service.GetCategories()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)

}

func GetMakes(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	category := params["category"]

	makes := service.GetMakesByCategory(category)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(makes)

}

func GetMfgYears(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	category := params["category"]
	make := params["make"]

	years := service.GetYearByCategoryAndMake(category, make)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(years)

}
