package parking

import (
	"testing"

	. "git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"github.com/stretchr/testify/assert"
)

func TestPark(t *testing.T) {
	t.Run("Should return ticket when park a car", func(t *testing.T) {
		parkingLot := NewParkingLot(2)
		car := Car{PlateNumber: "B456"}
		result, _ := parkingLot.Park(&car)

		assert.NotNil(t, result)
	})

	t.Run("should return error when parking is full", func(t *testing.T) {
		parkingLot := NewParkingLot(2)
		car := Car{PlateNumber: "B456"}
		car2 := Car{PlateNumber: "C"}
		car3 := Car{PlateNumber: "D"}

		_, err := parkingLot.Park(&car)
		_, err2 := parkingLot.Park(&car2)
		_, err3 := parkingLot.Park(&car3)

		assert.Nil(t, err)
		assert.Nil(t, err2)
		assert.ErrorIs(t, err3, ErrParkFull)
	})

	t.Run("should return error when park twice", func(t *testing.T) {
		parkingLot := NewParkingLot(2)
		car := Car{PlateNumber: "B"}

		_, err := parkingLot.Park(&car)
		_, err2 := parkingLot.Park(&car)

		assert.Nil(t, err)
		assert.ErrorIs(t, err2, ErrParkTwice)
	})
}

func TestUnpark(t *testing.T) {
	t.Run("Should return car when unpark a car", func(t *testing.T) {
		parkingLot := NewParkingLot(2)
		car := Car{PlateNumber: "B456"}
		ticket, _ := parkingLot.Park(&car)
		unparkedCar, _ := parkingLot.Unpark(ticket)

		assert.Equal(t, &car, unparkedCar)
	})

	t.Run("should return error after unpark with invalid Ticket", func(t *testing.T) {
		parkingLot := NewParkingLot(2)
		car := Car{PlateNumber: "B456"}
		parkingLot.Park(&car)
		ticket2 := NewTicket()
		_, err := parkingLot.Unpark(&ticket2)

		assert.ErrorIs(t, err, ErrParkTicket)
	})
}
