package elevator

type Direction int

const (
	GoingUp   Direction = 1
	Stopped             = 0
	GoingDown           = -1
)

// Set is a collection of elevators managed by a single control system.
type Set []Car

// PingEvent is an object that holds details of the last known location and
// movement of an elevator
// type PingEvent struct {
//   Time time.Time
//   Floor int
//   Direction ElevatorDirection
// }

type PickupRequest struct {
	// PickupFloor is the floor index of the requests origin (requester of the elevator).
	PickupFloor int

	// DesiredDirection is the direction the requester wants to travel.
	// In most cases, this is just a hint, not a mandate, since the person can
	// press a lower floor number, or not enter the car at all.
	DesiredDirection Direction
}

type TravelRequest struct {
	// DesiredFloor is the floor index of the car occupant's desired destination.
	DesiredFloor int
}
