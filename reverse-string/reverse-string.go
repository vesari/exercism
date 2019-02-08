package reverse

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//String reverses a string even if it's multi byte
func String(initial string) string {
	splitInitial := strings.Split(initial, "")
	reversed := []string{}
	for i := len(splitInitial) - 1; i >= 0; i-- {
		v := splitInitial[i]
		fmt.Printf("")
		for len(v) > 0 {
			r, size := utf8.DecodeRuneInString(v)
			v = v[size:]
			reversed = append(reversed, string(r))
		}
	}
	return strings.Join(reversed, "")
}
