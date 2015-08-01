package elevator

import (
	"sort"
	"time"
)

// Car is the state of a single elevator car.
type Car struct {
	// LastPing is the time of the last status update from the elevator.
	LastPing time.Time

	// Floor is the index of the floor the elevator was at or passing as of LastPing.
	Floor int

	// Direction is the direction the elevator was moving as of LastPing.
	Direction Direction

	// Momentum is used to keep the elevator moving is the same direction to
	// preserve fairness between the occupants. The elevator should keep moving in
	// the same direction until it reaches the highest or lowest requested floor,
	// before going in the other direction to fulfill remaining requests.
	Momentum Direction

	// Aligned indicates whether the elevator was aligned with a floor, ready for
	// the doors to open, as of LastPing.
	Aligned bool

	// DoorsOpen indicates whether the elevator's doors were open as of LastPing.
	DoorsOpen bool

	// DesiredFloors is a set of floor indexes for the currently lit buttons.
	DesiredFloors sort.IntSlice
}

func (car *Car) step() {
	// Update the floor based on the current movement.
	car.Floor = car.Floor + int(car.Direction)

	if !car.shouldStop() {
		return
	}

	if car.Direction != Stopped {
		// stop the car, but preserve the momentum
		car.Momentum, car.Direction = car.Direction, Stopped
		return // delay alignment by one tick
	}
	if !car.Aligned {
		car.Aligned = true
		return // delay opening the doors by one tick
	}

	// decide if we should stop, keep going, or change direction
	car.checkMomentum()

	if !car.DoorsOpen {
		car.DoorsOpen = true
		return // delay closing the doors by one tick
	}

	car.DoorsOpen = false
	car.Direction = car.Momentum // car will start moving on the next tick
}

func (car *Car) shouldStop() bool {
	// Is the car at one of the desired floors?
	return car.DesiredFloors[car.DesiredFloors.Search(car.Floor)] == car.Floor
}

func (car *Car) checkMomentum() {
	i := car.DesiredFloors.Search(car.Floor)
}
