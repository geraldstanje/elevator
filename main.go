package main

import (
	"bufio"
	"elevator"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StringToInt(value string) int {
	result, _ := strconv.ParseInt(value, 10, 64)
	return int(result)
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
	Pickup
	Step
	Exit
	Error
)

func formatCmd(line string) (control_state, []string) {
	if line == "status" {
		return Status, nil
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

	ecs := elevator.NewElevatorControlSystem(*flagNumElev)

	for {
		line := readFromStdin()
		cmd, args := formatCmd(line)

		switch cmd {
		case Status:
			stat := ecs.Status()
			for _, e := range stat {
				fmt.Println(e)
			}
		case Pickup:
			ecs.Pickup(StringToInt(args[0]), StringToInt(args[1]))
		case Step:
			ecs.Step()
		case Exit:
			return
		case Error:
			fmt.Println("invalid cmd")
		}
	}
}
