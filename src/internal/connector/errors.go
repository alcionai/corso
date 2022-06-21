package connector

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	msgraph_errors "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/pkg/logger"
)

// WrapErrorAndAppend helper function used to attach identifying information to an error
// and return it as a mulitierror
func WrapAndAppend(identifier string, e error, previous error) error {
	return multierror.Append(previous, errors.Wrap(e, identifier))
}

// WrapErrorAndAppendf format version of WrapErrorAndAppend
func WrapAndAppendf(identifier interface{}, e error, previous error) error {
	return multierror.Append(previous, errors.Wrapf(e, "%v", identifier))
}

// GetErrors Helper method to return the integer amount of errors in multi error
func GetNumberOfErrors(err error) int {
	words := strings.Split(err.Error(), " ")
	aNum, err := strconv.Atoi(words[0])
	if err != nil {
		return 1 // Base case
	}
	return aNum
}

// ListErrors is a helper method used to return the string of errors when
// the multiError library is used.
// depends on ConnectorStackErrorTrace
func ListErrors(ctx context.Context, multi multierror.Error) string {
	aString := ""
	for idx, err := range multi.Errors {
		detail := ConnectorStackErrorTrace(ctx, err)
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
func ConnectorStackErrorTrace(ctx context.Context, e error) string {
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
			logger.Ctx(ctx).Debug("error contains inner errors")
			eMessage = eMessage + "\nConnector Section:"
			client := inners.GetClientRequestId()
			rId := inners.GetRequestId()
			eMessage = concatenateStringFromPointers(eMessage,
				[]*string{client, rId})
		}
	}
	return eMessage
}
