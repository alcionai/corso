package onedrive

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	gapi "github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/internal/connector/onedrive/api"
	odConsts "github.com/alcionai/corso/src/internal/connector/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	maxDrivesRetries = 3

	// nextLinkKey is used to find the next link in a paged
	// graph response
	nextLinkKey           = "@odata.nextLink"
	itemChildrenRawURLFmt = "https://graph.microsoft.com/v1.0/drives/%s/items/%s/children"
	itemNotFoundErrorCode = "itemNotFound"
)

// DeltaUpdate holds the results of a current delta token.  It normally
// gets produced when aggregating the addition and removal of items in
// a delta-queryable folder.
// FIXME: This is same as exchange.api.DeltaUpdate
type DeltaUpdate struct {
	// the deltaLink itself
	URL string
	// true if the old delta was marked as invalid
	Reset bool
}

func PagerForSource(
	source driveSource,
	servicer graph.Servicer,
	resourceOwner string,
	fields []string,
) (api.DrivePager, error) {
	switch source {
	case OneDriveSource:
		return api.NewUserDrivePager(servicer, resourceOwner, fields), nil
	case SharePointSource:
		return api.NewSiteDrivePager(servicer, resourceOwner, fields), nil
	default:
		return nil, clues.New("unrecognized drive data source")
	}
}

type pathPrefixerFunc func(driveID string) (path.Path, error)

func pathPrefixerForSource(
	tenantID, resourceOwner string,
	source driveSource,
) pathPrefixerFunc {
	cat := path.FilesCategory
	serv := path.OneDriveService

	if source == SharePointSource {
		cat = path.LibrariesCategory
		serv = path.SharePointService
	}

	return func(driveID string) (path.Path, error) {
		return path.Build(tenantID, resourceOwner, serv, cat, false, odConsts.DrivesPathDir, driveID, odConsts.RootPathDir)
	}
}

// itemCollector functions collect the items found in a drive
type itemCollector func(
	ctx context.Context,
	driveID, driveName string,
	driveItems []models.DriveItemable,
	oldPaths map[string]string,
	newPaths map[string]string,
	excluded map[string]struct{},
	itemCollections map[string]map[string]string,
	validPrevDelta bool,
	errs *fault.Bus,
) error

type itemPager interface {
	GetPage(context.Context) (gapi.DeltaPageLinker, error)
	SetNext(nextLink string)
	Reset()
	ValuesIn(gapi.DeltaPageLinker) ([]models.DriveItemable, error)
}

func defaultItemPager(
	servicer graph.Servicer,
	driveID, link string,
) itemPager {
	return api.NewItemPager(
		servicer,
		driveID,
		link,
		[]string{
			"content.downloadUrl",
			"createdBy",
			"createdDateTime",
			"file",
			"folder",
			"id",
			"lastModifiedDateTime",
			"name",
			"package",
			"parentReference",
			"root",
			"sharepointIds",
			"size",
			"deleted",
			"malware",
			"shared",
		},
	)
}

// collectItems will enumerate all items in the specified drive and hand them to the
// provided `collector` method
func collectItems(
	ctx context.Context,
	pager itemPager,
	driveID, driveName string,
	collector itemCollector,
	oldPaths map[string]string,
	prevDelta string,
	errs *fault.Bus,
) (DeltaUpdate, map[string]string, map[string]struct{}, error) {
	var (
		newDeltaURL      = ""
		newPaths         = map[string]string{}
		excluded         = map[string]struct{}{}
		invalidPrevDelta = len(prevDelta) == 0

		// itemCollection is used to identify which collection a
		// file belongs to. This is useful to delete a file from the
		// collection it was previously in, in case it was moved to a
		// different collection within the same delta query
		// drive ID -> item ID -> item ID
		itemCollection = map[string]map[string]string{
			driveID: {},
		}
	)

	if !invalidPrevDelta {
		maps.Copy(newPaths, oldPaths)
		pager.SetNext(prevDelta)
	}

	for {
		// assume delta urls here, which allows single-token consumption
		page, err := pager.GetPage(graph.ConsumeNTokens(ctx, graph.SingleGetOrDeltaLC))

		if graph.IsErrInvalidDelta(err) {
			logger.Ctx(ctx).Infow("Invalid previous delta link", "link", prevDelta)

			invalidPrevDelta = true
			newPaths = map[string]string{}

			pager.Reset()

			continue
		}

		if err != nil {
			return DeltaUpdate{}, nil, nil, graph.Wrap(ctx, err, "getting page")
		}

		vals, err := pager.ValuesIn(page)
		if err != nil {
			return DeltaUpdate{}, nil, nil, graph.Wrap(ctx, err, "extracting items from response")
		}

		err = collector(
			ctx,
			driveID,
			driveName,
			vals,
			oldPaths,
			newPaths,
			excluded,
			itemCollection,
			invalidPrevDelta,
			errs)
		if err != nil {
			return DeltaUpdate{}, nil, nil, err
		}

		nextLink, deltaLink := gapi.NextAndDeltaLink(page)

		if len(deltaLink) > 0 {
			newDeltaURL = deltaLink
		}

		// Check if there are more items
		if len(nextLink) == 0 {
			break
		}

		logger.Ctx(ctx).Debugw("Found nextLink", "link", nextLink)
		pager.SetNext(nextLink)
	}

	return DeltaUpdate{URL: newDeltaURL, Reset: invalidPrevDelta}, newPaths, excluded, nil
}

// Create a new item in the specified folder
func CreateItem(
	ctx context.Context,
	service graph.Servicer,
	driveID, parentFolderID string,
	newItem models.DriveItemable,
) (models.DriveItemable, error) {
	// Graph SDK doesn't yet provide a POST method for `/children` so we set the `rawUrl` ourselves as recommended
	// here: https://github.com/microsoftgraph/msgraph-sdk-go/issues/155#issuecomment-1136254310
	rawURL := fmt.Sprintf(itemChildrenRawURLFmt, driveID, parentFolderID)
	builder := drives.NewItemItemsRequestBuilder(rawURL, service.Adapter())

	newItem, err := builder.Post(ctx, newItem, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating item")
	}

	return newItem, nil
}

// newItem initializes a `models.DriveItemable` that can be used as input to `createItem`
func newItem(name string, folder bool) models.DriveItemable {
	itemToCreate := models.NewDriveItem()
	itemToCreate.SetName(&name)

	if folder {
		itemToCreate.SetFolder(models.NewFolder())
	} else {
		itemToCreate.SetFile(models.NewFile())
	}

	return itemToCreate
}

type Displayable struct {
	models.DriveItemable
}

func (op *Displayable) GetDisplayName() *string {
	return op.GetName()
}

// GetAllFolders returns all folders in all drives for the given user. If a
// prefix is given, returns all folders with that prefix, regardless of if they
// are a subfolder or top-level folder in the hierarchy.
func GetAllFolders(
	ctx context.Context,
	gs graph.Servicer,
	pager api.DrivePager,
	prefix string,
	errs *fault.Bus,
) ([]*Displayable, error) {
	drvs, err := api.GetAllDrives(ctx, pager, true, maxDrivesRetries)
	if err != nil {
		return nil, clues.Wrap(err, "getting OneDrive folders")
	}

	var (
		folders = map[string]*Displayable{}
		el      = errs.Local()
	)

	for _, d := range drvs {
		if el.Failure() != nil {
			break
		}

		var (
			id   = ptr.Val(d.GetId())
			name = ptr.Val(d.GetName())
		)

		ictx := clues.Add(ctx, "drive_id", id, "drive_name", clues.Hide(name))
		collector := func(
			_ context.Context,
			_, _ string,
			items []models.DriveItemable,
			_ map[string]string,
			_ map[string]string,
			_ map[string]struct{},
			_ map[string]map[string]string,
			_ bool,
			_ *fault.Bus,
		) error {
			for _, item := range items {
				// Skip the root item.
				if item.GetRoot() != nil {
					continue
				}

				// Only selecting folders right now, not packages.
				if item.GetFolder() == nil {
					continue
				}

				itemID := ptr.Val(item.GetId())
				if len(itemID) == 0 {
					logger.Ctx(ctx).Info("folder missing ID")
					continue
				}

				if !strings.HasPrefix(ptr.Val(item.GetName()), prefix) {
					continue
				}

				// Add the item instead of the folder because the item has more
				// functionality.
				folders[itemID] = &Displayable{item}
			}

			return nil
		}

		_, _, _, err = collectItems(
			ictx,
			defaultItemPager(gs, id, ""),
			id,
			name,
			collector,
			map[string]string{},
			"",
			errs)
		if err != nil {
			el.AddRecoverable(clues.Wrap(err, "enumerating items in drive"))
		}
	}

	res := make([]*Displayable, 0, len(folders))

	for _, f := range folders {
		res = append(res, f)
	}

	return res, el.Failure()
}

// deletes require unique http clients
// https://github.com/alcionai/corso/issues/2707
func DeleteItem(
	ctx context.Context,
	gs graph.Servicer,
	driveID string,
	itemID string,
) error {
	err := gs.Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting item").With("item_id", itemID)
	}

	return nil
}
