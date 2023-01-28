package robot

import "errors"

var (
	ErrInvalidMovement = errors.New("invalid movement")
	ErrInvalidPosition = errors.New("invalid position")
)
