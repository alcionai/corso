package graph

import (
	"context"
	"net/url"
	"os"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common"
)

// ---------------------------------------------------------------------------
// Error Interpretation Helpers
// ---------------------------------------------------------------------------

const (
	errCodeActivityLimitReached        = "activityLimitReached"
	errCodeItemNotFound                = "ErrorItemNotFound"
	errCodeEmailFolderNotFound         = "ErrorSyncFolderNotFound"
	errCodeResyncRequired              = "ResyncRequired"
	errCodeSyncFolderNotFound          = "ErrorSyncFolderNotFound"
	errCodeSyncStateNotFound           = "SyncStateNotFound"
	errCodeResourceNotFound            = "ResourceNotFound"
	errCodeMailboxNotEnabledForRESTAPI = "MailboxNotEnabledForRESTAPI"
	tooManyRequests                    = 429
	serviceUnavailable                 = 503
	gatewayTimeout                     = 504
	internalServerError                = 500
)

var (
	Err401Unauthorized = errors.New("401 unauthorized")
	// normally the graph client will catch this for us, but in case we
	// run our own client Do(), we need to translate it to a timeout type
	// failure locally.
	Err429TooManyRequests     = errors.New("429 too many requests")
	Err503ServiceUnavailable  = errors.New("503 Service Unavailable")
	Err504GatewayTimeout      = errors.New("504 Gateway Timeout")
	Err500InternalServerError = errors.New("500 Internal Server Error")
)

// The folder or item was deleted between the time we identified
// it and when we tried to fetch data for it.
type ErrDeletedInFlight struct {
	common.Err
}

func IsErrDeletedInFlight(err error) bool {
	e := ErrDeletedInFlight{}
	if errors.As(err, &e) {
		return true
	}

	if hasErrorCode(err, errCodeItemNotFound, errCodeSyncFolderNotFound) {
		return true
	}

	return false
}

// Delta tokens can be desycned or expired.  In either case, the token
// becomes invalid, and cannot be used again.
// https://learn.microsoft.com/en-us/graph/errors#code-property
type ErrInvalidDelta struct {
	common.Err
}

func IsErrInvalidDelta(err error) bool {
	e := ErrInvalidDelta{}
	if errors.As(err, &e) {
		return true
	}

	if hasErrorCode(err, errCodeSyncStateNotFound, errCodeResyncRequired) {
		return true
	}

	return false
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

func IsErrTimeout(err error) bool {
	e := ErrTimeout{}
	if errors.As(err, &e) {
		return true
	}

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

type ErrThrottled struct {
	common.Err
}

func IsErrThrottled(err error) bool {
	if errors.Is(err, Err429TooManyRequests) {
		return true
	}

	if hasErrorCode(err, errCodeActivityLimitReached) {
		return true
	}

	e := ErrThrottled{}

	return errors.As(err, &e)
}

type ErrUnauthorized struct {
	common.Err
}

func IsErrUnauthorized(err error) bool {
	// TODO: refine this investigation.  We don't currently know if
	// a specific item download url expired, or if the full connection
	// auth expired.
	if errors.Is(err, Err401Unauthorized) {
		return true
	}

	e := ErrUnauthorized{}

	return errors.As(err, &e)
}

type ErrInternalServerError struct {
	common.Err
}

func IsInternalServerError(err error) bool {
	if errors.Is(err, Err500InternalServerError) {
		return true
	}

	e := ErrInternalServerError{}

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

	return slices.Contains(codes, *oDataError.GetError().GetCode())
}
