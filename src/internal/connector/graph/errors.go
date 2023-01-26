package graph

import (
	"context"
	"net/url"
	"os"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// Error Interpretation Helpers
// ---------------------------------------------------------------------------

const (
	errCodeItemNotFound                = "ErrorItemNotFound"
	errCodeEmailFolderNotFound         = "ErrorSyncFolderNotFound"
	errCodeResyncRequired              = "ResyncRequired"
	errCodeSyncFolderNotFound          = "ErrorSyncFolderNotFound"
	errCodeSyncStateNotFound           = "SyncStateNotFound"
	errCodeResourceNotFound            = "ResourceNotFound"
	errCodeMailboxNotEnabledForRESTAPI = "MailboxNotEnabledForRESTAPI"
)

var (
	err401Unauthorized = errors.New("401 unauthorized intercepted")
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

	if hasErrorCode(err, errCodeItemNotFound, errCodeSyncFolderNotFound) {
		return ErrDeletedInFlight{*common.EncapsulateError(err)}
	}

	return nil
}

func asDeletedInFlight(err error) bool {
	e := ErrDeletedInFlight{}
	return errors.As(err, &e)
}

// Delta tokens can be desycned or Unauthorized.  In either case, the token
// becomes invalid, and cannot be used again.
// https://learn.microsoft.com/en-us/graph/errors#code-property
type ErrInvalidDelta struct {
	common.Err
}

func IsErrInvalidDelta(err error) error {
	if asInvalidDelta(err) {
		return err
	}

	if hasErrorCode(err, errCodeSyncStateNotFound, errCodeResyncRequired) {
		return ErrInvalidDelta{*common.EncapsulateError(err)}
	}

	return nil
}

func asInvalidDelta(err error) bool {
	e := ErrInvalidDelta{}
	return errors.As(err, &e)
}

func IsErrExchangeMailFolderNotFound(err error) bool {
	return hasErrorCode(err, errCodeResourceNotFound, errCodeMailboxNotEnabledForRESTAPI)
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

// isTimeoutErr is used to determine if the Graph error returned is
// because of Timeout. This is used to restrict retries to just
// timeouts as other errors are handled within a middleware in the
// client.
func isTimeoutErr(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) || os.IsTimeout(err) {
		return true
	}

	switch err := err.(type) {
	case *url.Error:
		return err.Timeout()
	default:
		return false
	}
}

type ErrUnauthorized struct {
	common.Err
}

func IsErrUnauthorized(err error) error {
	// TODO: refine this investigation.  We don't currently know if
	// a specific item download url expired, or if the full connection
	// auth expired.
	if errors.Is(err, err401Unauthorized) {
		return err
	}

	if asUnauthorized(err) {
		return err
	}

	if hasErrorCode(err, "foo") {
		return ErrInvalidDelta{*common.EncapsulateError(err)}
	}

	return nil
}

func asUnauthorized(err error) bool {
	e := ErrUnauthorized{}
	return errors.As(err, &e)
}

// ---------------------------------------------------------------------------
// error parsers
// ---------------------------------------------------------------------------

func hasErrorCode(err error, codes ...string) bool {
	if err == nil {
		return false
	}

	var oDataError *odataerrors.ODataError
	if !errors.As(err, &oDataError) {
		return false
	}

	if oDataError.GetError().GetCode() == nil {
		return false
	}

	logger.Ctx(context.Background()).Errorw("ERR CODE", "code", *oDataError.GetError().GetCode())

	return slices.Contains(codes, *oDataError.GetError().GetCode())
}
