package longestCommonPrefix

import "testing"

func TestDataSetOne(t *testing.T) {
	// Arrange
	input := []string{"flower", "flow", "flight"}
	expectedResult := "fl"

	// Act
	result := longestCommonPrefix(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}

func TestDataSetTwo(t *testing.T) {
	// Arrange
	input := []string{"dog", "racecar", "car"}
	expectedResult := ""

	// Act
	result := longestCommonPrefix(input)

	// Assert
	if result != expectedResult {
		t.Errorf(`Expected: %v, Recieved: %v`, expectedResult, result)
	}
}
