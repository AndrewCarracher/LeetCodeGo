package romanToInt

import "bytes"

func romanToInt(s string) int {

	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int

	for b := 0; b < len(s); b++ {

		edgeCases := []byte{'I', 'X', 'C'}
		charArr := []byte{s[b]}

		if bytes.Contains(edgeCases, charArr) {
			if b < len(s)-1 {
				if s[b] == 'I' {
					if s[b+1] == 'V' {
						result += 4
						b++
					} else if s[b+1] == 'X' {
						result += 9
						b++
					} else {
						result += m['I']
					}
				} else if s[b] == 'X' {
					if s[b+1] == 'L' {
						result += 40
						b++
					} else if s[b+1] == 'C' {
						result += 90
						b++
					} else {
						result += m['X']
					}
				} else if s[b] == 'C' {
					if s[b+1] == 'D' {
						result += 400
						b++
					} else if s[b+1] == 'M' {
						result += 900
						b++
					} else {
						result += m['C']
					}
				}
			} else {
				result += m[s[b]]
			}
		} else {
			// treat as normal, fetch from map
			result += m[s[b]]
		}
	}

	return result
}
