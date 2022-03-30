package models

import "goproject/pkg/constants"

type CustomVehicle struct {
	CVId         int          `json:"cv_id"`
	UserId       int          `json:"user_id"`
	Vehicle      Vehicle      `json:"vehicle"`
	TotalPrice   float64      `json:"total_price"`
	Engine       Engine       `json:"engine"`
	Transmission Transmission `json:"transmission"`
	Tyre         Tyre         `json:"tyre"`
	Rooftop      Rooftop      `json:"rooftop"`
	Color        Color        `json:"color"`
	Accessories  []Accessory  `json:"accessories"`
}

type Engine struct {
	EngineId    int                   `json:"engine_id"`
	EngineType  string                `json:"engine_type"`
	Power       string                `json:"power"`
	Description string                `json:"description"`
	Make        string                `json:"make"`
	Price       float64               `json:"price"`
	Vehicletype constants.VehicleType `json:"vehicle_type"`
}

type Transmission struct {
	TransmissionId   int                   `json:"transmission_id"`
	TransmissionType string                `json:"transmission_type"`
	Description      string                `json:"description"`
	Make             string                `json:"make"`
	Model            string                `json:"model"`
	Submodel         string                `json:"submodel"`
	Price            string                `json:"price"`
	Vehicletype      constants.VehicleType `json:"vehicle_type"`
}

type Tyre struct {
	Tyreid      int                   `json:"tyre_id"`
	TyreType    string                `json:"type_of_tyre"`
	Size        string                `json:"size"`
	Description string                `json:"description"`
	Make        string                `json:"make"`
	Model       string                `json:"model"`
	Price       string                `json:"price"`
	Vehicletype constants.VehicleType `json:"vehicle_type"`
}

type Rooftop struct {
	RooftopId   int                   `json:"rooftop_id"`
	RooftopType string                `json:"rooftop_type"`
	Description string                `json:"description"`
	Make        string                `json:"make"`
	Model       string                `json:"model"`
	Price       float64               `json:"price"`
	Vehicletype constants.VehicleType `json:"vehicle_type"`
}

type Color struct {
	ColorId     int                   `json:"color_id"`
	ColorName   string                `json:"color_name"`
	Make        string                `json:"make"`
	Model       string                `json:"model"`
	Price       float64               `json:"price"`
	Vehicletype constants.VehicleType `json:"vehicle_type"`
}

type Accessory struct {
	AccessoryId   int                   `json:"accessory_id"`
	AccessoryName string                `json:"accessories_name"`
	Description   string                `json:"description"`
	Make          string                `json:"make"`
	Model         string                `json:"model"`
	Price         float64               `json:"price"`
	Vehicletype   constants.VehicleType `json:"vehicle_type"`
}
