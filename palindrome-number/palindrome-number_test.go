package palindromeNumber

import "testing"

func TestDataSetOne(t *testing.T) {
	// Arrange
	input := 0
	expectedResult := true

	// Act
	result := isPalindrome(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetTwo(t *testing.T) {
	// Arrange
	input := 1
	expectedResult := true

	// Act
	result := isPalindrome(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetThree(t *testing.T) {
	// Arrange
	input := 121
	expectedResult := true

	// Act
	result := isPalindrome(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetFour(t *testing.T) {
	// Arrange
	input := -121
	expectedResult := false

	// Act
	result := isPalindrome(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetFive(t *testing.T) {
	// Arrange
	input := 10
	expectedResult := false

	// Act
	result := isPalindrome(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetSix(t *testing.T) {
	// Arrange
	input := 278872
	expectedResult := true

	// Act
	result := isPalindrome(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetSeven(t *testing.T) {
	// Arrange
	input := 2780872
	expectedResult := true

	// Act
	result := isPalindrome(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}
