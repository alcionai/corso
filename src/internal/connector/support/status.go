package support

import (
	"context"
	"fmt"

	"github.com/alcionai/corso/pkg/logger"
)

type ConnectorOperationStatus struct {
	lastOperation    Operation
	ObjectCount      int
	folderCount      int
	successful       int
	errorCount       int
	incomplete       bool
	incompleteReason string
}

type Operation int

//go:generate stringer -type=Operation
const (
	OpUnknown Operation = iota
	Backup
	Restore
)

// Constructor for ConnectorOperationStatus. If the counts do not agree, an error is returned.
func CreateStatus(ctx context.Context, op Operation, objects, success, folders int, err error) *ConnectorOperationStatus {
	hasErrors := err != nil
	var reason string
	if err != nil {
		reason = err.Error()
	}
	numErr := GetNumberOfErrors(err)
	status := ConnectorOperationStatus{
		lastOperation:    op,
		ObjectCount:      objects,
		folderCount:      folders,
		successful:       success,
		errorCount:       numErr,
		incomplete:       hasErrors,
		incompleteReason: reason,
	}
	if status.ObjectCount != status.errorCount+status.successful {
		logger.Ctx(ctx).DPanicw(
			"status object count does not match errors + successes",
			"objects", objects,
			"successes", success,
			"errors", numErr)
	}
	return &status
}

// MergeStatus combines ConnectorOperationsStatus value into a single status
func MergeStatus(one, two *ConnectorOperationStatus) *ConnectorOperationStatus {
	if one == nil && two == nil {
		return nil
	}
	if one != nil && two == nil {
		return one

	}
	if one == nil && two != nil {
		return two
	}

	var hasErrors bool
	if one.incomplete || two.incomplete {
		hasErrors = true
	}

	status := ConnectorOperationStatus{
		lastOperation:    one.lastOperation,
		ObjectCount:      one.ObjectCount + two.ObjectCount,
		folderCount:      one.folderCount + two.folderCount,
		successful:       one.successful + two.successful,
		errorCount:       one.errorCount + two.errorCount,
		incomplete:       hasErrors,
		incompleteReason: one.incompleteReason + " " + two.incompleteReason,
	}
	return &status
}

func (cos *ConnectorOperationStatus) String() string {
	message := fmt.Sprintf("Action: %s performed on %d of %d objects within %d directories.", cos.lastOperation.String(),
		cos.successful, cos.ObjectCount, cos.folderCount)
	if cos.incomplete {
		message += " " + cos.incompleteReason
	}
	message = message + "\n"
	return message
}
