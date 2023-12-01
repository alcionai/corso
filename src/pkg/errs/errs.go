package errs

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// expose enums, rather than errors, for Is checks.  The enum should
// map to a specific internal error that can be used for the actual
// errors.Is comparison.
type errEnum string

const (
	ApplicationThrottled      errEnum = "application-throttled"
	BackupNotFound            errEnum = "backup-not-found"
	InsufficientAuthorization errEnum = "insufficient-authorization"
	RepoAlreadyExists         errEnum = "repository-already-exists"
	ResourceNotAccessible     errEnum = "resource-not-accesible"
	ResourceOwnerNotFound     errEnum = "resource-owner-not-found"
	ServiceNotEnabled         errEnum = "service-not-enabled"
)

// map of enums to errors.  We might want to re-use an enum for multiple
// internal errors (ex: "ServiceNotEnabled" may exist in both graph and
// non-graph producers).
var externalToInternal = map[errEnum][]error{
	ApplicationThrottled:  {graph.ErrApplicationThrottled},
	BackupNotFound:        {repository.ErrorBackupNotFound},
	RepoAlreadyExists:     {repository.ErrorRepoAlreadyExists},
	ResourceNotAccessible: {graph.ErrResourceLocked},
	ResourceOwnerNotFound: {graph.ErrResourceOwnerNotFound},
	ServiceNotEnabled:     {graph.ErrServiceNotEnabled},
}

type ErrCheck func(error) bool

// map of enums to error comparators.  The above map assumes that we
// always stack or wrap the sentinel error in the returned error.  But in
// many places of error handling, we primarily rely on error comparison
// checks.  This allows us to apply those comparison checks instead of relying
// only on sentinels.
var externalToInternalCheck = map[errEnum][]ErrCheck{
	ApplicationThrottled:      {graph.IsErrApplicationThrottled},
	ResourceNotAccessible:     {graph.IsErrResourceLocked},
	ResourceOwnerNotFound:     {graph.IsErrItemNotFound},
	InsufficientAuthorization: {graph.IsErrInsufficientAuthorization},
}

// Internal returns the internal errors and error checking functions which
// match to the public error enum.
func Internal(enum errEnum) ([]error, []ErrCheck) {
	return externalToInternal[enum], externalToInternalCheck[enum]
}

// Is checks if the provided error contains an internal error that matches
// the public error category.
func Is(err error, enum errEnum) bool {
	internalErrs, ok := externalToInternal[enum]
	if ok {
		for _, target := range internalErrs {
			if errors.Is(err, target) {
				return true
			}
		}
	}

	internalChecks, ok := externalToInternalCheck[enum]
	if ok {
		for _, check := range internalChecks {
			if check(err) {
				return true
			}
		}
	}

	return false
}
