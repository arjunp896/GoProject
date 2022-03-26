package models

import (
	"goproject/pkg/constants"
)

type Vehicle struct {
	VehicleId   int32                 `json:"id" db:"vehicle_id"`
	Make        string                `json:"make" db:"make"`
	Model       string                `json:"model" db:"model"`
	Submodel    string                `json:"submodel" db:"submodel"`
	YearOfMfg   string                `json:"yearofmfg" db:"year_of_mfg"`
	Price       string                `json:"price" db:"price"`
	ModelType   string                `json:"modeltype" db:"model_type"`
	Category    string                `json:"category" db:"category"`
	VehicleType constants.VehicleType `json:"vehicletype" db:"type_of_vehicle"`
	ImageUrl    string                `json:"imageurl" db:"image_url"`
}
