package abstractFactory

import "goproject/pkg/models"

type Bike struct {
	models.Vehicle
}

type IBikeBuilder interface{}
