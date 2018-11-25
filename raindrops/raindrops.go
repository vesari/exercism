package raindrops

import (
	"strconv"
	"strings"
)

// Convert turns an amount of raindrops into "raindrops speak"
func Convert(number int) string {
	// Make an array slice
	// If 3 is a factor add Pling to the slice
	// If 5 is a factor add Plang to the slice
	// If 7 is a factor add Plong to the slice
	//
	// If list is not empty return slice converted to string
	// Else return number as string
	// version is formatted
	rain := make([]string, 0, 3)

	if number%3 == 0 {
		rain = append(rain, "Pling")

	}
	if number%5 == 0 {
		rain = append(rain, "Plang")

	}
	if number%7 == 0 {
		rain = append(rain, "Plong")

	}
	if len(rain) > 0 {
		return strings.Join(rain, "")
	}

	return strconv.Itoa(number)
}
