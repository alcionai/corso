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

	hasErrors := err != nil
	if hasErrors {
		reason = err.Error()
	}

	numErr := GetNumberOfErrors(err)
	status := ConnectorOperationStatus{
		lastOperation:    op,
		ObjectCount:      objects,
		folderCount:      folders,
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

func (cos *ConnectorOperationStatus) String() string {
	message := fmt.Sprintf("Action: %s performed on %d of %d objects within %d directories.", cos.lastOperation.String(),
		cos.Successful, cos.ObjectCount, cos.folderCount)
	if cos.incomplete {
		message += " " + cos.incompleteReason
	}

	message = message + "\n"

	return message
}
