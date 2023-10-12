package onedrive

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ inject.ServiceHandler = &onedriveHandler{}

func NewOneDriveHandler(
	opts control.Options,
) *onedriveHandler {
	return &onedriveHandler{
		opts: opts,
	}
}

type onedriveHandler struct {
	opts control.Options
}

func (h *onedriveHandler) CacheItemInfo(v details.ItemInfo) {}

// ProduceExportCollections will create the export collections for the
// given restore collections.
func (h *onedriveHandler) ProduceExportCollections(
	ctx context.Context,
	backupVersion int,
	exportCfg control.ExportConfig,
	dcs []data.RestoreCollection,
	stats *data.ExportStats,
	errs *fault.Bus,
) ([]export.Collectioner, error) {
	var (
		el = errs.Local()
		ec = make([]export.Collectioner, 0, len(dcs))
	)

	for _, dc := range dcs {
		drivePath, err := path.ToDrivePath(dc.FullPath())
		if err != nil {
			return nil, clues.Wrap(err, "transforming path to drive path").WithClues(ctx)
		}

		baseDir := path.Builder{}.Append(drivePath.Folders...)

		ec = append(
			ec,
			drive.NewExportCollection(
				baseDir.String(),
				[]data.RestoreCollection{dc},
				backupVersion,
				stats))
	}

	return ec, el.Failure()
}
