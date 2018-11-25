// Package hamming contains functions to compare DNA
package hamming

import "fmt"

// Distance quantifies the discrepancies between 2 DNA portions of the same length
// and returns an error if their length differs
func Distance(a, b string) (int, error) {
	if a == b {
		fmt.Printf("Distance: %v == %v\n", a, b)

		return 0, nil

	}
	if len(a) != len(b) {
		fmt.Printf("Distance: len(%v) != len(%v)\n", a, b)
		return -1, fmt.Errorf("strand lengths do not match")
	}

	diffAminoCount := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			fmt.Printf("Distance: a[%v] (%v) != b[%v] (%v)\n ", i, a[i], i, b[i])
			diffAminoCount++
		}
	}

	return diffAminoCount, nil
}
