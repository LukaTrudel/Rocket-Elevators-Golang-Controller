package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewFloorRequestButton(_id int, _status string, _floor int, _direction string) *FloorRequestButton {
	n := new(FloorRequestButton)
	n.ID = _id
	n.status = _status
	n.floor = _floor
	n.direction = _direction
	return n
}
