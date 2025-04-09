package palindromeNumber

import "strconv"

func isPalindrome(x int) bool {
	// if single digit return true
	if x >= 0 && x < 10 {
		return true
	}

	numberString := strconv.Itoa(x)

	// negative numbers can't be palindromes
	if numberString[0] == '-' {
		return false
	}

	for i := 0; i < len(numberString)/2; i++ {
		if numberString[i] != numberString[len(numberString)-1-i] {
			return false
		}
	}

	return true
}
