package main

type Elevator struct {
	ID               int
	status           string
	amountOfFloors   int
	direction        string
	currentFloor     int
	door             Door
	floorRequestList []int
}

func NewElevator(_elevatorID string) *Elevator {
	e := Elevator{}
	e.ID = _elevatorID
	e.status = "online"
	e.amountOfFloors = amountOfFloors
	e.direction = "null"
	e.currentFloor = currentFloor
	e.door = Door{}
	e.floorRequestList = []int{}

	return e

}

func (e *Elevator) move() {

}
