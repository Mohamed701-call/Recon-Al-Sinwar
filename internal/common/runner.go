package common

import (
	"bytes"
	"context"
	"os/exec"
	"time"
)

type Runner struct {
	Timeout time.Duration
}

func NewRunner() *Runner {

	return &Runner{
		Timeout: 10 * time.Minute,
	}

}

func (r *Runner) Run(program string, args ...string) (string, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		r.Timeout,
	)

	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		program,
		args...,
	)

	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	return out.String(), err
}
