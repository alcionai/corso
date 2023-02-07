package common

import (
	"fmt"
	"io"
)

// TODO: Remove in favor of clues.Stack()

// Err provides boiler-plate functions that other types of errors can use
// if they wish to be compared with `errors.As()`. This struct ensures that
// stack traces are printed when requested (if present) and that Err
// chains `errors.As()`, `errors.Is()`, and `errors.Cause()` calls properly.
//
// When using errors.As, note that the variable that is passed as the second
// parameter must be a pointer to a type that exactly matches the returned type of the error previously. For
// example, if a struct was returned, the second parameter should be a pointer
// to said struct. If a pointer to a struct was returned, then a pointer to a
// pointer of the struct should be passed.
type Err struct {
	Err error
}

func EncapsulateError(e error) *Err {
	return &Err{Err: e}
}

func (e Err) Error() string {
	return e.Err.Error()
}

func (e Err) Cause() error {
	return e.Err
}

func (e Err) Unwrap() error {
	return e.Err
}

// Format complies with the Formatter interface and gives pretty printing when
// functions like `fmt.Printf("%+v")` are called. Implementing this allows Err
// to print stack traces from the encapsulated error.
func (e Err) Format(s fmt.State, verb rune) {
	if f, ok := e.Err.(fmt.Formatter); ok {
		f.Format(s, verb)
		return
	}

	// Formatting magic courtesy of github.com/pkg/errors.
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", e.Cause())
			return
		}

		fallthrough
	case 's', 'q':
		// nolint:errcheck
		_, _ = io.WriteString(s, e.Error())
	}
}
