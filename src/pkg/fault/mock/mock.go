package mock

import "github.com/alcionai/corso/src/pkg/fault"

// Adder mocks an adder interface for testing.
type Adder struct {
	Errs []error
}

func NewAdder() *Adder {
	return &Adder{Errs: []error{}}
}

func (ma *Adder) Add(err error) *fault.Errors {
	ma.Errs = append(ma.Errs, err)
	return fault.New(true)
}
