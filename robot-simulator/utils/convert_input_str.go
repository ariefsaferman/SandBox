package utils

import "strings"

func ConvertInputCommand(str string) []string {
	return strings.Split(str, "")
}
