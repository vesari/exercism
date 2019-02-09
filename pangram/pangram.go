package pangram

import (
	"strings"
	"unicode"
)

//IsPangram finds out whether a string contains all the 26 letters of the English alphabet
//(not-case sensitive)
func IsPangram(word string) bool {
	trimmedWord := strings.ToLower(strings.TrimSpace(word))
	count := map[rune]int{}
	if len(trimmedWord) < 26 {
		return false
	}

	for _, c := range trimmedWord {
		if !unicode.IsLetter(c) {
			continue
		}
		count[c]++
	}

	if len(count) == 26 {
		return true
	}
	return false
}
