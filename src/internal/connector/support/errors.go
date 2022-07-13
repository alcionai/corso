package support

import (
	"fmt"
	"strconv"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	msgraph_errors "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/common"
)

// GraphConnector has two types of errors that are exported
// RecoverableGCError is a query error that can be overcome with time
type RecoverableGCError struct {
	common.BaseError
}

func SetRecoverableError(e error) error {
	return RecoverableGCError{common.BaseError{Err: e}}
}

// NonRecoverableGCError is a permanent query error
type NonRecoverableGCError struct {
	common.BaseError
}

func SetNonRecoverableError(e error) error {
	return NonRecoverableGCError{common.BaseError{Err: e}}
}

// WrapErrorAndAppend helper function used to attach identifying information to an error
// and return it as a mulitierror
func WrapAndAppend(identifier string, e, previous error) error {
	return multierror.Append(previous, errors.Wrap(e, identifier))
}

// WrapErrorAndAppendf format version of WrapErrorAndAppend
func WrapAndAppendf(identifier interface{}, e, previous error) error {
	return multierror.Append(previous, errors.Wrapf(e, "%v", identifier))
}

// GetErrors Helper method to return the integer amount of errors in multi error
func GetNumberOfErrors(err error) int {
	if err == nil {
		return 0
	}
	result, _, wasFound := strings.Cut(err.Error(), " ")
	if wasFound {
		aNum, err := strconv.Atoi(result)
		if err == nil {
			return aNum
		}
	}
	return 1
}

// ListErrors is a helper method used to return the string of errors when
// the multiError library is used.
// depends on ConnectorStackErrorTrace
func ListErrors(multi multierror.Error) string {
	aString := ""
	for idx, err := range multi.Errors {
		detail := ConnectorStackErrorTrace(err)
		if detail == "" {
			detail = fmt.Sprintf("%v", err)
		}
		aString = aString + fmt.Sprintf("\n\tErr: %d %v", idx+1, detail)
	}
	return aString
}

// concatenateStringFromPointers is a helper funtion that adds
// strings to the originalMessage iff the pointer is not nil
func concatenateStringFromPointers(orig string, pointers []*string) string {
	for _, pointer := range pointers {
		if pointer != nil {
			orig = strings.Join([]string{orig, *pointer}, " ")
		}
	}
	return orig
}

// ConnectorStackErrorTrace is a helper function that wraps the
// stack trace for oDataError types from querying the M365 back store.
func ConnectorStackErrorTrace(e error) string {
	eMessage := ""
	if oDataError, ok := e.(msgraph_errors.ODataErrorable); ok {
		// Get MainError
		mainErr := oDataError.GetError()
		// message *string
		// target *string
		// code *string
		// details ErrorDetailsable
		// Ignoring Additonal Detail
		code := mainErr.GetCode()
		subject := mainErr.GetMessage()
		target := mainErr.GetTarget()
		details := mainErr.GetDetails()
		inners := mainErr.GetInnererror()
		eMessage = concatenateStringFromPointers(eMessage,
			[]*string{code, subject, target})
		// Get Error Details
		// code, message, target
		if details != nil {
			eMessage = eMessage + "\nDetails Section:"
			for idx, detail := range details {
				dMessage := fmt.Sprintf("Detail %d:", idx)
				c := detail.GetCode()
				m := detail.GetMessage()
				t := detail.GetTarget()
				dMessage = concatenateStringFromPointers(dMessage,
					[]*string{c, m, t})
				eMessage = eMessage + dMessage
			}
		}
		if inners != nil {
			eMessage = eMessage + "\nConnector Section:"
			client := inners.GetClientRequestId()
			rId := inners.GetRequestId()
			eMessage = concatenateStringFromPointers(eMessage,
				[]*string{client, rId})
		}
	}
	return eMessage
}
