package groups

import (
	"context"
	"errors"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ConsumeRestoreCollections will restore the specified data collections into OneDrive
func (h *groupsHandler) ConsumeRestoreCollections(
	ctx context.Context,
	rcc inject.RestoreConsumerConfig,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.Details, *data.CollectionStats, error) {
	if len(dcs) == 0 {
		return nil, nil, clues.WrapWC(ctx, data.ErrNoData, "performing restore")
	}

	// TODO(ashmrtn): We should stop relying on the context for rate limiter stuff
	// and instead configure this when we make the handler instance. We can't
	// initialize it in the NewHandler call right now because those functions
	// aren't (and shouldn't be) returning a context along with the handler. Since
	// that call isn't directly calling into this function even if we did
	// initialize the rate limiter there it would be lost because it wouldn't get
	// stored in an ancestor of the context passed to this function.
	ctx = graph.BindRateLimiterConfig(
		ctx,
		graph.LimiterCfg{Service: path.GroupsService})

	var (
		deets          = &details.Builder{}
		restoreMetrics support.CollectionMetrics
		caches         = drive.NewRestoreCaches(h.backupDriveIDNames)
		lrh            = drive.NewSiteRestoreHandler(
			h.apiClient,
			rcc.Selector.PathService())
		el                = errs.Local()
		webURLToSiteNames = map[string]string{}
	)

	// Reorder collections so that the parents directories are created
	// before the child directories; a requirement for permissions.
	data.SortRestoreCollections(dcs)

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		if el.Failure() != nil {
			break
		}

		var (
			err      error
			siteName string
			category = dc.FullPath().Category()
			metrics  support.CollectionMetrics
			ictx     = clues.Add(ctx,
				"category", category,
				"restore_location", clues.Hide(rcc.RestoreConfig.Location),
				"protected_resource", clues.Hide(dc.FullPath().ProtectedResource()),
				"full_path", dc.FullPath())
		)

		switch dc.FullPath().Category() {
		case path.LibrariesCategory:
			siteID := dc.FullPath().Folders()[1]

			webURL, ok := h.backupSiteIDWebURL.NameOf(siteID)
			if !ok {
				// This should not happen, but just in case
				logger.Ctx(ictx).With("site_id", siteID).Info("site weburl not found, using site id")
			}

			siteName, err = getSiteName(ictx, siteID, webURL, h.apiClient.Sites(), webURLToSiteNames)
			if err != nil {
				el.AddRecoverable(ictx, clues.Wrap(err, "getting site").
					With("web_url", webURL, "site_id", siteID))
			} else if len(siteName) == 0 {
				// Site was deleted in between and restore and is not
				// available anymore.
				continue
			}

			pr := idname.NewProvider(siteID, siteName)
			srcc := inject.RestoreConsumerConfig{
				BackupVersion:     rcc.BackupVersion,
				Options:           rcc.Options,
				ProtectedResource: pr,
				RestoreConfig:     rcc.RestoreConfig,
				Selector:          rcc.Selector,
			}

			err = caches.Populate(ictx, lrh, srcc.ProtectedResource.ID())
			if err != nil {
				return nil, nil, clues.Wrap(err, "initializing restore caches")
			}

			users, ierr := h.apiClient.Users().GetAllIDsAndNames(ctx, errs)
			if err != nil {
				return nil, nil, clues.Wrap(ierr, "getting users")
			}

			groups, ierr := h.apiClient.Groups().GetAllIDsAndNames(ctx, errs)
			if err != nil {
				return nil, nil, clues.Wrap(ierr, "getting groups")
			}

			caches.AvailableEntities.Users = users
			caches.AvailableEntities.Groups = groups

			metrics, err = drive.RestoreCollection(
				ictx,
				lrh,
				srcc,
				dc,
				caches,
				deets,
				control.DefaultRestoreContainerName(dttm.HumanReadableDriveItem),
				errs,
				ctr)
		case path.ChannelMessagesCategory:
			// Message cannot be restored as of now using Graph API.
			logger.Ctx(ictx).Debug("Skipping restore for channel messages")
		default:
			return nil, nil, clues.NewWC(ictx, "data category not supported").
				With("category", category)
		}

		restoreMetrics = support.CombineMetrics(restoreMetrics, metrics)

		if err != nil {
			el.AddRecoverable(ictx, err)
		}

		if errors.Is(err, context.Canceled) {
			break
		}
	}

	status := support.CreateStatus(
		ctx,
		support.Restore,
		len(dcs),
		restoreMetrics,
		rcc.RestoreConfig.Location)

	return deets.Details(), status.ToCollectionStats(), el.Failure()
}

func getSiteName(
	ctx context.Context,
	siteID string,
	webURL string,
	ac api.GetByIDer[models.Siteable],
	webURLToSiteNames map[string]string,
) (string, error) {
	siteName, ok := webURLToSiteNames[webURL]
	if ok {
		return siteName, nil
	}

	site, err := ac.GetByID(ctx, siteID, api.CallConfig{})
	if err != nil {
		webURLToSiteNames[webURL] = ""

		if graph.IsErrSiteNotFound(err) {
			// TODO(meain): Should we surface this to the user somehow?
			// In case a site that we had previously backed up was
			// deleted, skip that site with a warning.
			logger.Ctx(ctx).With("web_url", webURL, "site_id", siteID).
				Info("Site does not exist, skipping restore.")

			return "", nil
		}

		return "", err
	}

	siteName = ptr.Val(site.GetDisplayName())
	webURLToSiteNames[webURL] = siteName

	return siteName, nil
}
