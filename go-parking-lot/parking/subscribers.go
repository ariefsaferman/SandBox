package parking

type Subscriber interface {
	AddFullLot(lot *Lot)
	AddAvailableLot(lot *Lot)
}
