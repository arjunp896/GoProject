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

	r.HandleFunc("/car/cars/", controller.GetCars).Methods("GET")
	r.HandleFunc("/car/{id}", controller.GetCar).Methods("GET")

	// load index page
	r.Handle(string(constants.INDEX_ROUTE), http.FileServer(http.Dir("./web/")))

	// load login page
	r.HandleFunc(string(constants.LOGIN_ROUTE), controller.LoadLoginPage).Methods("GET")

	// validate user for login
	r.HandleFunc(string(constants.LOGIN_ROUTE), controller.ValidateUser).Methods("POST")

	// create new user
	r.HandleFunc(string(constants.SIGNUP_ROUTE), controller.CreateUser).Methods("POST")

	// return all vehicle categories
	r.HandleFunc(string(constants.GET_CATEGORIES_ROUTE), controller.GetCategories)

	// return all companies by category
	r.HandleFunc(string(constants.GET_MAKES_ROUTE), controller.GetMakes)

	// retrurn all years of mfg by category and make
	r.HandleFunc(string(constants.GET_YEARS_ROUTE), controller.GetMfgYears)

	return r
}
