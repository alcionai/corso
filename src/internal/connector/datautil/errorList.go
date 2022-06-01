// datautil package provides a series of helper functions to be better
// interpret, respond, or manipulate data that interacts with the M365 objects.
package datautil

import (
	"fmt"

	msgraph_errors "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ErrorList a struct for holding errors experienced in sequence or in a
// scenario where the errors are later required for deeper analysis.
type ErrorList struct {
	errorList []error
}

func NewErrorList() ErrorList {
	el := ErrorList{
		errorList: make([]error, 0),
	}
	return el
}

func (el *ErrorList) AddError(err *error) {
	newErr := *err
	el.errorList = append(el.errorList, newErr)
}

func (el *ErrorList) GetLength() int {
	return len(el.errorList)
}

// GetErrors returns the string value of all errors experienced.
func (el ErrorList) GetErrors() string {
	aString := ""
	if len(el.errorList) != 0 {
		for idx, err := range el.errorList {
			aString = aString + fmt.Sprintf("\n\tErr %d %v", idx, err)
		}
	}
	return aString
}

// GetDetailedErrors is a helper method for returning a the string
// representation of failures  out of the layers of msgraph-sdk-go using
// GetPostErrorsString.
func (el ErrorList) GetDetailedErrors() string {
	aString := ""
	for idx, err := range el.errorList {
		detail := GetPostErrorsString(err)
		if detail == "" {
			detail = fmt.Sprintf("%v", err)
		}
		aString = aString + fmt.Sprintf("\n\tErr %d %v", idx, detail)
	}
	return aString
}

// GetPostErrorsString is a helper method for burrowing into the oDataError
// created by M365. These errors  may have several cascading failures
// associated with them. The method returns a string.
func GetPostErrorsString(e error) string {
	eMessage := ""
	if oDataError, ok := e.(msgraph_errors.ODataErrorable); ok {
		// Get MainError
		mainErr := oDataError.GetError()
		// message *string
		// target *string
		// code *string
		// details ErrorDetailsable
		// Ignoring Additonal Detail
		space := " "
		code := mainErr.GetCode()
		subject := mainErr.GetMessage()
		target := mainErr.GetTarget()
		details := mainErr.GetDetails()
		inners := mainErr.GetInnererror()
		if code != nil {
			eMessage = eMessage + *code + space
		}
		if subject != nil {
			eMessage = eMessage + *subject + space
		}
		if target != nil {
			eMessage = eMessage + *target
		}
		// Get Error Details
		// code, message, target
		if details != nil {
			eMessage = eMessage + "\nDetails Section:"
			for idx, detail := range details {
				dMessage := fmt.Sprintf("Detail %d:", idx)
				c := detail.GetCode()
				m := detail.GetMessage()
				t := detail.GetTarget()
				if c != nil {
					dMessage = dMessage + space + *c + space
				}
				if m != nil {
					dMessage = dMessage + *m + space
				}
				if t != nil {
					dMessage = dMessage + *t + space
				}
				eMessage = eMessage + dMessage
			}
		}
		if inners != nil {
			fmt.Println("Inners not nil")
			eMessage = eMessage + "\nConnector Section:"
			client := inners.GetClientRequestId()
			rId := inners.GetRequestId()
			if client != nil {
				eMessage = eMessage + space + *client + space
			}
			if rId != nil {
				eMessage = eMessage + *rId + space
			}
		}
	}
	return eMessage
}

func PrintPostErrors(e error) {
	errorString := GetPostErrorsString(e)
	fmt.Println(errorString)
}
