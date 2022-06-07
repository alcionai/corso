package connector

import (
	"fmt"

	"github.com/pkg/errors"

	msgraph_errors "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ErrorList is a typical list of errors. The additional functionality comes
// from helper functions that are specific to Microsoft Graph API
type ErrorList []error

func NewErrorList() ErrorList {
	errors := make([]error, 0)
	return errors
}

// Returns generic list of error strings
func (el ErrorList) GetErrors() string {
	aString := ""
	if len(el) != 0 {
		for idx, err := range el {
			aString = aString + fmt.Sprintf("\n\tErr: %d %v", idx+1, err)
		}
	}
	return aString
}

// WrapErrorAndAppend helper function used to attach identifying information to the error
// and append to the list
func WrapErrorAndAppend(eventId string, e error, eList ErrorList) ErrorList {
	newError := errors.Wrap(e, eventId)
	eList = append(eList, newError)
	return eList
}

// GetDetailedErrors returns a message, a compilation of errors stored within the list.
// Extends "Post Errors" by default using the GetPostErrorString method
func (el *ErrorList) GetDetailedErrors() string {
	aString := ""
	for idx, err := range *el {
		detail := GetPostErrorsString(err)
		if detail == "" {
			detail = fmt.Sprintf("%v", err)
		}
		aString = aString + fmt.Sprintf("\n\tErr: %d %v", idx+1, detail)
	}
	return aString
}

// GetPostErrorString returns string identifying the chain of errors that
// occurred while retrieving data from the M365 back store.
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
