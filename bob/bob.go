// Package bob contains the functions describing Bob's behaviour
package bob

import (
	"strings"
	"unicode"
)

// The Hey function describes how Bob would typically answer
func Hey(remark string) string {
	remarkTrimmed := strings.TrimSpace(remark)
	if remarkTrimmed == "" {
		return "Fine. Be that way!"
	}

	if remarkTrimmed[len(remarkTrimmed)-1] == '?' {
		if isUppercase(remarkTrimmed) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	} else {
		if isUppercase(remarkTrimmed) {
			return "Whoa, chill out!"
		} else {
			return "Whatever."
		}
	}
}

func isUppercase(remark string) bool {
	letterCount := 0
	upperCaseLetterCount := 0
	for _, c := range remark {
		if unicode.IsLetter(c) {
			letterCount++
			if unicode.IsUpper(c) {
				upperCaseLetterCount++
			}

		}
	}
	return letterCount == upperCaseLetterCount && letterCount > 0
}
