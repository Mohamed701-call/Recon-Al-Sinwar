package common

import (
	"strings"
)

func SplitOutput(output string) []string {

	lines := strings.Split(output, "\n")

	var result []string

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		result = append(result, line)
	}

	return result
}
