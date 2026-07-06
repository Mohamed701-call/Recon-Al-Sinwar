package common

import (
	"bufio"
	"os"
)

func AppendLine(file, line string) error {

	f, err := os.OpenFile(
		file,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return err
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	_, err = writer.WriteString(line + "\n")

	if err != nil {
		return err
	}

	return writer.Flush()
}
