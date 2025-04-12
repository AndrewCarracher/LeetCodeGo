package validParentheses

func isValid(s string) bool {

	parenthesesMap := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	var openParenthesesOrder []byte

	strLength := len(s)

	if strLength%2 != 0 {
		return false
	}

	for i := range len(s) {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			openParenthesesOrder = append(openParenthesesOrder, s[i])
		} else {
			if len(openParenthesesOrder) == 0 || len(openParenthesesOrder) > (len(s)-i) {
				return false
			}

			if parenthesesMap[openParenthesesOrder[len(openParenthesesOrder)-1]] == s[i] {
				openParenthesesOrder = openParenthesesOrder[:len(openParenthesesOrder)-1]
			} else {
				return false
			}
		}
	}

	return len(openParenthesesOrder) == 0
}
