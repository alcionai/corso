package datautil

import (
	"fmt"

	msgraph_errors "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

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

func (el ErrorList) GetErrors() string {
	aString := ""
	if len(el.errorList) != 0 {
		for idx, err := range el.errorList {
			aString = aString + fmt.Sprintf("\n\tErr %d %v", idx, err)
		}
	}
	return aString
}

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
