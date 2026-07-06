package types

type Module interface {
	Name() string
	Run() (*Result, error)
}
