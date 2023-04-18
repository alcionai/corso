package support

import (
	"context"
	"fmt"

	"github.com/dustin/go-humanize"
)

// ConnectorOperationStatus is a data type used to describe the state of
// the sequence of operations.
// @param ObjectCount integer representation of how many objects have downloaded or uploaded.
// @param Successful: Number of objects that are sent through the connector without incident.
// @param incomplete: Bool representation of whether all intended items were download or uploaded.
// @param bytes: represents the total number of bytes that have been downloaded or uploaded.
type ConnectorOperationStatus struct {
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
	details string,
) *ConnectorOperationStatus {
	status := ConnectorOperationStatus{
		Folders: folders,
		Metrics: cm,
		details: details,
		op:      op,
	}

	return &status
}

// Function signature for a status updater
// Used to define a function that an async connector task can call
// to on completion with its ConnectorOperationStatus
type StatusUpdater func(*ConnectorOperationStatus)

// MergeStatus combines ConnectorOperationsStatus value into a single status
func MergeStatus(one, two ConnectorOperationStatus) ConnectorOperationStatus {
	if one.op == OpUnknown {
		return two
	}

	if two.op == OpUnknown {
		return one
	}

	status := ConnectorOperationStatus{
		Folders: one.Folders + two.Folders,
		Metrics: CombineMetrics(one.Metrics, two.Metrics),
		details: one.details + ", " + two.details,
		op:      one.op,
	}

	return status
}

func (cos *ConnectorOperationStatus) String() string {
	var operationStatement string

	switch cos.op {
	case Backup:
		operationStatement = "Downloaded from "
	case Restore:
		operationStatement = "Restored content to "
	}

	return fmt.Sprintf("Action: %s performed on %d of %d objects (%s) within %d directories.  %s %s",
		cos.op.String(),
		cos.Metrics.Successes,
		cos.Metrics.Objects,
		humanize.Bytes(uint64(cos.Metrics.Bytes)),
		cos.Folders,
		operationStatement,
		cos.details)
}
