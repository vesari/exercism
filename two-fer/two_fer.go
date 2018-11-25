

// Package twofer refers to "One for you, one for me"
package twofer
import "fmt"

// ShareWith returns either "you" or a given name in the twofer context
func ShareWith(name string) string {
	subject := "you"
	if name != "" {
 		subject = name
	}
	
	return fmt.Sprintf("One for %v, one for me.", subject)
}
