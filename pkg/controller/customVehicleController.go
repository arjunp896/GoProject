package controller

import (
	"goproject/pkg/service"
	"net/http"
)

func CreateNewCustomVehicle(w http.ResponseWriter, r *http.Request) {

	service.CreateNewCustomVehicle()

}
