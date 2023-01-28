package parking

type DefaultStyle struct{}

func (defaultStyle DefaultStyle) SelectLot(lots []*Lot) *Lot {
	return lots[0]
}
