package grade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradeStudent(t *testing.T) {
	type errorTestCases struct {
		description string
		input       int
		expected    string
	}

	for _, scenario := range []errorTestCases{
		{
			description: "Should return A when grade is between 90 to 100",
			input:       95,
			expected:    "A",
		},
		{
			description: "Should return B when grade is between 80 to 89",
			input:       85,
			expected:    "B",
		},
		{
			description: "Should return C when grade is between 70 to 79",
			input:       75,
			expected:    "C",
		},
		{
			description: "Should return D when grade is between 60 to 69",
			input:       65,
			expected:    "D",
		},
		{
			description: "Should return E when grade is between 0 to 59",
			input:       35,
			expected:    "E",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			grade := scenario.expected
			result := GradeStudent(scenario.input)
			assert.Equal(t, grade, result)
		})
	}
}
