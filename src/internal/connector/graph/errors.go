package graph

import (
	"fmt"
	"net/url"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
)

// ---------------------------------------------------------------------------
// Error Interpretation Helpers
// ---------------------------------------------------------------------------

const (
	errCodeItemNotFound        = "ErrorItemNotFound"
	errCodeEmailFolderNotFound = "ErrorSyncFolderNotFound"
	errCodeResyncRequired      = "ResyncRequired"
	errCodeSyncStateNotFound   = "SyncStateNotFound"
)

// The folder or item was deleted between the time we identified
// it and when we tried to fetch data for it.
type ErrDeletedInFlight struct {
	common.Err
}

func IsErrDeletedInFlight(err error) error {
	if asDeletedInFlight(err) {
		return err
	}

	if hasErrorCode(err, errCodeItemNotFound) {
		return ErrDeletedInFlight{*common.EncapsulateError(err)}
	}

	return nil
}

func asDeletedInFlight(err error) bool {
	e := ErrDeletedInFlight{}
	return errors.As(err, &e)
}

// Delta tokens can be desycned or expired.  In either case, the token
// becomes invalid, and cannot be used again.
// https://learn.microsoft.com/en-us/graph/errors#code-property
type ErrInvalidDelta struct {
	common.Err
}

func IsErrInvalidDelta(err error) error {
	if asInvalidDelta(err) {
		return err
	}

	if hasErrorCode(err, errCodeSyncStateNotFound) ||
		hasErrorCode(err, errCodeResyncRequired) {
		return ErrInvalidDelta{*common.EncapsulateError(err)}
	}

	return nil
}

func asInvalidDelta(err error) bool {
	e := ErrInvalidDelta{}
	return errors.As(err, &e)
}

// Timeout errors are identified for tracking the need to retry calls.
// Other delay errors, like throttling, are already handled by the
// graph client's built-in retries.
// https://github.com/microsoftgraph/msgraph-sdk-go/issues/302
type ErrTimeout struct {
	common.Err
}

func IsErrTimeout(err error) error {
	if asTimeout(err) {
		return err
	}

	if isTimeoutErr(err) {
		return ErrTimeout{*common.EncapsulateError(err)}
	}

	return nil
}

func asTimeout(err error) bool {
	e := ErrTimeout{}
	return errors.As(err, &e)
}

// ---------------------------------------------------------------------------
// error parsers
// ---------------------------------------------------------------------------

func hasErrorCode(err error, code string) bool {
	if err == nil {
		fmt.Println("nil")
		return false
	}

	var oDataError *odataerrors.ODataError
	if !errors.As(err, &oDataError) {
		return false
	}

	return oDataError.GetError().GetCode() != nil &&
		*oDataError.GetError().GetCode() == code
}

// isTimeoutErr is used to determine if the Graph error returned is
// because of Timeout. This is used to restrict retries to just
// timeouts as other errors are handled within a middleware in the
// client.
func isTimeoutErr(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		return err.Timeout()
	default:
		return false
	}
}
