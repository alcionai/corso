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
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/metrics"
	"github.com/alcionai/corso/src/pkg/path"
)

func NewExportCollection(
	baseDir string,
	backingCollection []data.RestoreCollection,
	backupVersion int,
	stats *metrics.ExportStats,
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
	stats *metrics.ExportStats,
) {
	defer close(ch)

	errs := fault.New(false)

	for _, rc := range drc {
		ictx := clues.Add(ctx, "path_short_ref", rc.FullPath().ShortRef())

		for item := range rc.Items(ictx, errs) {
			id := item.ID()
			name := id + ".eml"

			stats.UpdateResourceCount(path.EmailCategory)

			reader := item.ToReader()
			content, err := io.ReadAll(reader)

			reader.Close()

			if err != nil {
				err = clues.Wrap(err, "reading export item")

				logger.CtxErr(ctx, err).Info("processing collection item")

				ch <- export.Item{
					ID:    id,
					Error: err,
				}

				continue
			}

			email, err := eml.FromJSON(ictx, content)
			if err != nil {
				err = clues.Wrap(err, "converting JSON to eml")

				logger.CtxErr(ctx, err).Info("processing collection item")

				ch <- export.Item{
					ID:    id,
					Error: err,
				}

				continue
			}

			emlReader := io.NopCloser(bytes.NewReader([]byte(email)))
			body := metrics.ReaderWithStats(emlReader, path.EmailCategory, stats)

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
