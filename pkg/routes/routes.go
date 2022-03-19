package routes

import (
	"goproject/pkg/constants"
	"goproject/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeHandlers() *mux.Router {

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	// init router
	r := mux.NewRouter()

	// route handler or endpoints
	r.HandleFunc("/car/cars/", controllers.GetCars).Methods("GET")
	r.HandleFunc("/car/{id}", controllers.GetCar).Methods("GET")

	r.Handle(constants.INDEX_ROUTE, http.FileServer(http.Dir("./web")))

	return r
}
