# Elevators (working title)

A highly-available elevator control system for 16 elevators, running on Mesos.

## Installation

```
go get github.com/betawaffle/elevators
```

## Usage

```
elevators -logtostderr -stderrthreshold INFO
```

## Algorithm

### Deciding which floors to stop at

When someone presses a button in the elevator, the pressed floor number should
be added to a sorted set of target floor numbers. Each car has an associated
momentum. The car should stop at each floor in target set of floors until there
are no more in the direction of momentum. When the number of target floors in
the direction of momentum is empty, the direction of momentum should be
re-evaluated to either: stop if there are no target floors remaining in any
direction; reverse direction if there are floors remaining.

With this system, passengers can press buttons for intermediate floors, and the
elevator will stop at them on the way to the ultimate floor. Non-intermediate
floors will be arrived at after the ultimate floor has been reached.

### Deciding which elevator to call

When an elevator is called, the control system should calculate a rough
time-to-arrive for each elevator, and select the one with the lowest value.

#### Calculating time-to-arrive

If the elevator's momentum is opposite the callers desired direction, the
calculation should include the time for the elevator to go passed the caller and
come back, to minimize the time the caller would need to be in the elevator.

Each skipped floor counts as 1, and each stopped-on floor counts as 4 (stop,
align, open, close). See the elevator movement section for more information on
these numbers.

### Elevator Movement

This simple implementation will consider a simple tick-based flow of time, with
somewhat instantaneous acceleration and deceleration.

A moving elevator takes 1 tick to move from one floor to another.
A moving elevator takes 4 ticks to stop on a floor:

 - 1 tick to arrive from previous floor and stop
 - 1 tick to align to the floor
 - 1 tick to open the doors
 - 1 tick to close the doors and start moving

### Example

 - Elevator 1 on floor 7, stopping at 4, 2, and 1
 - Elevator 2 on floor 7, stopping at 6, 4, and 2
 - Elevator 3 on floor 10, stopping at 13, 15, and 16
 - Caller 1 on floor 3 wanting to travel up
 - Caller 2 on floor 9 wanting to travel up

Elevator 1 for caller 1: 6,5,4,4,4,4,3,2,2,2,2,1,1,1,1:2,3,3,3,3 = 15 + 5 = 20
Elevator 2 for caller 1: 6,6,6,6,5,4,4,4,4,3,2,2,2,2:3,3,3,3 = 14 + 4 = 18
Elevator 3 for caller 2: 10,11,12,13,13,13,13,14,15,15,15,15,16,16,16,16:15,14,13,12,11,10,9,8,7,6,5,4,3,3,3,3 = 16 + 16 = 32

Elevator 1 for caller 2: 6,5,4,4,4,4,3,2,2,2,2,1,1,1,1:2,3,4,5,6,7,8,9,9,9,9 = 15 + 11 = 26
Elevator 2 for caller 2: 6,6,6,6,5,4,4,4,4,3,2,2,2,2:3,4,5,6,7,8,9,9,9,9 = 14 + 10 = 24
Elevator 3 for caller 2: 10,11,12,13,13,13,13,14,15,15,15,15,16,16,16,16:15,14,13,12,11,10,9,9,9,9 = 16 + 10 = 26

Elevator 2 for caller 1 and 2: 6,6,6,6,5,4,4,4,4,3,2,2,2,2:3,3,3,3,4,5,6,7,8,9,9,9,9 = 14 + 3 + 10 = 27

Elevator 2 would be selected to stop at floor 3 and floor 9, which is probably
optimal for saving elevator movement, but not for the seconds caller's waiting
time. This shows something is missing from my algorithm. Shoot.

If we applied caller 1 to the elevators before caller 2, we would get a
different outcome: either elevator 1 or 3 depending on what we use to break the
tie.
