package types

import "time"

type Result struct {
	Module string

	OutputFile string

	Count int

	Started time.Time

	Finished time.Time

	Success bool

	Err error
}
