package core

import "errors"

// -----------------------------------------------------------------------------------------------
// core.Err sentinels are provided to maintain a reference of commonplace errors throughout Corso.
//
// The general idea is that these errors allow the repo (and consumers of its CLI and SDK apis)
// to communicate clearly about the central identity of an error (ie: its "core"), without leaking
// service-specific details and imports from low-level apis, clients, and other packages.
//
// In order to maintain sanity here, a couple rules should be followed.
//
// 1. Sentinels should have generic messages. No references to downstream concepts.
//		Basic cleanliness here.  Downstream references contaminate the sentinel purpose.
//
// 2. Maintain coarseness.
//		We won't need a core.Err version of every lower-level error.  Try, where possible,
//		to group concepts into broad categories.  Ex: prefer "resource not found" over
//		"user not found" or "site not found".
//
// 3. Always Stack/Wrap core.Errs.  Only once.
//		`return core.ErrFoo` should be avoided.  Also, if you're handling a error returned
//		by some internal package, do your due diligence and make sure it isn't already
//		identified by a core.Err at a lower level.
//
// 4. Stacking/Wrapping is the lowest layer's job.
//		We prefer to returning sentinels at lower layers instead of parsing errors at
//		higher layers.  This ensures higher layers only need to run errors.Is and .As
//		checks, without needing take on low-level error details.
//
// 5. Add comments to explain the sentinels.
//		Future maintainers may not easily grok the intent behind an existing sentinel.
//		Because we want to keep the error messages themselves small and clean, a short
//		explanation in the comments, even a basic one, can help a lot.
//
// 6. This package gets more important at higher layers.
//		The goal is to make life easier for layers that are the most detached from low-
//		level and internal packages.  The closer that code gets to those lower layers,
//		the less important it is to strictly use this package.  But since most errors
// 		bubble up to the SDK and CLI APIs, it is eventually a critical issue that we
//		categorize our errors smartly for those end users.
// -----------------------------------------------------------------------------------------------

type Err struct {
	msg string
}

func (e Err) Error() string {
	return e.msg
}

var (
	// currently we have no internal throttling controls.  We only try to match
	// external throttling requirements.  This sentinel assumes that an external
	// server has returned one or more throttling errors which has stopped
	// operation progress.
	ErrApplicationThrottled = &Err{msg: "application throttled"}
	// for use when a short-lived auth token (a jwt or something similar) expires.
	ErrAuthTokenExpired = &Err{msg: "auth token expired"}
	// about what it sounds like: we tried to look for a backup by ID, but the
	// storage layer couldn't find anything for that ID.
	ErrBackupNotFound = &Err{msg: "backup not found"}
	// occurs when creation of an entity (usually by restful POST or PUT) errors
	// because some other entity already already exists with a conflicting identifier.
	// The identifier is not always the id.  For example: duplicate filenames
	// in the same directory will cause conflicts, even with different IDs.
	ErrConflictAlreadyExists = &Err{msg: "conflict: already exists"}
	// a catch-all for downstream api auth issues.  doesn't matter which api.
	ErrInsufficientAuthorization = &Err{msg: "insufficient authorization"}
	// basically what it sounds like: we went looking for something by ID and
	// it wasn't found.  This might be because it was deleted in flight, or
	// was never created, or some other reason.
	ErrNotFound = &Err{msg: "not found"}
	// specifically for repository creation: if we tried to create a repo and
	// it already exists with those credentials, we return this error.
	ErrRepoAlreadyExists = &Err{msg: "repository already exists"}
	// use this when a resource (user, etc; whatever owner is used to own the
	// data in the given backup) is unable to be used for backup or restore.
	// some nuance here: this is not the same as a broad-scale auth issue.
	// it is also not the same as a "not found" issue.  it's specific to
	// cases where we can find the resource, and have authorization to access
	// it, but are told by the external system that the resource is somehow
	// unusable.
	ErrResourceNotAccessible = &Err{msg: "resource not accesible"}
	// use this when a resource (user, etc; whatever owner is used to own the
	// data in the given backup) cannot be found in the system by the ID that
	// the end user provided.
	ErrResourceOwnerNotFound = &Err{msg: "resource owner not found"}
	// a service is the set of application data within a given provider.  eg:
	// if m365 is the provider, then exchange is a service, so is oneDrive.
	// this sentinel is used to indicate that the service in question is not
	// accessible to the user.  this is not the same as an auth error.  more
	// often its a license issue.  as in: the tenant hasn't purchased the use
	// of this service (but may have purchased the use of other services in
	// the same provider).
	ErrServiceNotEnabled = &Err{msg: "service not enabled"}
)

// As is a quality-of-life wrapper around errors.As, to retrieve the core.Err
// out of any arbitrary error.
func As(err error) (*Err, bool) {
	if err == nil {
		return nil, false
	}

	var (
		ce *Err
		ok = errors.As(err, &ce)
	)

	if !ok {
		return nil, ok
	}

	return ce, ok
}
