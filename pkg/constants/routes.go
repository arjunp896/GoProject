package constants

type Route string

const (
	INDEX_ROUTE          Route = "/"
	LOGIN_ROUTE          Route = "/login"
	SIGNUP_ROUTE         Route = "/signup"
	GET_CATEGORIES_ROUTE Route = "/{vehicle}/getcategories"
	GET_MAKES_ROUTE      Route = "/{vehicle}/getmakes/{category}"
	GET_YEARS_ROUTE      Route = "/{vehicle}/getyears/{category}/{make}"
	DASHBOARD_ROUTE      Route = "/dashboard"

	GET_ALL_VEHICLES_ROUTE Route = "/vehicles"

	GET_ALL_CARS_ROUTE Route = "/car/cars/"
	CAR_DETAILS_ROUTE  Route = "/car/{id}"

	BIKE_DETAILS_ROUTE  Route = "/bike/bikes/"
	GET_ALL_BIKES_ROUTE Route = "/bike/{id}"

	GET_CARS_BY_CATEGORY           Route = "/car/cars/{category}"
	GET_CARS_BY_CATEGORY_MAKE      Route = "/car/cars/{category}/{make}"
	GET_CARS_BY_CATEGORY_MAKE_YEAR Route = "/car/cars/{category}/{make}/{year}"

	GET_BIKES_BY_CATEGORY           Route = "/bike/bikes/{category}"
	GET_BIKES_BY_CATEGORY_MAKE      Route = "/bike/bikes/{category}/{make}"
	GET_BIKES_BY_CATEGORY_MAKE_YEAR Route = "/bike/bikes/{category}/{make}/{year}"
)
