package robot

import (
	"fmt"
)

const (
	NORTH = "N"
	SOUTH = "S"
	EAST  = "E"
	WEST  = "W"

	LEFT     = "L"
	ADVANCED = "A"
	RIGHT    = "R"
)

type Robot struct {
	x, y      int
	direction string
}

func NewRobot(x, y int, s string) *Robot {
	r := Robot{
		x:         x,
		y:         y,
		direction: s,
	}
	return &r
}

func (r *Robot) Perform(input string) error {
	switch input {
	case LEFT:
		r.faceLeft()
	case ADVANCED:
		r.advanced()
	case RIGHT:
		r.faceRight()
	default:
		return ErrInvalidMovement
	}
	return nil
}

func (r *Robot) faceLeft() {
	switch r.direction {
	case NORTH:
		r.direction = WEST
	case SOUTH:
		r.direction = EAST
	case EAST:
		r.direction = NORTH
	case WEST:
		r.direction = SOUTH
	}
}

func (r *Robot) faceRight() {
	switch r.direction {
	case NORTH:
		r.direction = EAST
	case SOUTH:
		r.direction = WEST
	case EAST:
		r.direction = SOUTH
	case WEST:
		r.direction = NORTH
	}
}

func (r *Robot) advanced() {
	switch r.direction {
	case NORTH:
		r.y++
	case SOUTH:
		r.y--
	case EAST:
		r.x++
	case WEST:
		r.x--
	}
}

func (r *Robot) DisplayPosition() {
	fmt.Println(r.x, r.y, r.direction)
}
