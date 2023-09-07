package groups

import (
	"context"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
)

func NewExportCollection(
	baseDir string,
	backingCollection data.RestoreCollection,
	backupVersion int,
) export.Collectioner {
	return export.BaseCollection{
		BaseDir:           baseDir,
		BackingCollection: backingCollection,
		BackupVersion:     backupVersion,
		Stream:            streamItems,
	}
}

// streamItems streams the items in the backingCollection into the export stream chan
func streamItems(
	ctx context.Context,
	drc data.RestoreCollection,
	backupVersion int,
	ch chan<- export.Item,
) {
	defer close(ch)

	errs := fault.New(false)

	for item := range drc.Items(ctx, errs) {
		itemID := item.ID()

		// channel message items have no name
		name := itemID

		ch <- export.Item{
			ID:   itemID,
			Name: name,
			Body: item.ToReader(),
		}
	}

	items, recovered := errs.ItemsAndRecovered()

	// Return all the items that we failed to source from the persistence layer
	for _, err := range items {
		ch <- export.Item{
			ID:    err.ID,
			Error: &err,
		}
	}

	for _, err := range recovered {
		ch <- export.Item{
			Error: err,
		}
	}
}
