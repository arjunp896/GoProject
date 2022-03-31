package routes

import (
	"goproject/pkg/constants"
	"goproject/pkg/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeHandlers() *mux.Router {

	// init router
	r := mux.NewRouter()

	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	// ----------------- load page -----------------

	r.HandleFunc(string(constants.INDEX_ROUTE), controller.LoadLoginPage).Methods("GET")

	// load login page
	r.HandleFunc(string(constants.LOGIN_ROUTE), controller.LoadLoginPage).Methods("GET")

	// load dashboard page
	r.HandleFunc(string(constants.DASHBOARD_ROUTE), controller.LoadDashboardPage)

	// load dashboard page
	r.HandleFunc(string(constants.DASHBOARD_ROUTE), controller.LoadDashboardPage).Methods("GET")

	r.HandleFunc(string(constants.MYCARS_ROUTE), controller.LoadMyCars).Methods("GET")

	// -----------------------------------------------

	// validate user for login
	r.HandleFunc(string(constants.LOGIN_ROUTE), controller.ValidateUser).Methods("POST")

	// create new user
	r.HandleFunc(string(constants.SIGNUP_ROUTE), controller.CreateUser).Methods("POST")

	// ----------------- filter -----------------

	// return all vehicle categories
	r.HandleFunc(string(constants.GET_CATEGORIES_ROUTE), controller.GetCategories)

	// return all companies by category
	r.HandleFunc(string(constants.GET_MAKES_ROUTE), controller.GetMakes)

	// retrurn all years of mfg by category and make
	r.HandleFunc(string(constants.GET_YEARS_ROUTE), controller.GetMfgYears)

	// ------------------------------------------

	// ----------------- Get vehicle -------------
	r.HandleFunc(string(constants.GET_ALL_VEHICLES_ROUTE), controller.GetVehicles).Methods("GET")

	r.HandleFunc(string(constants.GET_ALL_CARS_ROUTE), controller.GetCars).Methods("GET")

	r.HandleFunc(string(constants.CAR_DETAILS_ROUTE), controller.GetCar).Methods("GET")

	r.HandleFunc(string(constants.GET_ALL_BIKES_ROUTE), controller.GetBikes).Methods("GET")

	r.HandleFunc(string(constants.BIKE_DETAILS_ROUTE), controller.GetBike).Methods("GET")

	// cars

	r.HandleFunc(string(constants.GET_VEHICLE_ROUTE), controller.GetVehicle).Methods("GET")

	r.HandleFunc(string(constants.GET_CARS_BY_CATEGORY), controller.GetCarsByFilter).Methods("GET")

	r.HandleFunc(string(constants.GET_CARS_BY_CATEGORY_MAKE), controller.GetCarsByFilter).Methods("GET")

	r.HandleFunc(string(constants.GET_CARS_BY_CATEGORY_MAKE_YEAR), controller.GetCarsByFilter).Methods("GET")

	// bikes

	r.HandleFunc(string(constants.GET_BIKES_BY_CATEGORY), controller.GetBikesByFilter).Methods("GET")

	r.HandleFunc(string(constants.GET_BIKES_BY_CATEGORY_MAKE), controller.GetBikesByFilter).Methods("GET")

	r.HandleFunc(string(constants.GET_BIKES_BY_CATEGORY_MAKE_YEAR), controller.GetBikesByFilter).Methods("GET")

	// ------------------------------------------

	// custom vehicle
	r.HandleFunc(string(constants.CREATE_CUSTOM_VEHICLE_ROUTE), controller.CreateNewCustomVehicle).Methods("POST")

	return r
}
