package scope

import (
	"bufio"
	"os"
	"strings"
)

func ReadTargets(file string) ([]string, error) {

	f, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var targets []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		targets = append(targets, line)

	}

	return targets, scanner.Err()
}
