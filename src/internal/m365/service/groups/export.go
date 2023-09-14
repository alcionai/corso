package groups

import (
	"context"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/groups"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ProduceExportCollections will create the export collections for the
// given restore collections.
func ProduceExportCollections(
	ctx context.Context,
	backupVersion int,
	exportCfg control.ExportConfig,
	opts control.Options,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Bus,
) ([]export.Collectioner, error) {
	var (
		el = errs.Local()
		ec = make([]export.Collectioner, 0, len(dcs))
	)

	for _, restoreColl := range dcs {
		var (
			fp      = restoreColl.FullPath()
			cat     = fp.Category()
			folders = []string{cat.String()}
		)

		switch cat {
		case path.ChannelMessagesCategory:
			folders = append(folders, fp.Folders()...)
		}

		coll := groups.NewExportCollection(
			path.Builder{}.Append(folders...).String(),
			[]data.RestoreCollection{restoreColl},
			backupVersion,
			exportCfg)

		ec = append(ec, coll)
	}

	return ec, el.Failure()
}
