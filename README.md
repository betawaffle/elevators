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

TODO
