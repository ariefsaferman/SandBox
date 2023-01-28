package parking

import (
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
)

type Lot struct {
	mapTicket      map[string]*entity.Car
	capacity       int
	subscriberList []Subscriber
}

func NewParkingLot(capacity int) Lot {
	pl := Lot{
		mapTicket: make(map[string]*entity.Car),
		capacity:  capacity,
	}
	return pl
}

func (pl *Lot) checkAvailableParkingLot(car *entity.Car) error {

	if len(pl.mapTicket) == pl.capacity {
		return ErrParkFull
	}

	for _, v := range pl.mapTicket {
		if v.PlateNumber == car.PlateNumber {
			return ErrParkTwice
		}
	}
	return nil
}

func (pl *Lot) Park(car *entity.Car) (*entity.Ticket, error) {

	// cek sebelum parkir
	if isFull := pl.checkAvailableParkingLot(car); isFull == ErrParkFull {
		return nil, ErrParkFull
	}

	for _, value := range pl.mapTicket {
		if value.PlateNumber == car.PlateNumber {
			return nil, ErrParkTwice
		}
	}

	ticket := entity.NewTicket()
	pl.mapTicket[ticket.ID] = car

	//cek setelah parkir
	if isFull := pl.checkAvailableParkingLot(car); isFull == ErrParkFull {
		for _, value := range pl.subscriberList {
			value.AddFullLot(pl)
		}
	}

	return &ticket, nil
}

func (pl *Lot) Unpark(ticket *entity.Ticket) (*entity.Car, error) {

	car, ok := pl.mapTicket[ticket.ID]
	if ok {
		delete(pl.mapTicket, ticket.ID)
		for _, value := range pl.subscriberList {
			value.AddAvailableLot(pl)
		}
		return car, nil
	}
	return nil, ErrParkTicket
}

func (pl *Lot) isHasBiggerCapacity(maxLot *Lot) bool {
	return pl.capacity > maxLot.capacity
}

func (pl *Lot) AddSubscriber(subscriber Subscriber) {
	pl.subscriberList = append(pl.subscriberList, subscriber)

}

func (pl *Lot) isHasBiggerSpace(lot *Lot) bool {
	return pl.capacity-len(pl.mapTicket) > lot.capacity-len(lot.mapTicket)
}

func (pl *Lot) GetSpaceLeft() int {
	return pl.capacity - len(pl.mapTicket)
}

func (pl *Lot) GetPlateCar() map[string]*entity.Car {
	return pl.mapTicket
}
