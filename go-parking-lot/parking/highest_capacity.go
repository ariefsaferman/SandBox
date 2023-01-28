package parking

type HighestCapacity struct{}

func (highestCapacity HighestCapacity) SelectLot(lots []*Lot) (selectedLot *Lot) {
	var newLot Lot
	selectedLot = &newLot
	for _, lot := range lots {
		if lot.isHasBiggerCapacity(selectedLot) {
			selectedLot = lot
		}
	}
	return selectedLot
}
