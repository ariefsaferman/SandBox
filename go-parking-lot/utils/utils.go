package utils

import (
	"strconv"
	"strings"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
)

func ConvertInputStringToIntSlices(input string) []int {
	arrOfString := strings.Split(input, ",")
	var result []int
	for _, str := range arrOfString {
		intInput, err := strconv.Atoi(str)
		if err == nil {
			result = append(result, intInput)
		}
	}
	return result
}

func ConvertTicket(ticket string) entity.Ticket {
	return entity.Ticket{ID: ticket[1:]}
}
