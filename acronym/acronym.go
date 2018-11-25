// Package acronym contains a function to convert a string into an acronym
package acronym

import (
	"strings"
	"regexp"
)
// Abbreviate converts a string into an acronym 
func Abbreviate(s string) string {
  sTrimmed:= strings.TrimSpace(s)
	if sTrimmed == "" {
		return ""
	}
	words := regexp.MustCompile("\\s|-").Split(s, -1)
	acronym := ""
	for _ , word:= range words {
		acronym += string(word[0])
	}
	return strings.ToUpper(acronym)
}
