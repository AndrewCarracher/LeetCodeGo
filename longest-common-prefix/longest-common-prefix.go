package longestCommonPrefix

import (
	"bytes"
	"slices"
)

func longestCommonPrefix(strs []string) string {

	result := bytes.NewBufferString("")
	slices.Sort(strs)
	firstElement, lastElement := strs[0], strs[len(strs)-1]

	for i := range len(firstElement) {
		if i == len(lastElement) || i == len(firstElement) {
			break
		}

		if firstElement[i] != lastElement[i] {
			break
		}

		result.WriteByte(firstElement[i])
	}

	return result.String()
}
