package core

import (
	"bytes"
	"os/exec"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Run(name string, args ...string) (string, error) {

	cmd := exec.Command(name, args...)

	var out bytes.Buffer
	var errb bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &errb

	err := cmd.Run()

	if err != nil {
		return errb.String(), err
	}

	return out.String(), nil
}
