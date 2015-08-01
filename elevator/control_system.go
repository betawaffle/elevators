package elevator

type ControlSystem struct {
	cars Set
}

func NewControlSystem(numCars int) *ControlSystem {
	return &ControlSystem{
		cars: make(Set, numCars),
	}
}

func (cs *ControlSystem) step() {
	for i := range cs.cars {
		car := &cs.cars[i]
		car.step()
	}
}
