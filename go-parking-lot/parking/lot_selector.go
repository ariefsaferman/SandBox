package parking

type LotSelector interface {
	SelectLot([]*Lot) *Lot
}
