package operations

import (
	"context"

	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
)

// LogFaultErrors is a helper function that logs all entries in the Errors struct.
func LogFaultErrors(ctx context.Context, fe *fault.Errors, prefix string) {
	if fe == nil {
		return
	}

	var (
		log        = logger.Ctx(ctx)
		pfxMsg     = prefix + ":"
		li, ls, lr = len(fe.Items), len(fe.Skipped), len(fe.Recovered)
	)

	if fe.Failure == nil && li+ls+lr == 0 {
		log.Info(pfxMsg, "no errors")
		return
	}

	if fe.Failure != nil {
		log.With("error", fe.Failure).Error(pfxMsg, "primary failure")
	}

	for i, item := range fe.Items {
		log.With("failed_item", item).Errorf("%s item failure %d of %d", pfxMsg, i+1, li)
	}

	for i, item := range fe.Skipped {
		log.With("skipped_item", item).Errorf("%s skipped item %d of %d", pfxMsg, i+1, ls)
	}

	for i, err := range fe.Recovered {
		log.With("recovered_error", err).Errorf("%s recoverable error %d of %d", pfxMsg, i+1, lr)
	}
}
