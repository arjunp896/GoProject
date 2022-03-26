package constants

type Route string

const (
	INDEX_ROUTE          Route = "/"
	LOGIN_ROUTE          Route = "/login"
	SIGNUP_ROUTE         Route = "/signup"
	GET_CATEGORIES_ROUTE Route = "/getcategories"
	GET_MAKES_ROUTE      Route = "/getmakes/{category}"
	GET_YEARS_ROUTE      Route = "/getyears/{category}/{make}"
	DASHBOARD_ROUTE      Route = "/dashboard"
	CAR_DETAILS_ROUTE    Route = "/car/{id}"
)
