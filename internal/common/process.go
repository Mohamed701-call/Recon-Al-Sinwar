package common

import (
	"os"
	"os/exec"
)

func FindTool(name string) (string, error) {

	local := Tool(name)

	if _, err := os.Stat(local); err == nil {
		return local, nil
	}

	return exec.LookPath(name)
}

func ToolExists(name string) bool {

	_, err := FindTool(name)

	return err == nil
}
