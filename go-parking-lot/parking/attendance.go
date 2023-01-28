package parking

import (
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
)

type Attendance struct {
	listParkingLot      []*Lot
	fullParkingLot      []*Lot
	availableParkingLot []*Lot
	parkingStyle        LotSelector
}

func NewAttendance(lot []*Lot) Attendance {
	attendance := Attendance{
		listParkingLot:      lot,
		fullParkingLot:      []*Lot{},
		availableParkingLot: lot,
		parkingStyle:        DefaultStyle{},
	}

	for _, value := range lot {
		value.AddSubscriber(&attendance)
	}
	return attendance
}

func (attendant *Attendance) AddFullLot(lot *Lot) {
	attendant.fullParkingLot = append(attendant.fullParkingLot, lot)
	for idx, value := range attendant.availableParkingLot {
		if value == lot {
			attendant.availableParkingLot = append(attendant.availableParkingLot[:idx], attendant.availableParkingLot[idx+1:]...)
			break
		}
	}
}

func (attendant *Attendance) AddAvailableLot(lot *Lot) {

	attendant.availableParkingLot = append(attendant.availableParkingLot, lot)
	for idx, value := range attendant.fullParkingLot {
		if value == lot {
			attendant.fullParkingLot = append(attendant.fullParkingLot[:idx], attendant.fullParkingLot[idx+1:]...)
			break
		}
	}
}

func (attendant *Attendance) Park(car *entity.Car) (*entity.Ticket, error) {
	parkingLot := attendant.parkingStyle.SelectLot(attendant.availableParkingLot)

	for _, lot := range attendant.listParkingLot {
		err := lot.checkAvailableParkingLot(car)
		if err != nil {
			if err == ErrParkFull {
				return nil, ErrParkFull
			} else if err == ErrParkTwice {
				return nil, ErrParkTwice
			}
		}
	}

	ticket, _ := parkingLot.Park(car)
	return ticket, nil
}

func (attendant *Attendance) Unpark(ticket *entity.Ticket) (*entity.Car, error) {
	for _, lot := range attendant.listParkingLot {
		car, ok := lot.mapTicket[ticket.ID]
		if ok {
			lot.Unpark(ticket)
			return car, nil
		}
	}
	return &entity.Car{}, ErrParkTicket
}

func (attendant *Attendance) ChangeParkingStyle(style LotSelector) {
	attendant.parkingStyle = style
}
