package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"queue"
	"strconv"
	"strings"
)

type Elevator struct {
	elevatorID         int
	currentFloorNumber int
	direction          int // -1 == down, 1 == up
	goalFloorNumber    map[int]bool
}

func StringToInt(value string) int {
	result, _ := strconv.ParseInt(value, 10, 64)
	return int(result)
}

func NewElevator(ID int) *Elevator {
	e := Elevator{elevatorID: ID, direction: 1}
	return &e
}

func (e *Elevator) GetElevatorID() int {
	return e.elevatorID
}

func (e *Elevator) GetCurrentFloorNumber() int {
  return e.currentFloorNumber
}

func (e *Elevator) GetGoalFloorNumbers() []int {
  goalFloors := make([]int, 0)

  for k, _ := range e.goalFloorNumber { 
    goalFloors = append(goalFloors, k)
  }

  return goalFloors
}

func (e *Elevator) addGoalFloor(floorNumber int, dir int) {
	if n := len(e.goalFloorNumber); n == 0 {
		e.direction = dir
	}
	e.goalFloorNumber[floorNumber] = true
}

func (e *Elevator) removeGloalFloor(floorNumber int) {
	delete(e.goalFloorNumber, floorNumber)
}

func (e *Elevator) Update(currentFloorNum int, goalFloorNum int, dir int) {
	e.currentFloorNumber = currentFloorNum

	e.removeGloalFloor(currentFloorNum)
	e.addGoalFloor(goalFloorNum, dir)
}

func (e *Elevator) NextFloor() int {
	if e.direction < 0 {
		return e.currentFloorNumber - 1
	}

	return e.currentFloorNumber + 1
}

type PickupReq struct {
	pickupFloor int
	direction   int // -1 == down, 1 == up
}

type ElevatorControlSystem struct {
	elevator       []*Elevator
	pickupRequests *queue.Queue
}

func NewElevatorControlSystem(NumberOfElevators int) *ElevatorControlSystem {
	ecs := ElevatorControlSystem{}

	for i := 0; i < NumberOfElevators; i++ {
		ecs.elevator = append(ecs.elevator, NewElevator(i))
	}
	ecs.pickupRequests = queue.NewQueue()

	return &ecs
}

func (ecs *ElevatorControlSystem) status() {
	for _, elev := range ecs.elevator {
    fmt.Println("elevatorID:", elev.GetElevatorID(), "currentFloor:", elev.GetCurrentFloorNumber(), "goalFloors:", elev.GetGoalFloorNumbers())
  }
}

func (ecs *ElevatorControlSystem) pickup(pickupFloorNumber int, direction int) {
	ecs.pickupRequests.Push(PickupReq{pickupFloorNumber, direction})
}

func (ecs *ElevatorControlSystem) update(elevatorID int, currentFloor int, goalFloor int, direction int) {
	ecs.elevator[elevatorID].Update(currentFloor, goalFloor, direction)
}

func (ecs *ElevatorControlSystem) step() {
	for _, elev := range ecs.elevator {
		if ecs.pickupRequests.Len() > 0 {
			req := ecs.pickupRequests.Pop()

			if e, ok := req.(PickupReq); ok {
				id := elev.GetElevatorID()
				ecs.update(id, ecs.elevator[id].NextFloor(), e.pickupFloor, e.direction)
			}
		} else {
			continue
		}
	}
}

func readFromStdin() string {
	fmt.Print("$ ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = line[0 : len(line)-1]
	return line
}

type control_state int

const (
	Status control_state = iota
	Update
	Pickup
	Step
	Exit
	Error
)

func formatCmd(line string) (control_state, []string) {
	if line == "status" {
		return Status, nil
	} else if line == "update" {
		return Update, nil
	} else if line == "step" {
		return Step, nil
	} else if line == "exit" {
		return Exit, nil
	} else if strings.HasPrefix(line, "pickup") {
		line = strings.Trim(line, "pickup ")
		args := strings.Split(line, " ")
		if len(args) == 2 {
			return Pickup, args
		}
	}

	return Error, nil
}

var flagNumElev = flag.Int("n", 0, "intflag")

func main() {
	flag.Parse()
	if *flagNumElev <= 0 {
		fmt.Printf("Usage: %s -n NumberOfElevators\n", os.Args[0])
		os.Exit(1)
	}

	ecs := NewElevatorControlSystem(*flagNumElev)

	for {
		line := readFromStdin()
		cmd, args := formatCmd(line)

		switch cmd {
		case Status:
			ecs.status()
		case Update:
			continue
		case Pickup:
			ecs.pickup(StringToInt(args[0]), StringToInt(args[1]))
		case Step:
			ecs.step()
		case Exit:
			return
		case Error:
			fmt.Println("invalid cmd")
		}
	}
}
