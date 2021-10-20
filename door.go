package main

type Door struct {
	ID     int
	status string
}

func NewDoor(_id int) *Door {
	doors := new(Door)
	doors.ID = _id
	doors.status = ""
	return doors
}
