package strand

import (
	"strings"
)

// ToRNA is a function that converts DNA strands into RNA ones.
func ToRNA(dna string) string {
	s := string(dna)
	sSep := strings.Split(s, "")
	sequence := make([]string, 0)

	for _, v := range sSep {
		if v == "G" {
			sequence = append(sequence, "C")
		}
		if v == "C" {
			sequence = append(sequence, "G")
		}
		if v == "T" {
			sequence = append(sequence, "A")
		}
		if v == "A" {
			sequence = append(sequence, "U")
		}
	}

	return strings.Join(sequence, "")
}
