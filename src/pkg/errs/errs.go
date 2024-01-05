package errs

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// map of enums to errors.  We might want to re-use an enum for multiple
// internal errors.
var externalToInternal = map[*core.Err][]error{
	core.ErrBackupNotFound:        {repository.ErrorBackupNotFound},
	core.ErrRepoAlreadyExists:     {repository.ErrorRepoAlreadyExists},
	core.ErrResourceNotAccessible: {graph.ErrResourceLocked},
}

type ErrCheck func(error) bool

// Internal returns the internal errors and error checking functions which
// match to the public error enum.
func Internal(ce *core.Err) []error {
	return externalToInternal[ce]
}

// Is checks if the provided error contains an internal error that matches
// the public error category.
func Is(err error, ce *core.Err) bool {
	if errors.Is(err, ce) {
		return true
	}

	internalErrs, ok := externalToInternal[ce]
	if ok {
		for _, target := range internalErrs {
			if errors.Is(err, target) {
				return true
			}
		}
	}

	return false
}
