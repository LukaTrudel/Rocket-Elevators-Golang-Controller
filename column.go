package main

import (
	"math"
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

func NewColumn(_id, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	c := Column{}
	c.ID = _id
	c.status = "online"
	c.amountOfFloors = amountOfFloors
	c.amountOfElevators = _amountOfElevators
	c.isBasement = _isBasement
	c.servedFloors = _servedFloors
	c.elevatorsList = []Elevator{}
	c.callButtonsList = []CallButton{}
	c.createElevators(amountOfFloors, _amountOfElevators)
	c.createCallButtons(amountOfFloors, _isBasement)

	return c
}
func (c *Column) createElevators(amountOfFloors int, amountOfElevators int) {
	elevatorID := 1
	for i := 0; i < amountOfElevators; i++ {
		elevator := newElevator(elevatorID, "idle", amountOfFloors, 1)
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

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator := c.findBestElevator(_requestedFloor, _direction)
	elevator.floorRequestList = append(elevator.floorRequestList, 1)
	elevator.sortFloorList()
	elevator.move()
	elevator.openDoors()
}

func (c *Column) findBestElevator(requestedFloor int, requestedDirection string) Elevator { // This function in conjuction wwith checkElevator will return the best elevator
	bestElevatorInfo := map[string]interface{}{
		"bestElevator": nil,
		"bestScore":    6,
		"referenceGap": math.Inf(1),
	}
	if requestedFloor == 1 {
		for _, e := range c.elevatorsList {
			if 1 == e.currentFloor && e.status == "stopped" {
				bestElevatorInfo = c.checkElevator(1, e, requestedFloor, bestElevatorInfo)
			} else if 1 == e.currentFloor && e.status == "idle" {
				bestElevatorInfo = c.checkElevator(2, e, requestedFloor, bestElevatorInfo)
			} else if 1 > e.currentFloor && e.direction == "up" {
				bestElevatorInfo = c.checkElevator(3, e, requestedFloor, bestElevatorInfo)
			} else if 1 < e.currentFloor && e.direction == "down" {
				bestElevatorInfo = c.checkElevator(3, e, requestedFloor, bestElevatorInfo)
			} else if e.status == "idle" {
				bestElevatorInfo = c.checkElevator(4, e, requestedFloor, bestElevatorInfo)
			} else {
				bestElevatorInfo = c.checkElevator(5, e, requestedFloor, bestElevatorInfo)
			}
		}
	} else {
		for _, e := range c.elevatorsList {
			if requestedFloor == e.currentFloor && e.status == "idle" && requestedDirection == e.direction {
				bestElevatorInfo = c.checkElevator(1, e, requestedFloor, bestElevatorInfo)
			} else if requestedFloor > e.currentFloor && e.direction == "up" && requestedDirection == e.direction {
				bestElevatorInfo = c.checkElevator(2, e, requestedFloor, bestElevatorInfo)
			} else if requestedFloor < e.currentFloor && e.direction == "down" && requestedDirection == e.direction {
				bestElevatorInfo = c.checkElevator(2, e, requestedFloor, bestElevatorInfo)
			} else if e.status == "stopped" {
				bestElevatorInfo = c.checkElevator(4, e, requestedFloor, bestElevatorInfo)
			} else {
				bestElevatorInfo = c.checkElevator(5, e, requestedFloor, bestElevatorInfo)
			}
		}
	}
	return bestElevatorInfo["bestElevator"].(Elevator)
}

func (c *Column) checkElevator(baseScore int, elevator Elevator, floor int, bestElevatorInfo map[string]interface{}) map[string]interface{} {
	if baseScore < bestElevatorInfo["bestScore"].(int) {
		bestElevatorInfo["bestScore"] = baseScore
		bestElevatorInfo["bestElevator"] = elevator
		bestElevatorInfo["referenceGap"] = Abs(elevator.currentFloor - floor)
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
