package support

import (
	"context"
	"fmt"

	bytesize "github.com/inhies/go-bytesize"

	"github.com/alcionai/corso/src/pkg/logger"
)

type ConnectorOperationStatus struct {
	lastOperation     Operation
	ObjectCount       int
	FolderCount       int
	Successful        int
	errorCount        int
	incomplete        bool
	incompleteReason  string
	additionalDetails string
	bytes             int
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
	objects, success, folders, bytes int,
	err error,
	details string,
) *ConnectorOperationStatus {
	var reason string

	if err != nil {
		reason = err.Error()
	}

	hasErrors := err != nil
	numErr := GetNumberOfErrors(err)
	status := ConnectorOperationStatus{
		lastOperation:     op,
		ObjectCount:       objects,
		FolderCount:       folders,
		Successful:        success,
		errorCount:        numErr,
		incomplete:        hasErrors,
		incompleteReason:  reason,
		bytes:             bytes,
		additionalDetails: details,
	}

	if status.ObjectCount != status.errorCount+status.Successful {
		logger.Ctx(ctx).DPanicw(
			"status object count does not match errors + successes",
			"objects", objects,
			"successes", success,
			"numErrors", numErr,
			"errors", err.Error())
	}

	return &status
}

// Function signature for a status updater
// Used to define a function that an async connector task can call
// to on completion with its ConnectorOperationStatus
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
		lastOperation:     one.lastOperation,
		ObjectCount:       one.ObjectCount + two.ObjectCount,
		FolderCount:       one.FolderCount + two.FolderCount,
		Successful:        one.Successful + two.Successful,
		errorCount:        one.errorCount + two.errorCount,
		bytes:             one.bytes + two.bytes,
		incomplete:        hasErrors,
		incompleteReason:  one.incompleteReason + " " + two.incompleteReason,
		additionalDetails: one.additionalDetails + " " + two.additionalDetails,
	}

	return status
}

func (cos *ConnectorOperationStatus) String() string {
	bs := bytesize.New(float64(cos.bytes))
	message := fmt.Sprintf("Action: %s performed on %d of %d objects within %d directories. Downloaded: %s",
		cos.lastOperation.String(),
		cos.Successful,
		cos.ObjectCount,
		cos.FolderCount,
		bs,
	)

	if cos.incomplete {
		message += " " + cos.incompleteReason
	}

	message += " " + cos.additionalDetails + "\n"

	return message
}
