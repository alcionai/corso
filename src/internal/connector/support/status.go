package support

import (
	"errors"
	"fmt"
)

type ConnectorOperationStatus struct {
	lastOperation    operation
	objectCount      int
	folderCount      int
	successful       int
	errorCount       int
	incomplete       bool
	incompleteReason string
}

type operation int

const (
	backup operation = iota
	restore
)

// Constructor for ConnectorOperationStatus. If the counts do not agree, an error is returned.
func CreateStatus(operationType, objects, success, folders, errCount int, errStatus string) (*ConnectorOperationStatus, error) {
	var op operation
	hasErrors := errCount > 0
	if operationType == 0 {
		op = backup
	} else {
		op = restore
	}
	status := ConnectorOperationStatus{
		lastOperation:    op,
		objectCount:      objects,
		folderCount:      folders,
		successful:       success,
		errorCount:       errCount,
		incomplete:       hasErrors,
		incompleteReason: errStatus,
	}
	if status.objectCount != status.errorCount+status.successful {
		return nil, errors.New("incorrect total on initialization")
	}
	return &status, nil
}

//  GetOperation helper function to standardize the instantiation of LastOperation
func GetOperation(selection operation) string {

	switch selection {
	case backup:
		return "Backup"
	default:
		return "Restore"
	}
}

func (cos *ConnectorOperationStatus) ToString() string {
	message := fmt.Sprintf("Action: %s performed on %d of %d for %d directories.", GetOperation(cos.lastOperation),
		cos.successful, cos.objectCount, cos.folderCount)
	if cos.incomplete {
		message = message + fmt.Sprintf(" %s", cos.incompleteReason)
	}
	message = message + "\n"
	return message
}
