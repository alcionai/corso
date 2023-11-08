package exchange

import (
	"bytes"
	"context"
	"io"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/converters/eml"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func NewExportCollection(
	baseDir string,
	backingCollection []data.RestoreCollection,
	backupVersion int,
	stats *data.ExportStats,
) export.Collectioner {
	return export.BaseCollection{
		BaseDir:           baseDir,
		BackingCollection: backingCollection,
		BackupVersion:     backupVersion,
		Stream:            streamItems,
		Stats:             stats,
	}
}

// streamItems streams the streamItems in the backingCollection into the export stream chan
func streamItems(
	ctx context.Context,
	drc []data.RestoreCollection,
	backupVersion int,
	cec control.ExportConfig,
	ch chan<- export.Item,
	stats *data.ExportStats,
) {
	defer close(ch)

	errs := fault.New(false)

	for _, rc := range drc {
		for item := range rc.Items(ctx, errs) {
			id := item.ID()
			name := id + ".eml"

			stats.UpdateResourceCount(path.EmailCategory)

			content, err := io.ReadAll(item.ToReader())
			if err != nil {
				ch <- export.Item{
					ID:    id,
					Error: clues.Wrap(err, "reading data"),
				}

				continue
			}

			msg, err := api.BytesToMessageable(content)
			if err != nil {
				ch <- export.Item{
					ID:    id,
					Error: clues.Wrap(err, "parsing email"),
				}

				continue
			}

			email, err := eml.ToEml(msg)
			if err != nil {
				ch <- export.Item{
					ID:    id,
					Error: clues.Wrap(err, "converting to eml"),
				}

				continue
			}

			reader := io.NopCloser(bytes.NewReader([]byte(email)))
			body := data.ReaderWithStats(reader, path.EmailCategory, stats)

			ch <- export.Item{
				ID:   id,
				Name: name,
				Body: body,
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
}
