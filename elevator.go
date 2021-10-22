package main

import (
	//"fmt"
	"sort"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Elevator struct {
	ID                    string
	status                string
	amountOfFloors        int
	direction             string
	currentFloor          int
	door                  Door
	floorRequestsList     []int
	completedRequestsList []int
}

func NewElevator(_id, _status string, _amountOfFloors, _currentFloor int) *Elevator {
	e := new(Elevator)
	e.ID = _id
	e.status = _status
	e.amountOfFloors = _amountOfFloors
	e.currentFloor = _currentFloor
	e.direction = ""
	e.door = Door{1, ""}

	return e

}

func (e *Elevator) move() {
	for len(e.floorRequestsList) != 0 {
		var destination int = e.floorRequestsList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			e.sortFloorList()
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			e.sortFloorList()
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.operateDoors()
		e.completedRequestsList = append(e.completedRequestsList, e.floorRequestsList[0])
		e.floorRequestsList = e.floorRequestsList[1:]
	}
	e.status = "idle"
}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Ints(e.floorRequestsList)
	} else if e.direction == "down" {
		sort.Ints(e.floorRequestsList)
		e.floorRequestsList = (e.floorRequestsList)
	}
}
func (elevator *Elevator) operateDoors() {
	if elevator.status == "stopped" || elevator.status == "idle" {
		elevator.door.status = "open"
		// fmt.Println("Doors: ", elevator.door.status)
		// fmt.Println("(doors stay open for 6 seconds)")
		if len(elevator.floorRequestsList) < 1 {
			elevator.direction = ""
			elevator.status = "idle"
			//fmt.Println("Elevator ", elevator.ID, " status: ", elevator.status)
		}
	}
}

func (e *Elevator) addNewRequest(requestedFloor int) {
	if !containsElement(requestedFloor, e.floorRequestsList) {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)
	}

	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}

	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}
}
