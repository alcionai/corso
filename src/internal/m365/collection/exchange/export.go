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
	config control.ExportConfig,
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

			reader := item.ToReader()
			content, err := io.ReadAll(reader)

			reader.Close()

			if err != nil {
				ch <- export.Item{
					ID:    id,
					Error: clues.Wrap(err, "reading data"),
				}

				continue
			}

			email, err := eml.FromJSON(ctx, content)
			if err != nil {
				ch <- export.Item{
					ID:    id,
					Error: clues.Wrap(err, "converting JSON to eml"),
				}

				continue
			}

			emlReader := io.NopCloser(bytes.NewReader([]byte(email)))
			body := data.ReaderWithStats(emlReader, path.EmailCategory, stats)

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
