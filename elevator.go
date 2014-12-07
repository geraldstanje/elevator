package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
  "queue"
  "strings"
  "strconv"
)

type Elevator struct {
	elevatorID        int
  currentFloorNumber int
	direction int     // -1 == down, 0 == nowhere, 1 == up
	goalFloorNumber   map[int]bool
}

func StringToInt(value string) int {
  result, _ := strconv.ParseInt(value, 10, 64)
  return int(result)
}

func NewElevator(ID int) *Elevator {
	e := Elevator{elevatorID: ID}
	return &e
}

func (e *Elevator) addGoalFloor(floorNumber int, dir int) {
  if len(e.goalFloorNumber[floorNumber]) == 0 {
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
  } else if e.direction > 0 {
    return e.currentFloorNumber + 1
  }
}

type pickup struct {
  pickupFloor int
  direction int // -1 == down, 1 == up
}

type ElevatorControlSystem struct {
	elevator []*Elevator
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

func (ecs *ElevatorControlSystem) status() string {
  return ""
}

func (ecs *ElevatorControlSystem) pickup(pickupFloorNumber int, direction int) {
  ecs.pickupRequests.Push(pickup{pickupFloorNumber, direction})
}

func (ecs *ElevatorControlSystem) update(elevatorID int, currentFloor int, goalFloor int, direction int) {
  ecs.elevator[elevatorID].Update(currentFloor, goalFloor, direction)
}

func (ecs *ElevatorControlSystem) step() {
  for e := range ecs.elevator {
    if ecs.pickupRequests.Len() > 0 {
      req := ecs.pickupRequests.Pop()
      ecs.update(i, ecs.elevator[elevatorID].NextFloor(), req.pickupFloor, req.direction)
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

var flagNumElev = flag.Int("n", 0, "intflag")

func main() {
	flag.Parse()
	if *flagNumElev <= 0 {
		fmt.Printf("Usage: %s -n NumberOfElevators\n", os.Args[0])
		os.Exit(1)
	}

	ecs := NewElevatorControlSystem(*flagNumElev)

	for {
		cmd := readFromStdin()

    if(strings.HasPrefix(cmd, "exit")) {
      return
    } else if (strings.HasPrefix(cmd, "pickup")) {
      cmd = strings.Trim(cmd, "pickup ")
      args := strings.Split(cmd, " ")

      if len(args) == 2 {
        ecs.pickup(StringToInt(args[0]), StringToInt(args[1]))
      }
    } else if (strings.HasPrefix(cmd, "step")) {
      ecs.step()
    }
	}

	fmt.Println(ecs)
}
