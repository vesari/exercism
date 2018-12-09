package grep

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// Search reads a file and looks for a text string in its content
func Search(pattern string, flags []string, files []string) []string {
	lineMatches := []string{}
	nFlagEnabled := false
	lFlagEnabled := false
	iFlagEnabled := false
	for _, flag := range flags {
		if flag == "-n" {
			nFlagEnabled = true
		}
		if flag == "-l" {
			lFlagEnabled = true
		}
		if flag == "-i" {
			iFlagEnabled = true
		}
	}

	if iFlagEnabled {
		pattern = strings.ToLower(pattern)
	}
	re := regexp.MustCompile(pattern)

	// fmt.Printf("Length of file: %v\n", len(files))
	for _, fileName := range files {
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		splitContent := strings.Split(string(content), "\n")

		for i, line := range splitContent {
			lineToMatch := line
			if iFlagEnabled {
				lineToMatch = strings.ToLower(line)
			}
			a := re.FindString(lineToMatch)
			if a != "" {
				if lFlagEnabled {
					lineMatches = append(lineMatches, fileName)

				} else {
					prefix := ""
					if nFlagEnabled {
						prefix = fmt.Sprintf("%v:", i+1)
					}
					lineMatches = append(lineMatches, fmt.Sprintf("%v%v", prefix, line))
				}
			}

		}
	}
	return lineMatches

}
