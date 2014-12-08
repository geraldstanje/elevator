package elevator

import (
	"testing"
)

func StrsEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestElevatorControlSystem(t *testing.T) {
	ecs := NewElevatorControlSystem(2)

	ecs.Pickup(2, 1)

	stat := ecs.Status()
	if !StrsEquals(stat, []string{"elevatorID: 0, currentFloor: 0, direction: 1, goalFloors: []", "elevatorID: 1, currentFloor: 0, direction: 1, goalFloors: []"}) {
		t.Errorf("Error: %v", stat)
	}

	ecs.Step()

	stat = ecs.Status()
	if !StrsEquals(stat, []string{"elevatorID: 0, currentFloor: 0, direction: 1, goalFloors: [2]", "elevatorID: 1, currentFloor: 0, direction: 1, goalFloors: []"}) {
		t.Errorf("Error: %v", stat)
	}

	ecs.Step()

	stat = ecs.Status()
	if !StrsEquals(stat, []string{"elevatorID: 0, currentFloor: 1, direction: 1, goalFloors: [2]", "elevatorID: 1, currentFloor: 0, direction: 1, goalFloors: []"}) {
		t.Errorf("Error: %v", stat)
	}

	ecs.Step()

	stat = ecs.Status()
	if !StrsEquals(stat, []string{"elevatorID: 0, currentFloor: 2, direction: 1, goalFloors: []", "elevatorID: 1, currentFloor: 0, direction: 1, goalFloors: []"}) {
		t.Errorf("Error: %v", stat)
	}

	ecs.Pickup(0, -1)

	stat = ecs.Status()
	if !StrsEquals(stat, []string{"elevatorID: 0, currentFloor: 2, direction: 1, goalFloors: []", "elevatorID: 1, currentFloor: 0, direction: 1, goalFloors: []"}) {
		t.Errorf("Error: %v", stat)
	}

	ecs.Step()

	stat = ecs.Status()
	if !StrsEquals(stat, []string{"elevatorID: 0, currentFloor: 2, direction: -1, goalFloors: [0]", "elevatorID: 1, currentFloor: 0, direction: 1, goalFloors: []"}) {
		t.Errorf("Error: %v", stat)
	}

	ecs.Step()

	stat = ecs.Status()
	if !StrsEquals(stat, []string{"elevatorID: 0, currentFloor: 1, direction: -1, goalFloors: [0]", "elevatorID: 1, currentFloor: 0, direction: 1, goalFloors: []"}) {
		t.Errorf("Error: %v", stat)
	}

	ecs.Step()

	stat = ecs.Status()
	if !StrsEquals(stat, []string{"elevatorID: 0, currentFloor: 0, direction: -1, goalFloors: []", "elevatorID: 1, currentFloor: 0, direction: 1, goalFloors: []"}) {
		t.Errorf("Error: %v", stat)
	}
}
