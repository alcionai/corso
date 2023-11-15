package groups

import (
	"context"
	stdlibpath "path"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/groups"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

var _ inject.ServiceHandler = &groupsHandler{}

func NewGroupsHandler(
	opts control.Options,
	apiClient api.Client,
	resourceGetter idname.GetResourceIDAndNamer,
) *groupsHandler {
	return &groupsHandler{
		baseGroupsHandler: baseGroupsHandler{
			opts:               opts,
			backupDriveIDNames: idname.NewCache(nil),
			backupSiteIDWebURL: idname.NewCache(nil),
		},
		apiClient:      apiClient,
		resourceGetter: resourceGetter,
	}
}

// ========================================================================== //
//                          baseGroupsHandler
// ========================================================================== //

// baseGroupsHandler contains logic for tracking data and doing operations
// (e.x. export) that don't require contact with external M356 services.
type baseGroupsHandler struct {
	opts control.Options

	backupDriveIDNames idname.CacheBuilder
	backupSiteIDWebURL idname.CacheBuilder
}

func (h *baseGroupsHandler) CacheItemInfo(v details.ItemInfo) {
	if v.Groups == nil {
		return
	}

	h.backupDriveIDNames.Add(v.Groups.DriveID, v.Groups.DriveName)
	h.backupSiteIDWebURL.Add(v.Groups.SiteID, v.Groups.WebURL)
}

// ProduceExportCollections will create the export collections for the
// given restore collections.
func (h *baseGroupsHandler) ProduceExportCollections(
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

	for _, restoreColl := range dcs {
		var (
			fp      = restoreColl.FullPath()
			cat     = fp.Category()
			folders = []string{cat.HumanString()}
			coll    export.Collectioner
		)

		switch cat {
		case path.ChannelMessagesCategory:
			folders = append(folders, fp.Folders()...)

			coll = groups.NewExportCollection(
				path.Builder{}.Append(folders...).String(),
				[]data.RestoreCollection{restoreColl},
				backupVersion,
				exportCfg,
				stats)

		case path.LibrariesCategory:
			drivePath, err := path.ToDrivePath(restoreColl.FullPath())
			if err != nil {
				return nil, clues.Wrap(err, "transforming path to drive path").WithClues(ctx)
			}

			driveName, ok := h.backupDriveIDNames.NameOf(drivePath.DriveID)
			if !ok {
				// This should not happen, but just in case
				logger.Ctx(ctx).With("drive_id", drivePath.DriveID).Info("drive name not found, using drive id")
				driveName = drivePath.DriveID
			}

			rfds := restoreColl.FullPath().Folders()
			siteName := rfds[1] // use siteID by default

			webURL, ok := h.backupSiteIDWebURL.NameOf(siteName)
			if !ok {
				// This should not happen, but just in case
				logger.Ctx(ctx).With("site_id", rfds[1]).Info("site weburl not found, using site id")
			}

			if len(webURL) != 0 {
				// We can't use the actual name anyways as it might
				// contain invalid characters. This should also avoid
				// possibility of name collisions.
				siteName = stdlibpath.Base(webURL)
			}

			baseDir := path.Builder{}.
				Append(folders...).
				Append(siteName).
				Append(driveName).
				Append(drivePath.Folders...)

			coll = drive.NewExportCollection(
				baseDir.String(),
				[]data.RestoreCollection{restoreColl},
				backupVersion,
				stats)
		default:
			el.AddRecoverable(
				ctx,
				clues.New("unsupported category for export").With("category", cat))

			continue
		}

		ec = append(ec, coll)
	}

	return ec, el.Failure()
}

// ========================================================================== //
//                              groupsHandler
// ========================================================================== //

// groupsHandler contains logic for handling data and performing operations
// (e.x. restore) regardless of whether they require contact with external M365
// services or not.
type groupsHandler struct {
	baseGroupsHandler
	apiClient      api.Client
	resourceGetter idname.GetResourceIDAndNamer
}

func (h *groupsHandler) IsServiceEnabled(
	ctx context.Context,
	resource string,
) (bool, error) {
	// TODO(ashmrtn): Move free function implementation to this function.
	res, err := IsServiceEnabled(ctx, h.apiClient.Groups(), resource)
	return res, clues.Stack(err).OrNil()
}

func (h *groupsHandler) PopulateProtectedResourceIDAndName(
	ctx context.Context,
	resource string, // Can be either ID or name.
	ins idname.Cacher,
) (idname.Provider, error) {
	if h.resourceGetter == nil {
		return nil, clues.Stack(graph.ErrNoResourceLookup).WithClues(ctx)
	}

	pr, err := h.resourceGetter.GetResourceIDAndNameFrom(ctx, resource, ins)

	return pr, clues.Wrap(err, "identifying resource owner").OrNil()
}
