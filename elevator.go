package main

import (
	"fmt"
	"sort"
)

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

func NewElevator(_elevatorID string) *Elevator {
	e := Elevator{}
	e.ID = _elevatorID
	e.status = "online"
	// e.amountOfFloors = amountOfFloors
	e.direction = "null"
	// e.currentFloor = currentFloor
	e.door = Door{}
	e.floorRequestsList = []int{}

	return e

}

func (e *Elevator) move() {
	for len(e.floorRequestsList) != 0 {
		destination := e.floorRequestsList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "idle"
		e.operateDoors()
		e.completedRequestsList = append(e.floorRequestsList, 0)
		e.floorRequestsList = RemoveIndex(e.floorRequestsList, 0)
	}
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...) // Function created to remove the first index of a list
}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Slice(e.floorRequestsList, func(i, j int) bool { return e.floorRequestsList[i] < e.floorRequestsList[j] })
	} else {
		sort.Slice(e.floorRequestsList, func(i, j int) bool { return e.floorRequestsList[i] > e.floorRequestsList[j] })
	}
}

func (elevator *Elevator) operateDoors() {
	if elevator.status == "stopped" || elevator.status == "idle" {
		elevator.door.status = "open"
		fmt.Println("Doors: ", elevator.door.status)
		fmt.Println("(doors stay open for 6 seconds)")
		if len(elevator.floorRequestsList) < 1 {
			elevator.direction = ""
			elevator.status = "idle"
			fmt.Println("Elevator ", elevator.ID, " status: ", elevator.status)
		}
	}
}
