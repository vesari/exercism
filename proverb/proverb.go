
package proverb

import (
	"fmt"
)

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	saying := make([]string, 0, len(rhyme))
	for i := 1; i < len(rhyme); i++ {
		saying = append(saying, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i - 1], rhyme[i]))
	}
	saying = append(saying, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return saying
}
