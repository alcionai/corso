package status

import (
	"context"
	"fmt"

	"github.com/dustin/go-humanize"
)

type Operation int

//go:generate stringer -type=Operation
const (
	OpUnknown Operation = iota
	Backup
	Restore
)

// ConnectorStatus is a data type used to describe the state of the sequence of operations.
type ConnectorStatus struct {
	Metrics Counts

	details    string
	incomplete bool
	op         Operation
}

type Counts struct {
	Bytes                       int64
	Folders, Objects, Successes int
}

func CombineCounts(a, b Counts) Counts {
	return Counts{
		Bytes:     a.Bytes + b.Bytes,
		Folders:   a.Folders + b.Folders,
		Objects:   a.Objects + b.Objects,
		Successes: a.Successes + b.Successes,
	}
}

// Constructor for ConnectorStatus. If the counts do not agree, an error is returned.
func New(
	ctx context.Context,
	op Operation,
	cs Counts,
	details string,
	incomplete bool,
) ConnectorStatus {
	status := ConnectorStatus{
		Metrics:    cs,
		details:    details,
		incomplete: incomplete,
		op:         op,
	}

	return status
}

// Combine aggregates both ConnectorStatus value into a single status.
func Combine(one, two ConnectorStatus) ConnectorStatus {
	if one.op == OpUnknown {
		return two
	}

	if two.op == OpUnknown {
		return one
	}

	status := ConnectorStatus{
		Metrics:    CombineCounts(one.Metrics, two.Metrics),
		details:    one.details + ", " + two.details,
		incomplete: one.incomplete || two.incomplete,
		op:         one.op,
	}

	return status
}

func (cos ConnectorStatus) String() string {
	var operationStatement string

	switch cos.op {
	case Backup:
		operationStatement = "Downloaded from "
	case Restore:
		operationStatement = "Restored content to "
	}

	var incomplete string
	if cos.incomplete {
		incomplete = "Incomplete "
	}

	message := fmt.Sprintf("%sAction: %s performed on %d of %d objects (%s) within %d directories.",
		incomplete,
		cos.op.String(),
		cos.Metrics.Successes,
		cos.Metrics.Objects,
		humanize.Bytes(uint64(cos.Metrics.Bytes)),
		cos.Metrics.Folders)

	return message + " " + operationStatement + cos.details
}
