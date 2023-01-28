package main

import (
	"fmt"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/utils"
)

func promptInput(text string) string {
	var x string
	fmt.Print(text)
	fmt.Scanln(&x)
	return x
}

func displayListParkingLotStatus(listParkingLot []*parking.Lot) {
	fmt.Println("\nParking Lot Status:")

	for idx, lot := range listParkingLot {
		fmt.Printf("Lot #%d: %d spaces left\n", idx+1, lot.GetSpaceLeft())

		listCar := lot.GetPlateCar()
		for ticket, car := range listCar {
			fmt.Printf("#%s %s\n", ticket, car.PlateNumber)
		}
		if len(listCar) > 0 {
			fmt.Println("")
		}
	}
}

func main() {
	var attendant parking.Attendance
	var listParkingLot []*parking.Lot
	exit := false
	menu := "Parking Lot\n" +
		"1. Setup\n" +
		"2. Park\n" +
		"3. Un Park\n" +
		"4. Parking Lot Status\n" +
		"5. Exit"

	for !exit {
		fmt.Println(menu)
		input := promptInput("input menu: ")

		switch input {
		case "1":
			inputStr := promptInput("input parking lot capacities: ")
			capacities := utils.ConvertInputStringToIntSlices(inputStr)
			for _, capacity := range capacities {
				lot := parking.NewParkingLot(capacity)
				listParkingLot = append(listParkingLot, &lot)
			}
			attendant = parking.NewAttendance(listParkingLot)
		case "2":
			plateNumber := promptInput("input plate number: ")
			car := entity.Car{PlateNumber: plateNumber}
			ticket, err := attendant.Park(&car)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("Car parked with ticket id #%v\n", ticket.ID)
		case "3":
			inputStr := promptInput("input ticked id: ")
			ticket := utils.ConvertTicket(inputStr)

			car, err := attendant.Unpark(&ticket)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("Car %v successfully unparked!\n", car.PlateNumber)
		case "4":
			displayListParkingLotStatus(listParkingLot)
		case "5":
			exit = true
		default:
			fmt.Println("invalid menu")
		}
		fmt.Println("\n-------------------")
	}

}
