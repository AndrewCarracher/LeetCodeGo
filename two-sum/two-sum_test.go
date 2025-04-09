package twoSum

import "testing"

func TestDataSetOne(t *testing.T) {
	// Arrange
	var testData []int
	var expectedResult []int
	testData = append(testData, 2, 7, 11, 15)
	// expectedResult = append(expectedResult, 0, 1)
	expectedResult = append(expectedResult, 0, 2)

	// Act
	result := twoSum(testData, 9)

	// Assert
	if result[0] != expectedResult[0] || result[1] != expectedResult[1] {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetTwo(t *testing.T) {
	// Arrange
	var testData []int
	var expectedResult []int
	testData = append(testData, 3, 2, 4)
	expectedResult = append(expectedResult, 1, 2)

	// Act
	result := twoSum(testData, 6)

	// Assert
	if result[0] != expectedResult[0] || result[1] != expectedResult[1] {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetThree(t *testing.T) {
	// Arrange
	var testData []int
	var expectedResult []int
	testData = append(testData, 3, 3)
	expectedResult = append(expectedResult, 0, 1)

	// Act
	result := twoSum(testData, 6)

	// Assert
	if result[0] != expectedResult[0] || result[1] != expectedResult[1] {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}
