package validParentheses

import (
	"testing"
)

func TestDataSetOne(t *testing.T) {
	// Arrange
	input := "()"
	expectedResult := true

	// Act
	result := isValid(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetTwo(t *testing.T) {
	// Arrange
	input := "()[]{}"
	expectedResult := true

	// Act
	result := isValid(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetThree(t *testing.T) {
	// Arrange
	input := "(]"
	expectedResult := false

	// Act
	result := isValid(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetFour(t *testing.T) {
	// Arrange
	input := "([])"
	expectedResult := true

	// Act
	result := isValid(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}
