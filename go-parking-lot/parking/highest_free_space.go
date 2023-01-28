package parking

type HighestFreeSpace struct {
}

func (highestFreeSpace HighestFreeSpace) SelectLot(lots []*Lot) (selectedLot *Lot) {
	var newLot Lot
	selectedLot = &newLot
	for _, lot := range lots {
		if lot.isHasBiggerSpace(selectedLot) {
			selectedLot = lot
		}
	}
	return selectedLot
}
