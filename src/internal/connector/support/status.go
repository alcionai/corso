package support

import (
	"context"
	"fmt"

	"github.com/dustin/go-humanize"
	multierror "github.com/hashicorp/go-multierror"

	"github.com/alcionai/corso/src/pkg/logger"
)

// ConnectorOperationStatus is a data type used to describe the state of
// the sequence of operations.
// @param ObjectCount integer representation of how many objects have downloaded or uploaded.
// @param Successful: Number of objects that are sent through the connector without incident.
// @param incomplete: Bool representation of whether all intended items were download or uploaded.
// @param bytes: represents the total number of bytes that have been downloaded or uploaded.
type ConnectorOperationStatus struct {
	lastOperation     Operation
	ObjectCount       int
	FolderCount       int
	Successful        int
	ErrorCount        int
	Err               error
	incomplete        bool
	incompleteReason  string
	additionalDetails string
	bytes             int64
}

type CollectionMetrics struct {
	Objects, Successes int
	TotalBytes         int64
}

func (cm *CollectionMetrics) Combine(additional CollectionMetrics) {
	cm.Objects += additional.Objects
	cm.Successes += additional.Successes
	cm.TotalBytes += additional.TotalBytes
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
	folders int,
	cm CollectionMetrics,
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
		ObjectCount:       cm.Objects,
		FolderCount:       folders,
		Successful:        cm.Successes,
		ErrorCount:        numErr,
		Err:               err,
		incomplete:        hasErrors,
		incompleteReason:  reason,
		bytes:             cm.TotalBytes,
		additionalDetails: details,
	}

	if status.ObjectCount != status.ErrorCount+status.Successful {
		logger.Ctx(ctx).Errorw(
			"status object count does not match errors + successes",
			"objects", cm.Objects,
			"successes", cm.Successes,
			"numErrors", numErr,
			"errors", err)
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
		ErrorCount:        one.ErrorCount + two.ErrorCount,
		Err:               multierror.Append(one.Err, two.Err).ErrorOrNil(),
		bytes:             one.bytes + two.bytes,
		incomplete:        hasErrors,
		incompleteReason:  one.incompleteReason + ", " + two.incompleteReason,
		additionalDetails: one.additionalDetails + ", " + two.additionalDetails,
	}

	return status
}

func (cos *ConnectorOperationStatus) String() string {
	var operationStatement string

	switch cos.lastOperation {
	case Backup:
		operationStatement = "Downloaded from "
	case Restore:
		operationStatement = "Restored content to "
	}

	message := fmt.Sprintf("Action: %s performed on %d of %d objects (%s) within %d directories.",
		cos.lastOperation.String(),
		cos.Successful,
		cos.ObjectCount,
		humanize.Bytes(uint64(cos.bytes)),
		cos.FolderCount,
	)

	if cos.incomplete {
		message += " " + cos.incompleteReason
	}

	message += " " + operationStatement + cos.additionalDetails + "\n"

	return message
}
