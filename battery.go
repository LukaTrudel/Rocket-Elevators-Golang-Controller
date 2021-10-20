package main

import (
	"math"
)

var floorRequestButtonID int = 1
var columnID int = 1

type Battery struct {
	ID                      int
	status                  string
	amountOfFloors          int
	amountOfColumns         int
	amountOfBasements       int
	columnsList             []Column
	floorRequestButtonsList []FloorRequestButton
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	b := Battery{}
	b.ID = _id
	b.status = "online"
	b.amountOfFloors = _amountOfFloors
	b.amountOfColumns = _amountOfColumns
	b.amountOfBasements = _amountOfBasements
	b.columnsList = []Column{}
	b.floorRequestButtonsList = []FloorRequestButton{}

	if _amountOfBasements > 0 {
		b.createBasmentColumn(b.amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	b.createFloorRequestButtons(_amountOfFloors)
	b.createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn)

	return b
}

func (b *Battery) createBasmentColumn(amountOfBasements int, amountOfElevatorPerColumn int) { // This will create the column for the basements
	servedFloors := []int{}
	floor := -1

	for i := 0; i < amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
	column := newColumn(columnID, amountOfElevator, amountOfBasements, servedFloors, true)
	b.columnsList = append(b.columnsList, column)
	columnID++
}
func (b *Battery) createColumns(amountOfColumns int, amountOfFloors int, amountOfElevatorPerColumn int) { // this will create the columns with thier floors
	amountOfFloorsPerColumn := math.Ceil(float64(amountOfFloors / amountOfColumns))
	n := int(amountOfFloorsPerColumn)
	floor := 1

	for i := 0; i < amountOfColumns; i++ {
		servedFloors := []int{}
		for j := 0; j < n; j++ {
			if floor <= amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column := newColumn(columnID, amountOfElevators, amountOfFloors, servedFloors, false)
		b.columnsList = append(b.columnsList, column)
		columnID++
	}
}

func (b *Battery) createFloorRequestButtons(amountOfFloors int) {
	buttonFloor := 1
	for i := 0; i < amountOfFloors; i++ {
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, FloorRequestButton{floorRequestButtonID, "off", buttonFloor})
		floorRequestButtonID++
		buttonFloor++
	}
}

func (b *Battery) createBasementFloorRequestButtons(amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < amountOfBasements; i++ {
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, FloorRequestButton{floorRequestButtonID, "off", buttonFloor})
		buttonFloor--
		floorRequestButtonID++
	}
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	col := Column{}
	for _, c := range b.columnsList {
		found := find(_requestedFloor, c.servedFloors)
		if found == true {
			col = c
		}
	}
	return col
}

func find(a int, list []int) bool { // function created to check if a list contains a given number
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := b.findBestColumn(_requestedFloor)
	elevator := column.findBestElevator(1, _direction)
	elevator.floorRequestsList = append(elevator.floorRequestsList, _requestedFloor)
	elevator.sortFloorList()
	elevator.move()
	elevator.operateDoors()

}
