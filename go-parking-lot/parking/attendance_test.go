package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/mocks"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	"github.com/stretchr/testify/assert"
)

func TestAttendancePark(t *testing.T) {
	t.Run("Should return ticket when park a car", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(2)
		listParkingLot := []*parking.Lot{&parkingLot}
		car := entity.Car{PlateNumber: "B456"}
		attendance := parking.NewAttendance(listParkingLot)

		result, _ := attendance.Park(&car)

		assert.NotNil(t, result)
	})

	t.Run("should return error when parking is full", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(2)
		listParkingLot := []*parking.Lot{&parkingLot}
		car := entity.Car{PlateNumber: "B456"}
		car2 := entity.Car{PlateNumber: "C"}
		car3 := entity.Car{PlateNumber: "D"}
		attendance := parking.NewAttendance(listParkingLot)

		_, err := attendance.Park(&car)
		_, err2 := attendance.Park(&car2)
		_, err3 := attendance.Park(&car3)

		assert.Nil(t, err)
		assert.Nil(t, err2)
		assert.ErrorIs(t, err3, parking.ErrParkFull)
	})

	t.Run("should return error when park twice", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(2)
		listParkingLot := []*parking.Lot{&parkingLot}
		car := entity.Car{PlateNumber: "B"}
		attendance := parking.NewAttendance(listParkingLot)

		_, err := attendance.Park(&car)
		_, err2 := attendance.Park(&car)

		assert.Nil(t, err)
		assert.ErrorIs(t, err2, parking.ErrParkTwice)
	})

	t.Run("should invoke notify to subscriber when lot is full", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(1)
		car := entity.Car{PlateNumber: "B"}
		mockSubs := &mocks.Subscriber{}
		mockSubs.On("AddFullLot", &parkingLot).Return()

		parkingLot.AddSubscriber(mockSubs)
		ticket, _ := parkingLot.Park(&car)

		assert.NotNil(t, ticket)
		mockSubs.AssertNumberOfCalls(t, "AddFullLot", 1)
	})

	t.Run("should park on highest capacity when parking style HighestCapacity", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(1)
		parkingLot2 := parking.NewParkingLot(2)
		parkingLot3 := parking.NewParkingLot(4)
		parkingLot4 := parking.NewParkingLot(3)
		listParkingLot := []*parking.Lot{&parkingLot, &parkingLot2, &parkingLot3, &parkingLot4}
		attendant := parking.NewAttendance(listParkingLot)
		car := entity.Car{PlateNumber: "ABC"}
		attendant.ChangeParkingStyle(parking.HighestCapacity{})

		ticket, _ := attendant.Park(&car)
		unparkedCar, _ := parkingLot3.Unpark(ticket)

		assert.Equal(t, &car, unparkedCar)
	})

	t.Run("should park on highest free space when parking style HighestFreeSpace", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(1)
		parkingLot2 := parking.NewParkingLot(2)
		parkingLot3 := parking.NewParkingLot(4)
		parkingLot4 := parking.NewParkingLot(3)
		listParkingLot := []*parking.Lot{&parkingLot, &parkingLot2, &parkingLot3, &parkingLot4}
		attendant := parking.NewAttendance(listParkingLot)
		car := entity.Car{PlateNumber: "ABC"}
		car2 := entity.Car{PlateNumber: "BAC"}
		car3 := entity.Car{PlateNumber: "BCA"}
		attendant.ChangeParkingStyle(parking.HighestFreeSpace{})

		ticket, _ := attendant.Park(&car)
		ticket2, _ := attendant.Park(&car2)
		ticket3, _ := attendant.Park(&car3)
		unparkedCar, _ := parkingLot3.Unpark(ticket)
		unparkedCar2, _ := parkingLot3.Unpark(ticket2)
		unparkedCar3, _ := parkingLot4.Unpark(ticket3)

		assert.Equal(t, &car, unparkedCar)
		assert.Equal(t, &car2, unparkedCar2)
		assert.Equal(t, &car3, unparkedCar3)
	})

}

func TestAttendanceUnpark(t *testing.T) {
	t.Run("Should return car when unpark a car", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(1)
		listParkingLot := []*parking.Lot{&parkingLot}
		car := entity.Car{PlateNumber: "B456"}
		attendance := parking.NewAttendance(listParkingLot)

		ticket, _ := attendance.Park(&car)
		unparkedCar, _ := attendance.Unpark(ticket)

		assert.Equal(t, &car, unparkedCar)
	})

	t.Run("should return error after unpark with invalid Ticket", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(2)
		car := entity.Car{PlateNumber: "B456"}
		listParkingLot := []*parking.Lot{&parkingLot}
		attendance := parking.NewAttendance(listParkingLot)

		attendance.Park(&car)
		ticket2 := entity.NewTicket()
		_, err := attendance.Unpark(&ticket2)

		assert.ErrorIs(t, err, parking.ErrParkTicket)
	})

	t.Run("should invoke notify to subscriber when lot is available", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(1)
		car := entity.Car{PlateNumber: "B"}
		mockSubs := &mocks.Subscriber{}
		mockSubs.On("AddFullLot", &parkingLot).Return()
		mockSubs.On("AddAvailableLot", &parkingLot).Return()

		parkingLot.AddSubscriber(mockSubs)
		ticket, _ := parkingLot.Park(&car)
		parkingLot.Unpark(ticket)

		assert.NotNil(t, ticket)
		mockSubs.AssertNumberOfCalls(t, "AddAvailableLot", 1)
	})

	t.Run("should return error when used ticket is submitted", func(t *testing.T) {
		lot1 := parking.NewParkingLot(1)
		lot2 := parking.NewParkingLot(1)
		car := entity.Car{PlateNumber: "B"}
		attend := parking.NewAttendance([]*parking.Lot{&lot1, &lot2})
		ticket := entity.NewTicket()

		attend.Park(&car)
		attend.Unpark(&ticket)
		_, err := attend.Unpark(&ticket)

		assert.ErrorIs(t, err, parking.ErrParkTicket)
	})

}

func TestFreeSpace(t *testing.T) {
	t.Run("should return a value when lot is available", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(1)

		freeSpace := parkingLot.GetSpaceLeft()

		assert.NotEmpty(t, freeSpace)
	})

	t.Run("should return nill value when lot is full", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(1)
		car := entity.Car{PlateNumber: "BCA"}
		attendant := parking.NewAttendance([]*parking.Lot{&parkingLot})

		attendant.Park(&car)
		freeSpace := parkingLot.GetSpaceLeft()

		assert.Empty(t, freeSpace)
	})
}

func TestPlateCar(t *testing.T) {
	t.Run("should return a value when there is available car in lots", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(1)
		car := entity.Car{PlateNumber: "BCA"}
		attendance := parking.NewAttendance([]*parking.Lot{&parkingLot})

		attendance.Park(&car)
		plateCar := parkingLot.GetPlateCar()

		assert.NotEmpty(t, plateCar)

	})

	t.Run("should return a nill value when there is no car in lots", func(t *testing.T) {
		parkingLot := parking.NewParkingLot(1)

		plateCar := parkingLot.GetPlateCar()

		assert.Empty(t, plateCar)

	})

}
