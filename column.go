package main

import (
	"math"
	"strconv"
)

var elevatorID int = 1
var callButtonID int = 1
var currentTime int = 1

type Column struct {
	ID               string
	status           string
	servedFloorsList []int
	isBasement       bool
	elevatorsList    []*Elevator
	callButtonsList  []CallButton
}

func NewColumn(_id string, _status string, _amountOfFloors, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	c := new(Column)
	c.ID = _id
	c.status = _status
	c.servedFloorsList = _servedFloors

	c.createElevators(_amountOfFloors, _amountOfElevators)
	c.createCallButtons(_amountOfFloors, _isBasement)
	return c
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	if _isBasement {
		var buttonFloor int = -1

		for i := 0; i < _amountOfFloors; i++ {
			c.callButtonsList = append(c.callButtonsList, *NewCallButton(callButtonID, buttonFloor, "up"))
			buttonFloor--
			callButtonID++
		}
	} else {
		var buttonFloor int = 1
		for i := 0; i < _amountOfFloors; i++ {
			c.callButtonsList = append(c.callButtonsList, *NewCallButton(callButtonID, buttonFloor, "up"))
			buttonFloor++
			callButtonID++
		}
	}
}

func (c *Column) createElevators(_amountOfFloors, _amountOfElevators int) {
	for i := 0; i < _amountOfElevators; i++ {
		c.elevatorsList = append(c.elevatorsList, NewElevator(strconv.Itoa(elevatorID), "idle", _amountOfFloors, 1))
		elevatorID++
	}
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator := c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	elevator.addNewRequest(1)
	elevator.move()
	return elevator
}

//We use a score system depending on the current elevators state. Since the bestScore and the referenceGap are
//higher values than what could be possibly calculated, the first elevator will always become the default bestElevator,
//before being compared with to other elevators. If two elevators get the same score, the nearest one is prioritized. Unlike
//the classic algorithm, the logic isn't exactly the same depending on if the request is done in the lobby or on a floor.

func (c *Column) findElevator(_requestedFloor int, _requestedDirection string) *Elevator {
	var bestElevator *Elevator
	var bestScore int = 6
	var referenceGap int = 100000

	if _requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			//The elevator is at the lobby and already has some requests. It is about to leave but has not yet departed
			if 1 == elevator.currentFloor && elevator.status == "stopped" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
				//The elevator is at the lobby and has no requests
			} else if 1 == elevator.currentFloor && elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
				//The elevator is lower than me and is coming up.
				//It means that I'm requesting an elevator to go to a basement, and the elevator is on it's way to me.
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
				//The elevator is above me and is coming down.
				//It means that I'm requesting an elevator to go to a floor, and the elevator is on it's way to me
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
				//The elevator is not at the first floor, but doesn't have any request
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
				//The elevator is not available, but still could take the call if nothing better is found
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			}
		}
	} else {
		for _, elevator := range c.elevatorsList {
			//The elevator is at the same level as me, and is about to depart to the first floor
			if _requestedFloor == elevator.currentFloor && elevator.status == "stopped" && _requestedDirection == elevator.direction {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
				//The elevator is lower than me and is going up. I'm on a basement, and the elevator can pick me up on it's way
			} else if _requestedFloor > elevator.currentFloor && elevator.direction == "up" && _requestedDirection == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
				//The elevator is higher than me and is going down. I'm on a floor, and the elevator can pick me up on it's way
			} else if _requestedFloor < elevator.currentFloor && elevator.direction == "down" && _requestedDirection == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
				//The elevator is idle and has no requests
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
				//The elevator is not available, but still could take the call if nothing better is found
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			}
		}
	}
	return bestElevator
}

func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator *Elevator, bestScore int, referenceGap int, bestElevator *Elevator, floor int) (*Elevator, int, int) {
	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator
		referenceGap = int(math.Abs(float64(newElevator.currentFloor) - float64(floor)))
	} else if bestScore == scoreToCheck {
		var gap int = int(math.Abs(float64(newElevator.currentFloor) - float64(floor)))
		if referenceGap > gap {
			bestElevator = newElevator
			referenceGap = gap
		}
	}
	return bestElevator, bestScore, referenceGap
}
