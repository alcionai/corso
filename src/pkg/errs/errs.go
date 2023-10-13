package errs

import (
	"errors"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/repository"
)

// expose enums, rather than errors, for Is checks.  The enum should
// map to a specific internal error that can be used for the actual
// errors.Is comparison.
type errEnum string

const (
	ApplicationThrottled  errEnum = "application-throttled"
	BackupNotFound        errEnum = "backup-not-found"
	RepoAlreadyExists     errEnum = "repository-already-exists"
	ResourceNotAccessible errEnum = "resource-not-accesible"
	ResourceOwnerNotFound errEnum = "resource-owner-not-found"
	ServiceNotEnabled     errEnum = "service-not-enabled"
)

// map of enums to errors.  We might want to re-use an enum for multiple
// internal errors (ex: "ServiceNotEnabled" may exist in both graph and
// non-graph producers).
var internalToExternal = map[errEnum][]error{
	ApplicationThrottled:  {graph.ErrApplicationThrottled},
	BackupNotFound:        {repository.ErrorBackupNotFound},
	RepoAlreadyExists:     {repository.ErrorRepoAlreadyExists},
	ResourceNotAccessible: {graph.ErrResourceLocked},
	ResourceOwnerNotFound: {graph.ErrResourceOwnerNotFound},
	ServiceNotEnabled:     {graph.ErrServiceNotEnabled},
}

// Internal returns the internal errors which match to the public error category.
func Internal(enum errEnum) []error {
	return internalToExternal[enum]
}

// Is checks if the provided error contains an internal error that matches
// the public error category.
func Is(err error, enum errEnum) bool {
	internalErrs, ok := internalToExternal[enum]
	if !ok {
		return false
	}

	for _, target := range internalErrs {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}
