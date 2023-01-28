package grade

func GradeStudent(grade int) string {

	var result string

	if grade >= 90 && grade <= 100 {
		result = "A"
	} else if grade >= 80 && grade <= 89 {
		result = "B"
	} else if grade >= 70 && grade <= 79 {
		result = "C"
	} else if grade >= 60 && grade <= 69 {
		result = "D"
	} else {
		result = "E"
	}

	return result
}
