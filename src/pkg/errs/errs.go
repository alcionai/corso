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

// map of enums to error comparators.  The above map assumes that we
// always stack or wrap the sentinel error in the returned error.  But in
// many places of error handling, we primarily rely on error comparison
// checks.  This allows us to apply those comparison checks instead of relying
// only on sentinels.
var externalToInternalCheck = map[*core.Err][]ErrCheck{
	core.ErrResourceOwnerNotFound: {graph.IsErrItemNotFound},
}

// Internal returns the internal errors and error checking functions which
// match to the public error enum.
func Internal(ce *core.Err) ([]error, []ErrCheck) {
	return externalToInternal[ce], externalToInternalCheck[ce]
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

	internalChecks, ok := externalToInternalCheck[ce]
	if ok {
		for _, check := range internalChecks {
			if check(err) {
				return true
			}
		}
	}

	return false
}
