package fuzz

import (
	"os"
)

func CheckWordlists() bool {

	files := []string{

		"tools/wordlists/directories/raft-small-directories.txt",

		"tools/wordlists/directories/raft-medium-directories.txt",

		"tools/wordlists/directories/raft-large-directories.txt",

		"tools/wordlists/files/raft-small-files.txt",

		"tools/wordlists/files/raft-medium-files.txt",

		"tools/wordlists/files/raft-large-files.txt",
	}

	for _, f := range files {

		if _, err := os.Stat(f); err != nil {

			return false

		}

	}

	return true
}
