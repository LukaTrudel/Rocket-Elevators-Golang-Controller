package main

import (
	"fmt"
	"math"
	"strconv"
)

type Column struct {
	ID                int
	status            string
	amountOfFloors    int
	amountOfElevators int
	isBasement        bool
	elevatorsList     []Elevator
	callButtonsList   []CallButton
	servedFloors      []int
}

func NewColumn(_id, _amountOfElevators, _amountOfFloors int, _servedFloors []int, _isBasement bool) *Column {
	c := Column{}
	c.ID = _id
	c.status = "online"
	c.amountOfFloors = _amountOfFloors
	c.amountOfElevators = _amountOfElevators
	c.isBasement = _isBasement
	c.servedFloors = _servedFloors
	c.elevatorsList = []Elevator{}
	c.callButtonsList = []CallButton{}
	c.createElevators(_amountOfFloors, _amountOfElevators)
	c.createCallButtons(_amountOfFloors, _isBasement)

	return c
}
func (c *Column) createElevators(amountOfFloors int, amountOfElevators int) {
	elevatorID := 1
	for i := 0; i < amountOfElevators; i++ {
		elevator := newElevator(strconv.Itoa(elevatorID))
		c.elevatorsList = append(c.elevatorsList, elevator)
		elevatorID++
	}
}
func (c *Column) createCallButtons(amountOfFloors int, isBasement bool) {
	callButtonID := 1
	if isBasement == true {
		buttonFloor := -1
		for i := 0; i < amountOfFloors; i++ {
			c.callButtonsList = append(c.callButtonsList, CallButton{callButtonID, "off", buttonFloor, "up"})
			buttonFloor--
			callButtonID++
		}
	} else {
		buttonFloor := 1
		for i := 0; i < amountOfFloors; i++ {
			c.callButtonsList = append(c.callButtonsList, CallButton{callButtonID, "off", buttonFloor, "down"})
			buttonFloor++
			callButtonID++
		}
	}
}
// func (column *Column) requestElevator(userFloor int, direction string) {
// 	fmt.Println("||Passenger requests elevator from", userFloor, "going", direction, "to the lobby||")
// 	var elevator Elevator = column.findBestElevator(userFloor, direction)
// 	fmt.Println("||", elevator.ID, "is the assigned elevator for this request||")
// 	elevator.floorRequestsList = append(elevator.floorRequestsList, userFloor)
// 	//elevator.sortFloorList()
// 	elevator.move()
// 	elevator.operateDoors()
}

//Simulate when a user press a button on a floor to go back to the first floor
func (column *Column) findBestElevator(floor int, direction string) Elevator {
	requestedFloor := floor
	requestedDirection := direction
	bestElevatorInfo := BestElevatorInfo{
		bestElevator: Elevator{},
		bestScore:    6,
		referenceGap: 1000000,
	}

	if requestedFloor == 1 {
		for _, elevator := range column.elevatorsList {
			if 1 == elevator.currentFloor && elevator.status == "stopped" {
				bestElevatorInfo = column.checkIfElevatorIsBetter(1, elevator, bestElevatorInfo, requestedFloor)
			} else if 1 == elevator.currentFloor && elevator.status == "idle" {
				bestElevatorInfo = column.checkIfElevatorIsBetter(2, elevator, bestElevatorInfo, requestedFloor)
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevatorInfo = column.checkIfElevatorIsBetter(3, elevator, bestElevatorInfo, requestedFloor)
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {
				bestElevatorInfo = column.checkIfElevatorIsBetter(3, elevator, bestElevatorInfo, requestedFloor)
			} else if elevator.status == "idle" {
				bestElevatorInfo = column.checkIfElevatorIsBetter(4, elevator, bestElevatorInfo, requestedFloor)
			} else {
				bestElevatorInfo = column.checkIfElevatorIsBetter(5, elevator, bestElevatorInfo, requestedFloor)
			}
		}
	} else {
		for _, elevator := range column.elevatorsList {
			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
				bestElevatorInfo = column.checkIfElevatorIsBetter(1, elevator, bestElevatorInfo, requestedFloor)
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {
				bestElevatorInfo = column.checkIfElevatorIsBetter(2, elevator, bestElevatorInfo, requestedFloor)
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {
				bestElevatorInfo = column.checkIfElevatorIsBetter(2, elevator, bestElevatorInfo, requestedFloor)
			} else if elevator.status == "idle" {
				bestElevatorInfo = column.checkIfElevatorIsBetter(3, elevator, bestElevatorInfo, requestedFloor)
			} else {
				bestElevatorInfo = column.checkIfElevatorIsBetter(4, elevator, bestElevatorInfo, requestedFloor)
			}
		}
	}
	return bestElevatorInfo.bestElevator
}

func (column *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator Elevator, bestElevatorInfo BestElevatorInfo, floor int) BestElevatorInfo {

	if scoreToCheck < bestElevatorInfo.bestScore {
		bestElevatorInfo.bestScore = scoreToCheck
		bestElevatorInfo.bestElevator = newElevator
		bestElevatorInfo.referenceGap = int(math.Abs(float64(newElevator.currentFloor - floor)))
	} else if bestElevatorInfo.bestScore == scoreToCheck {
		gap := int(math.Abs(float64(newElevator.currentFloor - floor)))
		if bestElevatorInfo.referenceGap > gap {
			bestElevatorInfo.bestScore = scoreToCheck
			bestElevatorInfo.bestElevator = newElevator
			bestElevatorInfo.referenceGap = gap
		}
	}
	return bestElevatorInfo
}

// Abs ...
func Abs(x int) int { // Function created to return the absolute value of an int
	if x < 0 {
		return -x
	}
	return x
}

type BestElevatorInfo struct {
	bestElevator Elevator
	bestScore    int
	referenceGap int
}
