package graph

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// Error Interpretation Helpers
// ---------------------------------------------------------------------------

const (
	errCodeActivityLimitReached        = "activityLimitReached"
	errCodeItemNotFound                = "ErrorItemNotFound"
	errCodeItemNotFoundShort           = "itemNotFound"
	errCodeEmailFolderNotFound         = "ErrorSyncFolderNotFound"
	errCodeResyncRequired              = "ResyncRequired" // alt: resyncRequired
	errCodeMalwareDetected             = "malwareDetected"
	errCodeSyncFolderNotFound          = "ErrorSyncFolderNotFound"
	errCodeSyncStateNotFound           = "SyncStateNotFound"
	errCodeResourceNotFound            = "ResourceNotFound"
	errCodeRequestResourceNotFound     = "Request_ResourceNotFound"
	errCodeMailboxNotEnabledForRESTAPI = "MailboxNotEnabledForRESTAPI"
)

const (
	mysiteURLNotFound = "unable to retrieve user's mysite URL"
	mysiteNotFound    = "user's mysite not found"
)

const (
	LabelsMalware        = "malware_detected"
	LabelsMysiteNotFound = "mysite_not_found"
)

var (
	// The folder or item was deleted between the time we identified
	// it and when we tried to fetch data for it.
	ErrDeletedInFlight = clues.New("deleted in flight")

	// Delta tokens can be desycned or expired.  In either case, the token
	// becomes invalid, and cannot be used again.
	// https://learn.microsoft.com/en-us/graph/errors#code-property
	ErrInvalidDelta = clues.New("inalid delta token")

	// Timeout errors are identified for tracking the need to retry calls.
	// Other delay errors, like throttling, are already handled by the
	// graph client's built-in retries.
	// https://github.com/microsoftgraph/msgraph-sdk-go/issues/302
	ErrTimeout = clues.New("communication timeout")
)

func IsErrDeletedInFlight(err error) bool {
	if errors.Is(err, ErrDeletedInFlight) {
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

func IsErrInvalidDelta(err error) bool {
	return hasErrorCode(err, errCodeSyncStateNotFound, errCodeResyncRequired) ||
		errors.Is(err, ErrInvalidDelta)
}

func IsErrExchangeMailFolderNotFound(err error) bool {
	return hasErrorCode(err, errCodeResourceNotFound, errCodeMailboxNotEnabledForRESTAPI)
}

func IsErrUserNotFound(err error) bool {
	return hasErrorCode(err, errCodeRequestResourceNotFound)
}

func IsErrTimeout(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		return err.Timeout()
	}

	return errors.Is(err, ErrTimeout) ||
		errors.Is(err, context.DeadlineExceeded) ||
		errors.Is(err, http.ErrHandlerTimeout) ||
		os.IsTimeout(err)
}

func IsErrUnauthorized(err error) bool {
	// TODO: refine this investigation.  We don't currently know if
	// a specific item download url expired, or if the full connection
	// auth expired.
	return clues.HasLabel(err, LabelStatus(http.StatusUnauthorized))
}

// LabelStatus transforms the provided statusCode into
// a standard label that can be attached to a clues error
// and later reviewed when checking error statuses.
func LabelStatus(statusCode int) string {
	return fmt.Sprintf("status_code_%d", statusCode)
}

// IsMalware is true if the graphAPI returns a "malware detected" error code.
func IsMalware(err error) bool {
	return hasErrorCode(err, errCodeMalwareDetected)
}

func IsMalwareResp(ctx context.Context, resp *http.Response) bool {
	// https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-wsshp/ba4ee7a8-704c-4e9c-ab14-fa44c574bdf4
	// https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-wdvmoduu/6fa6d4a9-ac18-4cd7-b696-8a3b14a98291
	if resp.Header.Get("X-Virus-Infected") == "true" {
		return true
	}

	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		logger.Ctx(ctx).Errorw("dumping http response", "error", err)
		return false
	}

	if strings.Contains(string(respDump), errCodeMalwareDetected) {
		return true
	}

	return false
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

	lcodes := []string{}
	for _, c := range codes {
		lcodes = append(lcodes, strings.ToLower(c))
	}

	return slices.Contains(lcodes, strings.ToLower(*oDataError.GetError().GetCode()))
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
		err = err.Label(LabelsMysiteNotFound)
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
	msgConcat := ptr.Val(mainErr.GetMessage()) + ptr.Val(mainErr.GetCode())

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
