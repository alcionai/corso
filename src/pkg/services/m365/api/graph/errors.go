package graph

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/jwt"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

// ---------------------------------------------------------------------------
// Error Interpretation Helpers
// ---------------------------------------------------------------------------

type errorCode string

const (
	ApplicationThrottled errorCode = "ApplicationThrottled"
	// this authN error is a catch-all used by graph in a variety of cases:
	// users without licenses, bad jwts, missing account permissions, etc.
	AuthenticationError errorCode = "AuthenticationError"
	// on the other hand, authZ errors apply specifically to authenticated,
	// but unauthorized, user requests
	AuthorizationRequestDenied errorCode = "Authorization_RequestDenied"
	// cannotOpenFileAttachment happen when an attachment is
	// inaccessible. The error message is usually "OLE conversion
	// failed for an attachment."
	cannotOpenFileAttachment errorCode = "ErrorCannotOpenFileAttachment"
	emailFolderNotFound      errorCode = "ErrorSyncFolderNotFound"
	ErrorAccessDenied        errorCode = "ErrorAccessDenied"
	ErrorItemNotFound        errorCode = "ErrorItemNotFound"
	// This error occurs when an email is enumerated but retrieving it fails
	// - we believe - due to it pre-dating mailbox creation. Possible explanations
	// are mailbox creation racing with email receipt or a similar issue triggered
	// due to on-prem->M365 mailbox migration.
	ErrorInvalidRecipients errorCode = "ErrorInvalidRecipients"
	// We have seen this graph error with a 500 response while backing up email
	// messages. It implies that the email message is corrupt and cannot be read.
	// Associated error message goes like "Data is corrupt.,Invalid global object ID:"
	ErrorCorruptData errorCode = "ErrorCorruptData"
	// This error occurs when an attempt is made to create a folder that has
	// the same name as another folder in the same parent. Such duplicate folder
	// names are not allowed by graph.
	folderExists   errorCode = "ErrorFolderExists"
	invalidRequest errorCode = "invalidRequest"
	// Some datacenters are returning this when we try to get the inbox of a user
	// that doesn't exist.
	invalidUser                 errorCode = "ErrorInvalidUser"
	invalidAuthenticationToken  errorCode = "InvalidAuthenticationToken"
	ItemNotFound                errorCode = "itemNotFound"
	MailboxNotEnabledForRESTAPI errorCode = "MailboxNotEnabledForRESTAPI"
	malwareDetected             errorCode = "malwareDetected"
	// nameAlreadyExists occurs when a request with
	// @microsoft.graph.conflictBehavior=fail finds a conflicting file.
	nameAlreadyExists       errorCode = "nameAlreadyExists"
	NotAllowed              errorCode = "notAllowed"
	notFound                errorCode = "NotFound"
	noResolvedUsers         errorCode = "noResolvedUsers"
	QuotaExceeded           errorCode = "ErrorQuotaExceeded"
	RequestResourceNotFound errorCode = "Request_ResourceNotFound"
	// Returned when we try to get the inbox of a user that doesn't exist.
	ResourceNotFound   errorCode = "ResourceNotFound"
	resyncRequired     errorCode = "ResyncRequired"
	syncFolderNotFound errorCode = "ErrorSyncFolderNotFound"
	syncStateInvalid   errorCode = "SyncStateInvalid"
	syncStateNotFound  errorCode = "SyncStateNotFound"
)

// inner error codes
const (
	ResourceLocked errorCode = "resourceLocked"
	// Returned when either the tenant-wide or site settings don't permit sharing
	// for external users to some extent.
	sharingDisabled errorCode = "sharingDisabled"
)

type errorMessage string

const (
	IOErrDuringRead                 errorMessage = "IO error during request payload read"
	MysiteURLNotFound               errorMessage = "unable to retrieve user's mysite url"
	MysiteNotFound                  errorMessage = "user's mysite not found"
	NoSPLicense                     errorMessage = "Tenant does not have a SPO license"
	ParameterDeltaTokenNotSupported errorMessage = "Parameter 'DeltaToken' not supported for this request"
	usersCannotBeResolved           errorMessage = "One or more users could not be resolved"
	requestedSiteCouldNotBeFound    errorMessage = "Requested site could not be found"
)

const (
	LabelsMalware             = "malware_detected"
	LabelsMysiteNotFound      = "mysite_not_found"
	LabelsNoSharePointLicense = "no_sharepoint_license"

	// LabelsSkippable is used to determine if an error is skippable
	LabelsSkippable = "skippable_errors"
)

// ErrServiceUnavailableEmptyResp indicates the remote service returned a 503
// with an empty response body. This can sometimes happen if a request times out
// during processing.
//
// TODO(ashmrtn): Either make a separate error struct for empty responses and
// implement Is() on it or start using tags on errors for the different status
// codes.
var (
	ErrServiceUnavailableEmptyResp = clues.New("service unavailable and no returned content")
	ErrNotFoundEmptyResp           = clues.New("not found and no returned content")
)

// ---------------------------------------------------------------------------
// error categorization
// ---------------------------------------------------------------------------

func stackWithCoreErr(ctx context.Context, err error, traceDepth int) error {
	if err == nil {
		return nil
	}

	ode := parseODataErr(err)
	labels := []string{}

	switch {
	case isErrBadJWTToken(ode, err):
		err = clues.Stack(core.ErrAuthTokenExpired)
	case isErrApplicationThrottled(ode, err):
		err = clues.Stack(core.ErrApplicationThrottled, err)
	case isErrUserNotFound(ode, err):
		err = clues.Stack(core.ErrNotFound, err)
	case isErrResourceLocked(ode, err):
		err = clues.Stack(core.ErrResourceNotAccessible, err)
	case isErrInsufficientAuthorization(ode, err):
		err = clues.Stack(core.ErrInsufficientAuthorization, err)
	case isErrItemAlreadyExists(ode, err):
		err = clues.Stack(core.ErrAlreadyExists, err)
	case isErrNotFound(ode, err):
		err = clues.Stack(core.ErrNotFound, err)
	default:
		err = toErrByRespCode(ode, err)

		labels = append(labels, core.LabelRootCauseUnknown)
	}

	stacked := stackWithDepth(ctx, err, 1+traceDepth).
		Label(LabelStatus(ode.Resp.StatusCode))

	// labeling here because we want the context from stackWithDepth first
	for _, label := range labels {
		stacked = stacked.Label(label)
	}

	return stacked
}

// unexported categorizers, for use with stackWithCoreErr

// for categorizing errors by their response code and no other response propery
func toErrByRespCode(ode oDataErr, err error) error {
	switch {
	case ode.hasResponseCode(http.StatusNotFound):
		return clues.Stack(core.ErrNotFound, err)
	case ode.hasResponseCode(http.StatusUnauthorized, http.StatusForbidden):
		return clues.Stack(core.ErrInsufficientAuthorization, err)
	case ode.hasResponseCode(http.StatusTooManyRequests):
		return clues.Stack(core.ErrApplicationThrottled, err)
	// catch any 5xx error
	case ode.hasResponseCode(5):
		return clues.Stack(core.ErrDownstreamServerError, err)
	default:
		return err
	}
}

func isErrApplicationThrottled(ode oDataErr, err error) bool {
	return ode.hasErrorCode(err, ApplicationThrottled)
}

func isErrInsufficientAuthorization(ode oDataErr, err error) bool {
	return ode.hasErrorCode(err, AuthorizationRequestDenied)
}

func isErrNotFound(ode oDataErr, err error) bool {
	return clues.HasLabel(err, LabelStatus(http.StatusNotFound)) ||
		ode.hasErrorCode(
			err,
			ErrorItemNotFound,
			ItemNotFound,
			syncFolderNotFound,
			notFound)
}

func isErrUserNotFound(ode oDataErr, err error) bool {
	if ode.hasErrorCode(err, RequestResourceNotFound, invalidUser) {
		return true
	}

	if ode.hasErrorCode(err, ResourceNotFound) {
		return strings.Contains(strings.ToLower(ode.Main.Message), "user")
	}

	return false
}

func isErrBadJWTToken(ode oDataErr, err error) bool {
	return ode.hasErrorCode(err, invalidAuthenticationToken)
}

func isErrItemAlreadyExists(ode oDataErr, err error) bool {
	return ode.hasErrorCode(err, nameAlreadyExists)
}

func isErrResourceLocked(ode oDataErr, err error) bool {
	return ode.hasErrorCode(err, ResourceLocked, NotAllowed) ||
		ode.errMessageMatchesAllFilters(
			err,
			filters.In([]string{"the service principal for resource"}),
			filters.In([]string{"this indicate that a subscription within the tenant has lapsed"}),
			filters.In([]string{"preventing tokens from being issued for it"}))
}

// exported categorizers

func IsErrAuthenticationError(err error) bool {
	return parseODataErr(err).hasErrorCode(err, AuthenticationError)
}

func IsErrInvalidDelta(err error) bool {
	return parseODataErr(err).hasErrorCode(err, syncStateNotFound, resyncRequired, syncStateInvalid)
}

func IsErrInvalidRequest(err error) bool {
	return parseODataErr(err).hasResponseCode(http.StatusBadRequest) &&
		parseODataErr(err).hasErrorCode(err, invalidRequest)
}

func IsErrDeltaNotSupported(err error) bool {
	return parseODataErr(err).hasErrorMessage(err, ParameterDeltaTokenNotSupported)
}

func IsErrQuotaExceeded(err error) bool {
	return parseODataErr(err).hasErrorCode(err, QuotaExceeded)
}

func IsErrExchangeMailFolderNotFound(err error) bool {
	// Not sure if we can actually see a resourceNotFound error here. I've only
	// seen the latter two.
	return parseODataErr(err).hasErrorCode(err, ResourceNotFound, ErrorItemNotFound, MailboxNotEnabledForRESTAPI)
}

func IsErrInvalidRecipients(err error) bool {
	return parseODataErr(err).hasErrorCode(err, ErrorInvalidRecipients)
}

func IsErrCorruptData(err error) bool {
	return parseODataErr(err).hasErrorCode(err, ErrorCorruptData)
}

func IsErrCannotOpenFileAttachment(err error) bool {
	return parseODataErr(err).hasErrorCode(err, cannotOpenFileAttachment)
}

func IsErrAccessDenied(err error) bool {
	return parseODataErr(err).hasErrorCode(err, ErrorAccessDenied) ||
		clues.HasLabel(err, LabelStatus(http.StatusForbidden))
}

func IsErrTimeout(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		return err.Timeout()
	}

	return errors.Is(err, context.Canceled) ||
		errors.Is(err, context.DeadlineExceeded) ||
		errors.Is(err, http.ErrHandlerTimeout) ||
		os.IsTimeout(err)
}

func IsErrConnectionReset(err error) bool {
	return errors.Is(err, syscall.ECONNRESET)
}

func IsErrUnauthorizedOrBadToken(err error) bool {
	return clues.HasLabel(err, LabelStatus(http.StatusUnauthorized)) ||
		parseODataErr(err).hasErrorCode(err, invalidAuthenticationToken) ||
		errors.Is(err, core.ErrAuthTokenExpired)
}

// LabelStatus transforms the provided statusCode into
// a standard label that can be attached to a clues error
// and later reviewed when checking error statuses.
func LabelStatus(statusCode int) string {
	return fmt.Sprintf("status_code_%d", statusCode)
}

// IsMalware is true if the graphAPI returns a "malware detected" error code.
func IsMalware(err error) bool {
	return parseODataErr(err).hasErrorCode(err, malwareDetected)
}

func IsMalwareResp(ctx context.Context, resp *http.Response) bool {
	// https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-wsshp/ba4ee7a8-704c-4e9c-ab14-fa44c574bdf4
	// https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-wdvmoduu/6fa6d4a9-ac18-4cd7-b696-8a3b14a98291
	return resp != nil &&
		len(resp.Header) > 0 &&
		resp.Header.Get("X-Virus-Infected") == "true"
}

func IsErrFolderExists(err error) bool {
	return parseODataErr(err).hasErrorCode(err, folderExists)
}

func IsErrUsersCannotBeResolved(err error) bool {
	ode := parseODataErr(err)

	return ode.hasErrorCode(err, noResolvedUsers) || ode.hasErrorMessage(err, usersCannotBeResolved)
}

func IsErrSiteNotFound(err error) bool {
	return parseODataErr(err).hasErrorMessage(err, requestedSiteCouldNotBeFound)
}

func IsErrSharingDisabled(err error) bool {
	return parseODataErr(err).hasErrorCode(err, sharingDisabled)
}

// ---------------------------------------------------------------------------
// quality of life wrappers
// ---------------------------------------------------------------------------

// Wrap is a helper function that extracts ODataError metadata from
// the error.  If the error is not an ODataError type, returns the error.
// You probably don't need this, because all calls to the graph client
// automatically Stack these details.  You probably want clues.Wrap.
// This is primarily exported for test helpers.
func Wrap(ctx context.Context, e error, msg string) *clues.Err {
	if e == nil {
		return nil
	}

	ode := parseODataErr(e)
	if !ode.isODataErr {
		return clues.StackWC(ctx, e).WithTrace(1)
	}

	if len(ode.Main.Message) > 0 {
		e = clues.Stack(e, clues.New(ode.Main.Message))
	}

	ce := clues.WrapWC(ctx, e, msg).
		With("graph_api_err", ode).
		WithTrace(1)

	return setLabels(ce, ode)
}

// Stack is a helper function that extracts ODataError metadata from
// the error.  If the error is not an ODataError type, returns the error.
// You probably don't need this, because all calls to the graph client
// automatically Stack these details.  You probably want clues.Stack.
// This is primarily exported for test helpers.
func Stack(ctx context.Context, e error) *clues.Err {
	return stackWithDepth(ctx, e, 1)
}

func stackWithDepth(ctx context.Context, e error, traceDepth int) *clues.Err {
	if e == nil {
		return nil
	}

	ode := parseODataErr(e)
	if !ode.isODataErr {
		return clues.StackWC(ctx, e).WithTrace(1)
	}

	if len(ode.Main.Message) > 0 {
		e = clues.Stack(e, clues.New(ode.Main.Message))
	}

	ce := clues.StackWC(ctx, e).
		With("graph_api_err", ode).
		WithTrace(1 + traceDepth)

	return setLabels(ce, ode)
}

// stackReq is a helper function that extracts ODataError metadata from
// the error, plus http req/resp data.  If the error is not an ODataError
// type, returns the error with only the req/resp values.
func stackReq(
	ctx context.Context,
	req *http.Request,
	resp *http.Response,
	e error,
) *clues.Err {
	// For service unavailable stack an error we can check against. We can't check
	// this at higher layers because we won't have the content length or server
	// response code thanks to graph SDK. We do need to check this before checking
	// if there was no error though because the status code itself doesn't count
	// as an error.
	//
	// Returning an error here is non-standard, but if we don't return something
	// then all we get from graph SDK is an error saying "content is empty" which
	// isn't particularly useful.
	if resp != nil &&
		resp.ContentLength == 0 {
		switch resp.StatusCode {
		case http.StatusServiceUnavailable:
			e = clues.Stack(ErrServiceUnavailableEmptyResp, e)

		case http.StatusNotFound:
			e = clues.Stack(ErrNotFoundEmptyResp, e)
		}
	}

	if e == nil {
		return nil
	}

	se := Stack(ctx, e).
		WithMap(reqData(req)).
		WithMap(respData(resp))

	return se
}

// Checks for the following conditions and labels the error accordingly:
// * mysiteNotFound | mysiteURLNotFound
// * malware
func setLabels(err *clues.Err, ode oDataErr) *clues.Err {
	if err == nil {
		return nil
	}

	f := filters.Contains([]string{ode.Main.Message + ode.Main.Code})

	if f.Compare(string(MysiteNotFound)) ||
		f.Compare(string(MysiteURLNotFound)) {
		err = err.Label(LabelsMysiteNotFound)
	}

	if f.Compare(string(NoSPLicense)) {
		err = err.Label(LabelsNoSharePointLicense)
	}

	if IsMalware(err) {
		err = err.Label(LabelsMalware)
	}

	return err
}

func reqData(req *http.Request) map[string]any {
	if req == nil {
		return nil
	}

	r := map[string]any{}
	r["req_method"] = req.Method
	r["req_len"] = req.ContentLength

	if req.URL != nil {
		r["req_url"] = LoggableURL(req.URL.String())
	}

	return r
}

func respData(resp *http.Response) map[string]any {
	if resp == nil {
		return nil
	}

	r := map[string]any{}
	r["resp_status"] = resp.Status
	r["resp_len"] = resp.ContentLength

	return r
}

// ItemInfo gathers potentially useful information about a drive item,
// and aggregates that data into a map.
func ItemInfo(item *custom.DriveItem) map[string]any {
	m := map[string]any{}

	creator := item.GetCreatedByUser()
	if creator != nil {
		m[fault.AddtlCreatedBy] = ptr.Val(creator.GetId())
	}

	lastmodder := item.GetLastModifiedByUser()
	if lastmodder != nil {
		m[fault.AddtlLastModBy] = ptr.Val(lastmodder.GetId())
	}

	parent := item.GetParentReference()
	if parent != nil {
		m[fault.AddtlContainerID] = ptr.Val(parent.GetId())
		m[fault.AddtlContainerName] = ptr.Val(parent.GetName())
		containerPath := ""

		// Remove the "/drives/b!vF-sdsdsds-sdsdsa-sdsd/root:" prefix
		splitPath := strings.SplitN(ptr.Val(parent.GetPath()), ":", 2)
		if len(splitPath) > 1 {
			containerPath = splitPath[1]
		}

		m[fault.AddtlContainerPath] = containerPath
	}

	malware := item.GetMalware()
	if malware != nil {
		m[fault.AddtlMalwareDesc] = ptr.Val(malware.GetDescription())
	}

	return m
}

// ---------------------------------------------------------------------------
// error parsers
// ---------------------------------------------------------------------------

type oDataErr struct {
	isODataErr bool
	Details    []oDataDeets `json:"details,omitempty"`
	Inner      innerErr     `json:"inner,omitempty"`
	Main       mainErr      `json:"main,omitempty"`
	Resp       apiResp      `json:"resp,omitempty"`
}

type oDataDeets struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Target  string `json:"target,omitempty"`
}

type innerErr struct {
	ClientRequestID string    `json:"clientRequestID,omitempty"`
	Code            string    `json:"code,omitempty"`
	Date            time.Time `json:"date,omitempty"`
	OdataType       string    `json:"odataType,omitempty"`
	RequestID       string    `json:"requestID,omitempty"`
}

type mainErr struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Target  string `json:"target,omitempty"`
}

type apiResp struct {
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
}

// parses the odataerror.ODataError interface to a local
// struct for easy logging/usage.
func parseODataErr(err error) oDataErr {
	if err == nil {
		return oDataErr{}
	}

	var (
		ode      oDataErr
		graphODE *odataerrors.ODataError
	)

	if !errors.As(err, &graphODE) {
		return oDataErr{}
	}

	ode = oDataErr{
		isODataErr: true,
		Resp: apiResp{
			Message:    strings.Clone(graphODE.ApiError.Message),
			StatusCode: graphODE.ApiError.ResponseStatusCode,
		},
	}

	var (
		main    = graphODE.GetErrorEscaped()
		details []odataerrors.ErrorDetailsable
		inner   odataerrors.InnerErrorable
	)

	if main != nil {
		ode.Main = mainErr{
			Code:    strings.Clone(ptr.Val(main.GetCode())),
			Message: strings.Clone(ptr.Val(main.GetMessage())),
			Target:  strings.Clone(ptr.Val(main.GetTarget())),
		}

		inner = main.GetInnerError()
		details = main.GetDetails()
	}

	if inner != nil {
		ode.Inner = innerErr{
			ClientRequestID: strings.Clone(ptr.Val(inner.GetClientRequestId())),
			Date:            ptr.Val(inner.GetDate()),
			OdataType:       strings.Clone(ptr.Val(inner.GetOdataType())),
			RequestID:       strings.Clone(ptr.Val(inner.GetRequestId())),
		}

		code, err := str.AnyValueToString("code", inner.GetAdditionalData())
		if err == nil {
			ode.Inner.Code = code
		}
	}

	ode.Details = make([]oDataDeets, 0, len(details))

	for _, deet := range details {
		d := oDataDeets{
			Code:    strings.Clone(ptr.Val(deet.GetCode())),
			Message: strings.Clone(ptr.Val(deet.GetMessage())),
			Target:  strings.Clone(ptr.Val(deet.GetTarget())),
		}

		ode.Details = append(ode.Details, d)
	}

	return ode
}

// hasResponseCode checks if the error is an ODataError and carries the provided
// http response code.
func (ode oDataErr) hasResponseCode(httpResponseCode ...int) bool {
	if !ode.isODataErr {
		return false
	}

	orsc := ode.Resp.StatusCode

	for _, hrc := range httpResponseCode {
		if orsc == hrc {
			return true
		}

		if hrc < 10 && orsc/100 == hrc {
			return true
		}
	}

	return false
}

func (ode oDataErr) hasErrorCode(err error, codes ...errorCode) bool {
	if !ode.isODataErr {
		return false
	}

	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = string(c)
	}

	eq := filters.Equal(cs)

	return eq.CompareAny(ode.Main.Code, ode.Inner.Code)
}

// only use this as a last resort.  Prefer the code or statuscode if possible.
func (ode oDataErr) hasErrorMessage(err error, msgs ...errorMessage) bool {
	if !ode.isODataErr {
		return false
	}

	cs := make([]string, len(msgs))
	for i, c := range msgs {
		cs[i] = string(c)
	}

	return filters.In(cs).Compare(ode.Main.Message)
}

// only use this as a last resort.  Prefer the code or statuscode if possible.
func (ode oDataErr) errMessageMatchesAllFilters(err error, fs ...filters.Filter) bool {
	if !ode.isODataErr {
		return false
	}

	return filters.Must(ode.Main.Message, fs...)
}

// ---------------------------------------------------------------------------
// other helpers
// ---------------------------------------------------------------------------
const (
	// JWTQueryParam is a query param embed in graph download URLs which holds
	// JWT token.
	JWTQueryParam = "tempauth"
	// base64 encoded json header. Contains {"alg":"HS256","typ":"JWT"}
	//
	// Hardcoding this instead of generating it every time on the fly.
	// The algorithm doesn't matter as we are not verifying the token.
	jwtHeader = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
)

func sanitizeToken(rawToken string) string {
	segments := strings.Split(rawToken, ".")

	// Check if the token has the old format, in which it has 3 segments and
	// conforms to jwt spec. Format is seg1.seg2.seg3.
	if len(segments) == 3 {
		return rawToken
	}

	// Check if it is a msft proprietary token in which it has 4 segments and
	// doesn't meet jwt spec. Format is v1.seg1.seg2.seg3. Return a token which
	// meets jwt spec.
	//
	// In this proprietary token, there is no jwt header segment. Also, the claims
	// section is split into first and segments. The first segment contains the
	// `exp` claim that we are interested in.
	//
	// The second segment contains the rest of the claims, but likely encrypted.
	// We don't need it so discard it. The last segment contains the signature which
	// we don't care about either, as we are not verifying the token. So append it as is.
	//
	// It's okay if the sanitized token still doesn't meet jwt spec. It'll fail decoding
	// later and we have fallbacks for that.
	if len(segments) == 4 && segments[0] == "v1" {
		return jwtHeader + "." + segments[1] + "." + segments[3]
	}

	// If MSFT change the token format again on us, just return empty string and let caller
	// handle it as an error.
	return ""
}

// IsURLExpired inspects the jwt token embed in the item download url
// and returns true if it is expired.
func IsURLExpired(
	ctx context.Context,
	urlStr string,
) (
	expiredErr error,
	err error,
) {
	ctx = clues.Add(ctx, "checked_url", urlStr)

	// Extract the raw JWT string from the download url.
	rawJWT, err := common.GetQueryParamFromURL(urlStr, JWTQueryParam)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "jwt query param not found")
	}

	// Token may have a proprietary format. Try to sanitize it to jwt format.
	rawJWT = sanitizeToken(rawJWT)
	if len(rawJWT) == 0 {
		return nil, clues.WrapWC(ctx, err, "sanitizing jwt")
	}

	expired, err := jwt.IsJWTExpired(rawJWT)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "checking jwt expiry")
	}

	if expired {
		return clues.StackWC(ctx, core.ErrAuthTokenExpired), nil
	}

	return nil, nil
}
