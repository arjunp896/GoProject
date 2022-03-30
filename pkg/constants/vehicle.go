package constants

type VehicleType string

type Category string

const (
	Car  VehicleType = "Car"
	Bike VehicleType = "Bike"

	Sport    Category = "Sport"
	Offroad  Category = "Offroad"
	Racing   Category = "Racing"
	Touring  Category = "Touring"
	Cruiser  Category = "Cruiser"
	Electric Category = "Electric"

	USER_COOKIE_NAME string = "user-cookie"

	AUTHENTICATED_COOKIE_VALUES string = "authenticated"
	USERID_COOKIE_VALUES        string = "userid"
)
