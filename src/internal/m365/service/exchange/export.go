package exchange

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ inject.ServiceHandler = &baseExchangeHandler{}

func NewExchangeHandler(
	opts control.Options,
) *baseExchangeHandler {
	return &baseExchangeHandler{
		opts: opts,
	}
}

type baseExchangeHandler struct {
	opts control.Options
}

func (h *baseExchangeHandler) CacheItemInfo(v details.ItemInfo) {}

// ProduceExportCollections will create the export collections for the
// given restore collections.
func (h *baseExchangeHandler) ProduceExportCollections(
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
		category := dc.FullPath().Category()

		switch category {
		case path.EmailCategory:
			folders := dc.FullPath().Folders()
			pth := path.Builder{}.Append(path.EmailCategory.HumanString()).Append(folders...)

			ec = append(
				ec,
				exchange.NewExportCollection(
					pth.String(),
					[]data.RestoreCollection{dc},
					backupVersion,
					stats))
		case path.EventsCategory, path.ContactsCategory:
			logger.Ctx(ctx).With("category", category.String()).Debugw("Skipping restore for category")
		default:
			return nil, clues.New("data category not supported").
				With("category", category).
				WithClues(ctx)
		}
	}

	return ec, el.Failure()
}
