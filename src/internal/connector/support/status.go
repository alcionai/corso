package support

import (
	"context"
	"fmt"

	"github.com/alcionai/corso/pkg/logger"
)

type ConnectorOperationStatus struct {
	lastOperation    Operation
	ObjectCount      int
	FolderCount      int
	Successful       int
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
func CreateStatus(
	ctx context.Context,
	op Operation,
	objects, success, folders int,
	err error,
) *ConnectorOperationStatus {
	var reason string

	if err != nil {
		reason = err.Error()
	}

	hasErrors := err != nil
	numErr := GetNumberOfErrors(err)
	status := ConnectorOperationStatus{
		lastOperation:    op,
		ObjectCount:      objects,
		FolderCount:      folders,
		Successful:       success,
		errorCount:       numErr,
		incomplete:       hasErrors,
		incompleteReason: reason,
	}

	if status.ObjectCount != status.errorCount+status.Successful {
		logger.Ctx(ctx).DPanicw(
			"status object count does not match errors + successes",
			"objects", objects,
			"successes", success,
			"errors", numErr)
	}

	return &status
}

// Function signature for a status updater
type StatusUpdater func(*ConnectorOperationStatus)

// MergeStatus combines ConnectorOperationsStatus value into a single status
func MergeStatus(one, two ConnectorOperationStatus) ConnectorOperationStatus {
	var hasErrors bool
	if one.lastOperation == OpUnknown {
		return two
	}

	if two.lastOperation == OpUnknown {
		return one
	}

	if one.incomplete || two.incomplete {
		hasErrors = true
	}

	status := ConnectorOperationStatus{
		lastOperation:    one.lastOperation,
		ObjectCount:      one.ObjectCount + two.ObjectCount,
		FolderCount:      one.FolderCount + two.FolderCount,
		Successful:       one.Successful + two.Successful,
		errorCount:       one.errorCount + two.errorCount,
		incomplete:       hasErrors,
		incompleteReason: one.incompleteReason + " " + two.incompleteReason,
	}

	return status
}

func (cos *ConnectorOperationStatus) String() string {
	message := fmt.Sprintf("Action: %s performed on %d of %d objects within %d directories.", cos.lastOperation.String(),
		cos.Successful, cos.ObjectCount, cos.FolderCount)
	if cos.incomplete {
		message += " " + cos.incompleteReason
	}

	message = message + "\n"

	return message
}
