package operations

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/store"
)

type opStatus int

const (
	Unknown opStatus = iota
	InProgress
	Successful
	Failed
)

// --------------------------------------------------------------------------------
// Operation Core
// --------------------------------------------------------------------------------

// An operation tracks the in-progress workload of some long-running process.
// Specific processes (eg: backups, restores, etc) are expected to wrap operation
// with process specific details.
type operation struct {
	ID        uuid.UUID `json:"id"`        // system generated identifier
	CreatedAt time.Time `json:"createdAt"` // datetime of the operation's creation
	Options   Options   `json:"options"`
	Status    opStatus  `json:"status"`

	kopia *kopia.Wrapper
	store *store.Wrapper
}

const (
	onErrBestEffort = "best-effort"
	onErrFailFast   = "fast-fail"
)

// Options configure some parameters of the operation
type Options struct {
	OnError string `json:"onError"`
	// todo: collision handling
}

func NewOptions(fastFail bool) Options {
	oe := onErrBestEffort
	if fastFail {
		oe = onErrFailFast
	}
	return Options{
		OnError: oe,
	}
}

func newOperation(
	opts Options,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
) operation {
	return operation{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Options:   opts,
		kopia:     kw,
		store:     sw,
		Status:    InProgress,
	}
}

func (op operation) validate() error {
	if op.kopia == nil {
		return errors.New("missing kopia connection")
	}
	if op.store == nil {
		return errors.New("missing modelstore")
	}
	return nil
}

// --------------------------------------------------------------------------------
// Results
// --------------------------------------------------------------------------------

// Summary tracks the total files touched and errors produced
// during an operation.
type summary struct {
	ItemsRead    int   `json:"itemsRead,omitempty"`
	ItemsWritten int   `json:"itemsWritten,omitempty"`
	ReadErrors   error `json:"readErrors,omitempty"`
	WriteErrors  error `json:"writeErrors,omitempty"`
}

// Metrics tracks performance details such as timing, throughput, etc.
type metrics struct {
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
}
