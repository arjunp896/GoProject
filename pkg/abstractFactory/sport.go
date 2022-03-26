package abstractFactory

type Sport struct{}

func (s *Sport) MakeBike() IBike {

	return &Bike{}
}

func (s *Sport) MakeCar() ICar {

	return &Car{}
}
