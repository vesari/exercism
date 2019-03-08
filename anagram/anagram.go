package anagram

import (
	"sort"
	"strings"
)

// Detect detects anagrams to a word from a list of possible candidates
func Detect(word string, candidates []string) []string {

	lowWord := strings.ToLower(word)
	wordSlice := []string{}
	matches := []string{}

	for _, char := range strings.Split(lowWord, "") {
		wordSlice = append(wordSlice, char)
	}

	sort.Strings(wordSlice)

	for _, candidate := range candidates {
		if len(candidate) != len(word) || strings.ToLower(candidate) == lowWord {
			continue
		}

		slicedCandidate := []string{}
		slicedCandidate = append(slicedCandidate, candidate)
		reslicedCandidate := []string{}

		for _, v := range strings.Split(strings.Join(slicedCandidate, ""), "") {
			reslicedCandidate = append(reslicedCandidate, strings.ToLower(v))
			sort.Strings(reslicedCandidate)
		}

		if strings.Join(reslicedCandidate, "") == strings.Join(wordSlice, "") {
			matches = append(matches, strings.Join(slicedCandidate, ""))
		}
	}
	return matches
}
