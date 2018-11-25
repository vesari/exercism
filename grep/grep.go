package grep

import (
	"fmt"
)

func createFiles(pattern string, flag []string, file []string) []string {
	if pattern != "" && flag == []
	// if len(pattern) == 0 {
	// 	return []string{}
	// }
	output := make([]string, 0, len(pattern))
	for i := 1; i < len(pattern); i++ {
		output = append(output, fmt.Sprintf("%s ", pattern[i-1], pattern[i]))
	}
	output = append(output, fmt.Sprintf(" %s.", pattern[0]))
	return output
}
