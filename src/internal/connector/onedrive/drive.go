package onedrive

import (
	"context"
	"fmt"
	"strings"
	"time"

	msdrive "github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/connector/graph"
	gapi "github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/internal/connector/onedrive/api"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/logger"
)

var errFolderNotFound = errors.New("folder not found")

const (
	getDrivesRetries = 3

	// nextLinkKey is used to find the next link in a paged
	// graph response
	nextLinkKey             = "@odata.nextLink"
	itemChildrenRawURLFmt   = "https://graph.microsoft.com/v1.0/drives/%s/items/%s/children"
	itemByPathRawURLFmt     = "https://graph.microsoft.com/v1.0/drives/%s/items/%s:/%s"
	itemNotFoundErrorCode   = "itemNotFound"
	userMysiteURLNotFound   = "BadRequest Unable to retrieve user's mysite URL"
	userMysiteNotFound      = "ResourceNotFound User's mysite not found"
	contextDeadlineExceeded = "context deadline exceeded"
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
		err             error
		page            gapi.PageLinker
		numberOfRetries = getDrivesRetries
		drives          = []models.Driveable{}
	)

	if !retry {
		numberOfRetries = 0
	}

	// Loop through all pages returned by Graph API.
	for {
		// Retry Loop for Drive retrieval. Request can timeout
		for i := 0; i <= numberOfRetries; i++ {
			page, err = pager.GetPage(ctx)
			if err != nil {
				// Various error handling. May return an error or perform a retry.
				detailedError := err.Error()
				if strings.Contains(detailedError, userMysiteURLNotFound) ||
					strings.Contains(detailedError, userMysiteNotFound) {
					logger.Ctx(ctx).Infof("resource owner does not have a drive")
					return make([]models.Driveable, 0), nil // no license or drives.
				}

				if strings.Contains(detailedError, contextDeadlineExceeded) && i < numberOfRetries {
					time.Sleep(time.Duration(3*(i+1)) * time.Second)
					continue
				}

				return nil, errors.Wrapf(
					err,
					"failed to retrieve drives. details: %s",
					detailedError,
				)
			}

			// No error encountered, break the retry loop so we can extract results
			// and see if there's another page to fetch.
			break
		}

		tmp, err := pager.ValuesIn(page)
		if err != nil {
			return nil, errors.Wrap(err, "extracting drives from response")
		}

		drives = append(drives, tmp...)

		nextLink := gapi.NextLink(page)
		if len(nextLink) == 0 {
			break
		}

		pager.SetNext(nextLink)
	}

	logger.Ctx(ctx).Debugf("Found %d drives", len(drives))

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
) error

type itemPager interface {
	GetPage(context.Context) (gapi.DeltaPageLinker, error)
	SetNext(nextLink string)
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
			"size",
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
	prevDelta string,
) (DeltaUpdate, map[string]string, map[string]struct{}, error) {
	var (
		newDeltaURL = ""
		// TODO(ashmrtn): Eventually this should probably be a parameter so we can
		// take in previous paths.
		oldPaths         = map[string]string{}
		newPaths         = map[string]string{}
		excluded         = map[string]struct{}{}
		invalidPrevDelta = false
		triedPrevDelta   = false
	)

	maps.Copy(newPaths, oldPaths)

	if len(prevDelta) != 0 {
		pager.SetNext(prevDelta)
	}

	for {
		page, err := pager.GetPage(ctx)

		if !triedPrevDelta && graph.IsErrInvalidDelta(err) {
			logger.Ctx(ctx).Infow("Invalid previous delta link", "link", prevDelta)

			triedPrevDelta = true // TODO(meain): Do we need this check?
			invalidPrevDelta = true

			pager.SetNext("")

			continue
		}

		if err != nil {
			return DeltaUpdate{}, nil, nil, errors.Wrapf(
				err,
				"failed to query drive items. details: %s",
				support.ConnectorStackErrorTrace(err),
			)
		}

		vals, err := pager.ValuesIn(page)
		if err != nil {
			return DeltaUpdate{}, nil, nil, errors.Wrap(err, "extracting items from response")
		}

		err = collector(ctx, driveID, driveName, vals, oldPaths, newPaths, excluded)
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

	err = graph.RunWithRetry(func() error {
		foundItem, err = builder.Get(ctx, nil)
		return err
	})

	if err != nil {
		var oDataError *odataerrors.ODataError
		if errors.As(err, &oDataError) &&
			oDataError.GetError() != nil &&
			oDataError.GetError().GetCode() != nil &&
			*oDataError.GetError().GetCode() == itemNotFoundErrorCode {
			return nil, errors.WithStack(errFolderNotFound)
		}

		return nil, errors.Wrapf(err,
			"failed to get folder %s/%s. details: %s",
			parentFolderID,
			folderName,
			support.ConnectorStackErrorTrace(err),
		)
	}

	// Check if the item found is a folder, fail the call if not
	if foundItem.GetFolder() == nil {
		return nil, errors.WithStack(errFolderNotFound)
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
		return nil, errors.Wrapf(
			err,
			"failed to create item. details: %s",
			support.ConnectorStackErrorTrace(err),
		)
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
) ([]*Displayable, error) {
	drives, err := drives(ctx, pager, true)
	if err != nil {
		return nil, errors.Wrap(err, "getting OneDrive folders")
	}

	folders := map[string]*Displayable{}

	for _, d := range drives {
		_, _, _, err = collectItems(
			ctx,
			defaultItemPager(
				gs,
				*d.GetId(),
				"",
			),
			*d.GetId(),
			*d.GetName(),
			func(
				innerCtx context.Context,
				driveID, driveName string,
				items []models.DriveItemable,
				oldPaths map[string]string,
				newPaths map[string]string,
				excluded map[string]struct{},
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

					if item.GetId() == nil || len(*item.GetId()) == 0 {
						logger.Ctx(ctx).Warn("folder without ID")
						continue
					}

					if !strings.HasPrefix(*item.GetName(), prefix) {
						continue
					}

					// Add the item instead of the folder because the item has more
					// functionality.
					folders[*item.GetId()] = &Displayable{item}
				}

				return nil
			},
			"",
		)
		if err != nil {
			return nil, errors.Wrapf(err, "getting items for drive %s", *d.GetName())
		}
	}

	res := make([]*Displayable, 0, len(folders))

	for _, f := range folders {
		res = append(res, f)
	}

	return res, nil
}

func DeleteItem(
	ctx context.Context,
	gs graph.Servicer,
	driveID string,
	itemID string,
) error {
	err := gs.Client().DrivesById(driveID).ItemsById(itemID).Delete(ctx, nil)
	if err != nil {
		return errors.Wrapf(err, "deleting item with ID %s", itemID)
	}

	return nil
}
