package constants

type VehicleType string

type Category string

const (
	Car  VehicleType = "car"
	Bike VehicleType = "bike"

	Sport    Category = "sport"
	Offroad  Category = "offroad"
	Racing   Category = "racing"
	Touring  Category = "touring"
	Cruiser  Category = "cruiser"
	Electric Category = "electric"

	USER_COOKIE_NAME string = "user-cookie"

	AUTHENTICATED_COOKIE_VALUES string = "authenticated"
	USERID_COOKIE_VALUES        string = "userid"
)
