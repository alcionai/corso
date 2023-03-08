package graph

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common"
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
	mysiteURLNotFound = "unable to retrieve user's mysite url"
	mysiteNotFound    = "user's mysite not found"
)

const (
	LabelsMalware        = "malware_detected"
	LabelsMysiteNotFound = "mysite_not_found"
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

// MalwareInfo gathers potentially useful information about a malware infected
// drive item, and aggregates that data into a map.
func MalwareInfo(item models.DriveItemable) map[string]any {
	m := map[string]any{}

	creator := item.GetCreatedByUser()
	if creator != nil {
		m["created_by"] = ptr.Val(creator.GetId())
	}

	lastmodder := item.GetLastModifiedByUser()
	if lastmodder != nil {
		m["last_modified_by"] = ptr.Val(lastmodder.GetId())
	}

	parent := item.GetParentReference()
	if parent != nil {
		m["container_id"] = ptr.Val(parent.GetId())
		m["container_name"] = ptr.Val(parent.GetName())
	}

	malware := item.GetMalware()
	if malware != nil {
		m["malware_desciption"] = ptr.Val(malware.GetDescription())
	}

	return m
}
