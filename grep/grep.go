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
	countFileName := map[string]int{}
	shouldIncludeFileName := len(files) > 1
	nFlagEnabled := false
	lFlagEnabled := false
	iFlagEnabled := false
	xFlagEnabled := false
	vFlagEnabled := false
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
		if flag == "-x" {
			xFlagEnabled = true
		}
		if flag == "-v" {
			vFlagEnabled = true
		}
	}

	if iFlagEnabled {
		pattern = strings.ToLower(pattern)
	}
	re := regexp.MustCompile(pattern)

	for _, fileName := range files {
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		splitContent := strings.Split(string(content), "\n")

		for i, line := range splitContent {
			if line == "" {
				continue //you skip but you don't exit
			}

			lineToMatch := line
			if iFlagEnabled {
				lineToMatch = strings.ToLower(line)
			}
			a := re.FindString(lineToMatch)
			isMatch := a != ""
			if xFlagEnabled && len(a) != len(lineToMatch) {
				isMatch = false
			}

			if vFlagEnabled {
				isMatch = !isMatch
			}

			if isMatch {

				if lFlagEnabled {
					//before the ; it looks up the fileName in the map countFileName and it uses 2 return values
					//1 is the value corresponding to the key in the map and the other is a bool describing if
					//the key is the map. After the ; we do a test (!ok)
					if _, ok := countFileName[fileName]; !ok {
						lineMatches = append(lineMatches, fileName)
						countFileName[fileName] = 1
					}
				} else {
					prefix := ""
					if nFlagEnabled {
						prefix = fmt.Sprintf("%v:", i+1)
					}
					if shouldIncludeFileName {
						prefix = fmt.Sprintf("%v:%v", fileName, prefix)
					}
					fmt.Printf("Appending '%v'\n", line)
					lineMatches = append(lineMatches, fmt.Sprintf("%v%v", prefix, line))
				}
			}

		}
	}
	return lineMatches

}
