package indenter

import (
	"fmt"
	"regexp"

	"github.com/eryljoy/veritas/src/scoper/internal/constants"
)

func CorrectIndentation(file string) (string, error) {
	// indentLevel := 0
	var result string
	r, err := regexp.Compile(constants.Indent)
	if err != nil {
		return "", err
	}
	for {
		stringMatch := r.FindString(file)
		if stringMatch == "" {
			break
		}
		indexMatch := r.FindStringIndex(file)
		file = file[indexMatch[1]:]
		result = result + stringMatch
	}

	fmt.Printf("%v", result)

	return fmt.Sprintf("%v", result), nil
}
