### Problem Specification

Design and implement an elevator control system. What data structures,
interfaces and algorithms will you need? Your elevator control system should
be able to handle a few elevators -- up to 16.

You can use the language of your choice to implement an elevator control
system. In the end, your control system should provide an interface for:

  * Querying the state of the elevators (what floor are they on and where they
    are going),

  * receiving an update about the status of an elevator,

  * receiving a pickup request,

  * time-stepping the simulation.

For example, we could imagine in Scala an interface like this:

```scala
  trait ElevatorControlSystem {
    def status(): Seq[(Int, Int, Int)]
    def update(Int, Int, Int)
    def pickup(Int, Int)
    def step()
  }
```

Here we have chosen to represent elevator state as 3 integers:

  Elevator ID, Floor Number, Goal Floor Number

An update alters these numbers for one elevator. A pickup request is two
integers:

  Pickup Floor, Direction (negative for down, positive for up)

This is not a particularly nice interface, and leaves some questions open. For
example, the elevator state only has one goal floor; but it is conceivable
that an elevator holds more than one person, and each person wants to go to a
different floor, so there could be a few goal floors queued up.

=========================================

#### Set the GOPATH:
  * run the command from install.sh

#### Restrictions:
  * by default the following parameters are set:
    - minFloorNumber is set to 0
    - maxFloorNumber is set to 10

#### Build the Elevator Control System:
  * go build main.go

#### Start the Elevator Control System:
  * $ ./main -n (number of elevators)
  * e.g. start with 2 elevators: $ ./main -n 2

#### Commands:
 * status - returns the status of all the elevators in the form of a list of triples. Each triple represents one elevator in the following format: (ElevatorID, CurrentFloorNumber, GoalFloors[]).
 * step - allows one unit of time to pass, effectively telling the elevators to go to the next goal floor.
 * pickup floorNumber direction - adds a pickup request to the pickupRequests queue. The arguments (floor_number, direction [1 == up, -1 == down]) are space separated.
 * exit - exits the program.

#### FCFC Algorithm:
  * FCFS-First Come First Serve is an algorithm which is the most simple and traditional way to control elevators. It doesn’t optimize the floors-searching and timing as well. Elevators work depends on the order of every request. The first one who sends the request can take the elevator first.

#### Scheduling Algorithm:
  * My scheduling algorithm is based on the elevator algorithm (also SCAN). The elevator algorithm has found an application in computer operating systems as an algorithm for scheduling hard disk requests.
  * The elevator moves in the same direction as long as there are remaining requests stored in the map called goalFloorNumber.
  * If the goalFloorNumber is empty, the elevator will go into an idle state and change the direction if there are requests in the opposite direction.

#### Improvements to Scheduling:
  * If there are two elevator moving in the same direction and the pickup request is in the same direction, we would assign the elevator with the lower number of passengers. This would add some load ballancing to distribute the passengers accross moving elevators.
  * Yet another improvement would be an algorithm that remembered days and patterns might be able to better predict the comings and goings and different times (the ground floor rush at 9am and the floors leaving at 5pm). So an idle elevator would move to the ground floor at 9am or to the upper floors at 5pm.