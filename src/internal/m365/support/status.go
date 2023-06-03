package support

import (
	"context"
	"fmt"

	"github.com/dustin/go-humanize"
)

// ControllerOperationStatus is a data type used to describe the state of
// the sequence of operations.
type ControllerOperationStatus struct {
	Folders int
	Metrics CollectionMetrics
	details string
	op      Operation
}

type CollectionMetrics struct {
	Objects, Successes int
	Bytes              int64
}

func CombineMetrics(a, b CollectionMetrics) CollectionMetrics {
	return CollectionMetrics{
		Objects:   a.Objects + b.Objects,
		Successes: a.Successes + b.Successes,
		Bytes:     a.Bytes + b.Bytes,
	}
}

type Operation string

const (
	Backup  = "backup"
	Restore = "restore"
)

func CreateStatus(
	ctx context.Context,
	op Operation,
	folders int,
	cm CollectionMetrics,
	details string,
) *ControllerOperationStatus {
	status := ControllerOperationStatus{
		Folders: folders,
		Metrics: cm,
		details: details,
		op:      op,
	}

	return &status
}

// Function signature for a status updater
// Used to define a function that an async controller task can call
// to on completion with its ControllerOperationStatus
type StatusUpdater func(*ControllerOperationStatus)

// MergeStatus combines ControllerOperationsStatus value into a single status
func MergeStatus(one, two ControllerOperationStatus) ControllerOperationStatus {
	if len(one.op) == 0 {
		return two
	}

	if len(two.op) == 0 {
		return one
	}

	status := ControllerOperationStatus{
		Folders: one.Folders + two.Folders,
		Metrics: CombineMetrics(one.Metrics, two.Metrics),
		details: one.details + ", " + two.details,
		op:      one.op,
	}

	return status
}

func (cos *ControllerOperationStatus) String() string {
	var operationStatement string

	switch cos.op {
	case Backup:
		operationStatement = "Downloaded from "
	case Restore:
		operationStatement = "Restored content to "
	}

	return fmt.Sprintf("Action: %s performed on %d of %d objects (%s) within %d directories.  %s %s",
		cos.op,
		cos.Metrics.Successes,
		cos.Metrics.Objects,
		humanize.Bytes(uint64(cos.Metrics.Bytes)),
		cos.Folders,
		operationStatement,
		cos.details)
}
