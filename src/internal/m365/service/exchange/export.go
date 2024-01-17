package exchange

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/metrics"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ inject.ServiceHandler = &exchangeHandler{}

func NewExchangeHandler(
	apiClient api.Client,
	resourceClient idname.GetResourceIDAndNamer,
) *exchangeHandler {
	return &exchangeHandler{
		baseExchangeHandler: baseExchangeHandler{},
		apiClient:           apiClient,
		resourceClient:      resourceClient,
	}
}

// ========================================================================== //
//                        baseExchangeHandler
// ========================================================================== //

// baseExchangeHandler contains logic for tracking data and doing operations
// (e.x. export) that don't require contact with external M356 services.
type baseExchangeHandler struct{}

func (h *baseExchangeHandler) CacheItemInfo(v details.ItemInfo) {}

// ProduceExportCollections will create the export collections for the
// given restore collections.
func (h *baseExchangeHandler) ProduceExportCollections(
	ctx context.Context,
	backupVersion int,
	exportCfg control.ExportConfig,
	dcs []data.RestoreCollection,
	stats *metrics.ExportStats,
	errs *fault.Bus,
) ([]export.Collectioner, error) {
	var (
		el = errs.Local()
		ec = make([]export.Collectioner, 0, len(dcs))
	)

	for _, dc := range dcs {
		category := dc.FullPath().Category()

		switch category {
		case path.ContactsCategory, path.EmailCategory, path.EventsCategory:
			folders := dc.FullPath().Folders()
			pth := path.Builder{}.Append(category.HumanString()).Append(folders...)

			ec = append(
				ec,
				exchange.NewExportCollection(
					pth.String(),
					[]data.RestoreCollection{dc},
					backupVersion,
					stats))
		default:
			return nil, clues.NewWC(ctx, "data category not supported").
				With("category", category)
		}
	}

	return ec, el.Failure()
}

// ========================================================================== //
//                            exchangeHandler
// ========================================================================== //

// exchangeHandler contains logic for handling data and performing operations
// (e.x. restore) regardless of whether they require contact with external M365
// services or not.
type exchangeHandler struct {
	baseExchangeHandler
	apiClient      api.Client
	resourceClient idname.GetResourceIDAndNamer
}

func (h *exchangeHandler) IsServiceEnabled(
	ctx context.Context,
	resourceID string,
) (bool, error) {
	// TODO(ashmrtn): Move free function implementation to this function.
	res, err := IsServiceEnabled(ctx, h.apiClient.Users(), resourceID)
	return res, clues.Stack(err).OrNil()
}

func (h *exchangeHandler) PopulateProtectedResourceIDAndName(
	ctx context.Context,
	resourceID string, // Can be either ID or name.
	ins idname.Cacher,
) (idname.Provider, error) {
	if h.resourceClient == nil {
		return nil, clues.StackWC(ctx, resource.ErrNoResourceLookup)
	}

	pr, err := h.resourceClient.GetResourceIDAndNameFrom(ctx, resourceID, ins)

	return pr, clues.Wrap(err, "identifying resource owner").OrNil()
}
