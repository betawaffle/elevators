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

	// Aligned indicates whether the elevator was aligned with a floor, ready for
	// the doors to open, as of LastPing.
	Aligned bool

	// DoorsOpen indicates whether the elevator's doors were open as of LastPing.
	DoorsOpen bool

	// DesiredFloors is a set of floor indexes for the currently lit buttons.
	DesiredFloors sort.IntSlice
}

func (car *Car) step() {
}
