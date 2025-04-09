package romanToInt

import "testing"

func TestDataSetOne(t *testing.T) {
	// Arrange
	input := "III"
	expectedResult := 3

	// Act
	result := romanToInt(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetTwo(t *testing.T) {
	// Arrange
	input := "LVIII"
	expectedResult := 58

	// Act
	result := romanToInt(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetThree(t *testing.T) {
	// Arrange
	input := "MCMXCIV"
	expectedResult := 1994

	// Act
	result := romanToInt(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}
