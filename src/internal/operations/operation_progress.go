package operations

type (
	progressChan chan string
	errorChan    chan error
)

// opProgress allows downstream writers to communicate async progress and
// errors to the operation.  Per-process wrappers of operation are required
// to implement receivers for each channel.
type opProgress struct {
	progressChan
	errorChan
}

func newOpProgress() *opProgress {
	return &opProgress{
		progressChan: make(progressChan),
		errorChan:    make(errorChan),
	}
}

// Report transmits a progress report to the operation.
func (rch progressChan) Report(rpt string) {
	if rch != nil {
		rch <- rpt
	}
}

// Error transmits an error report to the operation.
func (ech errorChan) Error(err error) {
	if ech != nil {
		ech <- err
	}
}

// Close closes all communication channels used by opProgress.
func (op *opProgress) Close() {
	if op.progressChan != nil {
		close(op.progressChan)
		op.progressChan = nil
	}
	if op.errorChan != nil {
		close(op.errorChan)
		op.errorChan = nil
	}
}
