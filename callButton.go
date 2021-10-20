package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewCallButton(_id int, _floor int, _direction string) *CallButton {
	n := new(CallButton)
	n.ID = _id
	n.status = ""
	n.floor = _floor
	n.direction = _direction
	return n

}
