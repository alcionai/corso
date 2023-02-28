package graph

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
)

// ---------------------------------------------------------------------------
// Error Interpretation Helpers
// ---------------------------------------------------------------------------

const (
	errCodeActivityLimitReached        = "activityLimitReached"
	errCodeItemNotFound                = "ErrorItemNotFound"
	errCodeItemNotFoundShort           = "itemNotFound"
	errCodeEmailFolderNotFound         = "ErrorSyncFolderNotFound"
	errCodeResyncRequired              = "ResyncRequired"
	errCodeSyncFolderNotFound          = "ErrorSyncFolderNotFound"
	errCodeSyncStateNotFound           = "SyncStateNotFound"
	errCodeResourceNotFound            = "ResourceNotFound"
	errCodeRequestResourceNotFound     = "Request_ResourceNotFound"
	errCodeMailboxNotEnabledForRESTAPI = "MailboxNotEnabledForRESTAPI"
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

var (
	mysiteURLNotFound = "unable to retrieve user's mysite URL"
	mysiteNotFound    = "user's mysite not found"
)

var Labels = struct {
	MysiteNotFound string
}{
	MysiteNotFound: "mysite_not_found",
}

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

	if hasErrorCode(
		err,
		errCodeItemNotFound,
		errCodeItemNotFoundShort,
		errCodeSyncFolderNotFound,
	) {
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

func IsErrUserNotFound(err error) bool {
	return hasErrorCode(err, errCodeRequestResourceNotFound)
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

	if errors.Is(err, context.DeadlineExceeded) || os.IsTimeout(err) || errors.Is(err, http.ErrHandlerTimeout) {
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

// Wrap is a helper function that extracts ODataError metadata from
// the error.  If the error is not an ODataError type, returns the error.
func Wrap(ctx context.Context, e error, msg string) *clues.Err {
	if e == nil {
		return nil
	}

	odErr, ok := e.(odataerrors.ODataErrorable)
	if !ok {
		return clues.Wrap(e, msg).WithClues(ctx)
	}

	data, innerMsg := errData(odErr)

	return setLabels(clues.Wrap(e, msg).WithClues(ctx).With(data...), innerMsg)
}

// Stack is a helper function that extracts ODataError metadata from
// the error.  If the error is not an ODataError type, returns the error.
func Stack(ctx context.Context, e error) *clues.Err {
	if e == nil {
		return nil
	}

	odErr, ok := e.(odataerrors.ODataErrorable)
	if !ok {
		return clues.Stack(e).WithClues(ctx)
	}

	data, innerMsg := errData(odErr)

	return setLabels(clues.Stack(e).WithClues(ctx).With(data...), innerMsg)
}

func setLabels(err *clues.Err, msg string) *clues.Err {
	if strings.Contains(msg, mysiteNotFound) || strings.Contains(msg, mysiteURLNotFound) {
		err = err.Label(Labels.MysiteNotFound)
	}

	return err
}

func errData(err odataerrors.ODataErrorable) ([]any, string) {
	data := make([]any, 0)

	// Get MainError
	mainErr := err.GetError()

	data = appendIf(data, "odataerror_code", mainErr.GetCode())
	data = appendIf(data, "odataerror_message", mainErr.GetMessage())
	data = appendIf(data, "odataerror_target", mainErr.GetTarget())
	msgConcat := ptr.Val(mainErr.GetMessage())

	for i, d := range mainErr.GetDetails() {
		pfx := fmt.Sprintf("odataerror_details_%d_", i)
		data = appendIf(data, pfx+"code", d.GetCode())
		data = appendIf(data, pfx+"message", d.GetMessage())
		data = appendIf(data, pfx+"target", d.GetTarget())
		msgConcat += ptr.Val(d.GetMessage())
	}

	inner := mainErr.GetInnererror()
	if inner != nil {
		data = appendIf(data, "odataerror_inner_cli_req_id", inner.GetClientRequestId())
		data = appendIf(data, "odataerror_inner_req_id", inner.GetRequestId())
	}

	return data, strings.ToLower(msgConcat)
}

func appendIf(a []any, k string, v *string) []any {
	if v == nil {
		return a
	}

	return append(a, k, *v)
}
