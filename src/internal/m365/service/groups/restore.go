package groups

import (
	"context"
	"errors"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ConsumeRestoreCollections will restore the specified data collections into OneDrive
func ConsumeRestoreCollections(
	ctx context.Context,
	rcc inject.RestoreConsumerConfig,
	ac api.Client,
	backupDriveIDNames idname.Cacher,
	backupSiteIDWebURL idname.Cacher,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Bus,
	ctr *count.Bus,
) (*support.ControllerOperationStatus, error) {
	var (
		restoreMetrics support.CollectionMetrics
		caches         = drive.NewRestoreCaches(backupDriveIDNames)
		lrh            = drive.NewLibraryRestoreHandler(ac, rcc.Selector.PathService())
		el             = errs.Local()
		siteNames      = map[string]string{}
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

			webURL, ok := backupSiteIDWebURL.NameOf(siteID)
			if !ok {
				// This should not happen, but just in case
				logger.Ctx(ctx).With("site_id", siteID).Info("site weburl not found, using site id")
			}

			// In case a site that we had previously backed up was
			// deleted, skip that site with a warning.
			siteName, ok := siteNames[webURL]
			if !ok {
				site, err := ac.Sites().GetByID(ctx, siteID, api.CallConfig{})
				if err != nil {
					siteNames[webURL] = ""

					if graph.IsErrSiteCouldNotBeFound(err) {
						// TODO(meain): Should we surface this to the user somehow?
						logger.Ctx(ctx).With("web_url", webURL, "site_id", siteID).
							Info("Site does not exist, skipping restore.")
					} else {
						el.AddRecoverable(ctx, clues.Wrap(err, "getting site").
							With("web_url", webURL, "site_id", siteID))
					}

					continue
				}

				siteName = ptr.Val(site.GetDisplayName())
				siteNames[webURL] = siteName
			}

			if ok && len(siteName) == 0 {
				// We have previously seen it but the site was
				// unavailable
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

			err = caches.Populate(ctx, lrh, srcc.ProtectedResource.ID())
			if err != nil {
				return nil, clues.Wrap(err, "initializing restore caches")
			}

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
			logger.Ctx(ctx).Debug("Skipping restore for channel messages")
		default:
			return nil, clues.New("data category not supported").
				With("category", category).
				WithClues(ictx)
		}

		restoreMetrics = support.CombineMetrics(restoreMetrics, metrics)

		if err != nil {
			el.AddRecoverable(ctx, err)
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

	return status, el.Failure()
}
