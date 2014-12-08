Build the Elevator Control System:
  - go build

Start the Elevator Control System:
  - start with 2 elevators: ./main -n 2

Improvements to Scheduling:
  - The elevator moves in the same direction as long as there are goalFloorNumber stored in the map of the elevator
  - If the goalFloorNumber is empty, the elevator will go into an idle state and change the direction if
    there are requests in the opposite direction