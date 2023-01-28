package main

import (
	"fmt"
	"robot-simulator/robot"
	"robot-simulator/utils"
	"strings"
)

func main() {
	var x, y int
	var face string
	var sequenceCommand string
	fmt.Scanf("%d %d %s", &x, &y, &face)
	fmt.Scanln(&sequenceCommand)

	robot := robot.NewRobot(x, y, face)

	if ok := checkValidPosition(face); ok {
		fmt.Println("Invalid position!")
		return
	}

	for _, r := range utils.ConvertInputCommand(sequenceCommand) {
		err := robot.Perform(r)
		if err != nil {
			fmt.Println("Invalid movement!")
			return
		}
	}
	robot.DisplayPosition()
}

func checkValidPosition(face string) bool {
	return !strings.Contains(face, "N") && !strings.Contains(face, "S") && !strings.Contains(face, "E") && !strings.Contains(face, "W")
}
