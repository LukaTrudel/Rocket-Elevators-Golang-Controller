package main

import (
	//"fmt"
	"math"
	"strconv"
)

var elevatorID int = 1
var callButtonID int = 1

type Column struct {
	ID               string
	status           string
	servedFloorsList []int
	isBasement       bool
	elevatorsList    []*Elevator
	callButtonsList  []CallButton
}

func NewColumn(_id string, _status string, _amountOfFloors, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	newColumn := new(Column)
	newColumn.ID = _id
	newColumn.status = _status
	newColumn.servedFloorsList = _servedFloors

	newColumn.createElevators(_amountOfFloors, _amountOfElevators)
	newColumn.createCallButtons(_amountOfFloors, _isBasement)
	return newColumn
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

func (c *Column) findElevator(_requestedFloor int, _requestedDirection string) *Elevator {
	var bestElevator *Elevator
	var bestScore int = 6
	var referenceGap int = 100000
	// type BestElevatorInformations struct {
	// 	bestElevator            *Elevator
	// 	bestScore, referenceGap int
	// }

	if _requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			if 1 == elevator.currentFloor && elevator.status == "stopped" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if 1 == elevator.currentFloor && elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			}
		}
	} else {
		for _, elevator := range c.elevatorsList {
			if _requestedFloor == elevator.currentFloor && elevator.status == "stopped" && _requestedDirection == elevator.direction {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if _requestedFloor > elevator.currentFloor && elevator.direction == "up" && _requestedDirection == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if _requestedFloor < elevator.currentFloor && elevator.direction == "down" && _requestedDirection == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
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
