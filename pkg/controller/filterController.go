package controller

import (
	"encoding/json"
	"goproject/pkg/service"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	categories := make(map[string][]string)

	result := service.GetCategories(params["vehicle"])

	categories["categories"] = result

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(categories)

}

func GetMakes(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	category := params["category"]
	vehicle := params["vehicle"]

	makes := make(map[string][]string)

	makes["makes"] = service.GetMakesByCategory(category, vehicle)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(makes)

}

func GetMfgYears(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	category := params["category"]
	vehicleMake := params["make"]
	vehicle := params["vehicle"]

	years := make(map[string][]string)

	years["years"] = service.GetYearByCategoryAndMake(category, vehicleMake, vehicle)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(years)

}
