package support

import (
	"errors"
	"fmt"
)

type ConnectorOperationStatus struct {
	lastOperation    Operation
	objectCount      int
	folderCount      int
	successful       int
	errorCount       int
	incomplete       bool
	incompleteReason string
}

type Operation int

//go:generate stringer -type=Operation
const (
	Backup Operation = iota
	Restore
)

// Constructor for ConnectorOperationStatus. If the counts do not agree, an error is returned.
func CreateStatus(op Operation, objects, success, folders int, err error) (*ConnectorOperationStatus, error) {
	hasErrors := err != nil
	var reason string
	if err != nil {
		reason = err.Error()
	}
	status := ConnectorOperationStatus{
		lastOperation:    op,
		objectCount:      objects,
		folderCount:      folders,
		successful:       success,
		errorCount:       GetNumberOfErrors(err),
		incomplete:       hasErrors,
		incompleteReason: reason,
	}
	if status.objectCount != status.errorCount+status.successful {
		return nil, errors.New("incorrect total on initialization")
	}
	return &status, nil
}

func (cos *ConnectorOperationStatus) String() string {
	message := fmt.Sprintf("Action: %s performed on %d of %d objects within %d directories.", cos.lastOperation.String(),
		cos.successful, cos.objectCount, cos.folderCount)
	if cos.incomplete {
		message = message + fmt.Sprintf(" %s", cos.incompleteReason)
	}
	message = message + "\n"
	return message
}
