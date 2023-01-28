package parking

import "errors"

var ErrParkTicket = errors.New("unrecognized parking ticket")
var ErrParkFull = errors.New("no available position")
var ErrParkTwice = errors.New("cannot park twice")
