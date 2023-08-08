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
	RepoAlreadyExists     errEnum = "repository-already-exists"
	BackupNotFound        errEnum = "backup-not-found"
	ServiceNotEnabled     errEnum = "service-not-enabled"
	ResourceOwnerNotFound errEnum = "resource-owner-not-found"
)

// map of enums to errors.  We might want to re-use an enum for multiple
// internal errors (ex: "ServiceNotEnabled" may exist in both graph and
// non-graph producers).
var internalToExternal = map[errEnum][]error{
	RepoAlreadyExists:     {repository.ErrorRepoAlreadyExists},
	BackupNotFound:        {repository.ErrorBackupNotFound},
	ServiceNotEnabled:     {graph.ErrServiceNotEnabled},
	ResourceOwnerNotFound: {graph.ErrResourceOwnerNotFound},
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
