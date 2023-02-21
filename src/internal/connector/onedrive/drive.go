package onedrive

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"
	msdrive "github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	gapi "github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/internal/connector/onedrive/api"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
)

var errFolderNotFound = clues.New("folder not found")

const (
	getDrivesRetries = 3

	// nextLinkKey is used to find the next link in a paged
	// graph response
	nextLinkKey              = "@odata.nextLink"
	itemChildrenRawURLFmt    = "https://graph.microsoft.com/v1.0/drives/%s/items/%s/children"
	itemByPathRawURLFmt      = "https://graph.microsoft.com/v1.0/drives/%s/items/%s:/%s"
	itemNotFoundErrorCode    = "itemNotFound"
	userMysiteURLNotFound    = "BadRequest Unable to retrieve user's mysite URL"
	userMysiteURLNotFoundMsg = "Unable to retrieve user's mysite URL"
	userMysiteNotFound       = "ResourceNotFound User's mysite not found"
	userMysiteNotFoundMsg    = "User's mysite not found"
	contextDeadlineExceeded  = "context deadline exceeded"
)

// DeltaUpdate holds the results of a current delta token.  It normally
// gets produced when aggregating the addition and removal of items in
// a delta-queriable folder.
// FIXME: This is same as exchange.api.DeltaUpdate
type DeltaUpdate struct {
	// the deltaLink itself
	URL string
	// true if the old delta was marked as invalid
	Reset bool
}

type drivePager interface {
	GetPage(context.Context) (gapi.PageLinker, error)
	SetNext(nextLink string)
	ValuesIn(gapi.PageLinker) ([]models.Driveable, error)
}

func PagerForSource(
	source driveSource,
	servicer graph.Servicer,
	resourceOwner string,
	fields []string,
) (drivePager, error) {
	switch source {
	case OneDriveSource:
		return api.NewUserDrivePager(servicer, resourceOwner, fields), nil
	case SharePointSource:
		return api.NewSiteDrivePager(servicer, resourceOwner, fields), nil
	default:
		return nil, errors.Errorf("unrecognized drive data source")
	}
}

func drives(
	ctx context.Context,
	pager drivePager,
	retry bool,
) ([]models.Driveable, error) {
	var (
		numberOfRetries = getDrivesRetries
		drives          = []models.Driveable{}
	)

	if !retry {
		numberOfRetries = 0
	}

	// Loop through all pages returned by Graph API.
	for {
		var (
			err  error
			page gapi.PageLinker
		)

		// Retry Loop for Drive retrieval. Request can timeout
		for i := 0; i <= numberOfRetries; i++ {
			page, err = pager.GetPage(ctx)
			if err != nil {
				// Various error handling. May return an error or perform a retry.
				// errMsg := support.ConnectorStackErrorTraceWrap(err, "").Error()
				// temporarily broken until ^ is fixed in next PR.
				errMsg := err.Error()
				if strings.Contains(errMsg, userMysiteURLNotFound) ||
					strings.Contains(errMsg, userMysiteURLNotFoundMsg) ||
					strings.Contains(errMsg, userMysiteNotFound) ||
					strings.Contains(errMsg, userMysiteNotFoundMsg) {
					logger.Ctx(ctx).Infof("resource owner does not have a drive")
					return make([]models.Driveable, 0), nil // no license or drives.
				}

				if strings.Contains(errMsg, contextDeadlineExceeded) && i < numberOfRetries {
					time.Sleep(time.Duration(3*(i+1)) * time.Second)
					continue
				}

				return nil, clues.Wrap(err, "retrieving drives").WithClues(ctx).With(graph.ErrData(err)...)
			}

			// No error encountered, break the retry loop so we can extract results
			// and see if there's another page to fetch.
			break
		}

		tmp, err := pager.ValuesIn(page)
		if err != nil {
			return nil, clues.Wrap(err, "extracting drives from response").WithClues(ctx).With(graph.ErrData(err)...)
		}

		drives = append(drives, tmp...)

		nextLink := ptr.Val(page.GetOdataNextLink())
		if len(nextLink) == 0 {
			break
		}

		pager.SetNext(nextLink)
	}

	logger.Ctx(ctx).Debugf("retrieved %d valid drives", len(drives))

	return drives, nil
}

// itemCollector functions collect the items found in a drive
type itemCollector func(
	ctx context.Context,
	driveID, driveName string,
	driveItems []models.DriveItemable,
	oldPaths map[string]string,
	newPaths map[string]string,
	excluded map[string]struct{},
	fileCollectionMap map[string]string,
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
		itemCollection = map[string]string{}
	)

	if !invalidPrevDelta {
		maps.Copy(newPaths, oldPaths)
		pager.SetNext(prevDelta)
	}

	for {
		page, err := pager.GetPage(ctx)

		if graph.IsErrInvalidDelta(err) {
			logger.Ctx(ctx).Infow("Invalid previous delta link", "link", prevDelta)

			invalidPrevDelta = true
			newPaths = map[string]string{}

			pager.Reset()

			continue
		}

		if err != nil {
			return DeltaUpdate{}, nil, nil, clues.Wrap(err, "getting page").WithClues(ctx).With(graph.ErrData(err)...)
		}

		vals, err := pager.ValuesIn(page)
		if err != nil {
			return DeltaUpdate{}, nil, nil, clues.Wrap(err, "extracting items from response").
				WithClues(ctx).
				With(graph.ErrData(err)...)
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

// getFolder will lookup the specified folder name under `parentFolderID`
func getFolder(
	ctx context.Context,
	service graph.Servicer,
	driveID, parentFolderID, folderName string,
) (models.DriveItemable, error) {
	// The `Children().Get()` API doesn't yet support $filter, so using that to find a folder
	// will be sub-optimal.
	// Instead, we leverage OneDrive path-based addressing -
	// https://learn.microsoft.com/en-us/graph/onedrive-addressing-driveitems#path-based-addressing
	// - which allows us to lookup an item by its path relative to the parent ID
	rawURL := fmt.Sprintf(itemByPathRawURLFmt, driveID, parentFolderID, folderName)
	builder := msdrive.NewItemsDriveItemItemRequestBuilder(rawURL, service.Adapter())

	var (
		foundItem models.DriveItemable
		err       error
	)

	foundItem, err = builder.Get(ctx, nil)
	if err != nil {
		if graph.IsErrDeletedInFlight(err) {
			return nil, clues.Stack(errFolderNotFound, err).WithClues(ctx).With(graph.ErrData(err)...)
		}

		return nil, clues.Wrap(err, "getting folder").WithClues(ctx).With(graph.ErrData(err)...)
	}

	// Check if the item found is a folder, fail the call if not
	if foundItem.GetFolder() == nil {
		return nil, clues.Stack(errFolderNotFound).WithClues(ctx).With(graph.ErrData(err)...)
	}

	return foundItem, nil
}

// Create a new item in the specified folder
func createItem(
	ctx context.Context,
	service graph.Servicer,
	driveID, parentFolderID string,
	newItem models.DriveItemable,
) (models.DriveItemable, error) {
	// Graph SDK doesn't yet provide a POST method for `/children` so we set the `rawUrl` ourselves as recommended
	// here: https://github.com/microsoftgraph/msgraph-sdk-go/issues/155#issuecomment-1136254310
	rawURL := fmt.Sprintf(itemChildrenRawURLFmt, driveID, parentFolderID)
	builder := msdrive.NewItemsRequestBuilder(rawURL, service.Adapter())

	newItem, err := builder.Post(ctx, newItem, nil)
	if err != nil {
		return nil, clues.Wrap(err, "creating item").WithClues(ctx).With(graph.ErrData(err)...)
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
	pager drivePager,
	prefix string,
	errs *fault.Bus,
) ([]*Displayable, error) {
	drives, err := drives(ctx, pager, true)
	if err != nil {
		return nil, errors.Wrap(err, "getting OneDrive folders")
	}

	var (
		folders = map[string]*Displayable{}
		el      = errs.Local()
	)

	for _, d := range drives {
		if el.Failure() != nil {
			break
		}

		var (
			id   = ptr.Val(d.GetId())
			name = ptr.Val(d.GetName())
		)

		ictx := clues.Add(ctx, "drive_id", id, "drive_name", name) // TODO: pii
		collector := func(
			innerCtx context.Context,
			driveID, driveName string,
			items []models.DriveItemable,
			oldPaths map[string]string,
			newPaths map[string]string,
			excluded map[string]struct{},
			itemCollection map[string]string,
			doNotMergeItems bool,
			errs *fault.Bus,
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

				if !strings.HasPrefix(*item.GetName(), prefix) {
					continue
				}

				// Add the item instead of the folder because the item has more
				// functionality.
				folders[itemID] = &Displayable{item}
			}

			return nil
		}

		_, _, _, err = collectItems(ictx, defaultItemPager(gs, id, ""), id, name, collector, map[string]string{}, "", errs)
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

func DeleteItem(
	ctx context.Context,
	gs graph.Servicer,
	driveID string,
	itemID string,
) error {
	err := gs.Client().DrivesById(driveID).ItemsById(itemID).Delete(ctx, nil)
	if err != nil {
		return clues.Wrap(err, "deleting item").
			WithClues(ctx).
			With("item_id", itemID).
			With(graph.ErrData(err)...)
	}

	return nil
}
