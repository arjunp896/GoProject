package controller

import (
	"goproject/config"
	"goproject/pkg/constants"
	"goproject/pkg/service"
	"net/http"
)

func CreateNewCustomVehicle(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	mapData := make(map[string]interface{})

	mapData["vehicleid"] = r.Form["vehicleid"][0]
	mapData["colorid"] = r.Form["colorid"][0]
	mapData["engineid"] = r.Form["engineid"][0]
	mapData["rooftopid"] = r.Form["rooftopid"][0]
	mapData["transmissionid"] = r.Form["transmissionid"][0]
	mapData["tyreid"] = r.Form["tyreid"][0]
	mapData["accessoriesid"] = r.Form["accessoriesid"]

	session, _ := config.GetSession(r, constants.USER_COOKIE_NAME)
	mapData["userid"] = session.Values[constants.USERID_COOKIE_VALUES].(int)

	service.CreateNewCustomVehicle(mapData)

	renderTemplate(w, "./web/vehicles.html")

}
