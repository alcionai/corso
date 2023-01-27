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

type drivePager interface {
	GetPage(context.Context) (api.PageLinker, error)
	SetNext(nextLink string)
	ValuesIn(api.PageLinker) ([]models.Driveable, error)
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

	for done := false; !done; {
		// Retry Loop for Drive retrieval. Request can timeout
		for i := 0; i <= numberOfRetries; i++ {
			page, err := pager.GetPage(ctx)
			if err == nil {
				// Success path, break out of inner loop at the end.
				tmp, err := pager.ValuesIn(page)
				if err != nil {
					return nil, errors.Wrap(err, "extracting user drives from response")
				}

				drives = append(drives, tmp...)

				nextLink := page.GetOdataNextLink()
				if nextLink == nil || len(*nextLink) == 0 {
					done = true
					break
				}

				pager.SetNext(*nextLink)

				break
			}

			// Various error handling. May return an error or perform a retry.
			detailedError := support.ConnectorStackErrorTrace(err)
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
	}

	logger.Ctx(ctx).Debugf("Found %d drives", len(drives))

	return drives, nil
}

// Enumerates the drives for the specified user
func drives(
	ctx context.Context,
	service graph.Servicer,
	resourceOwner string,
	source driveSource,
) ([]models.Driveable, error) {
	switch source {
	case OneDriveSource:
		return userDrives(ctx, service, resourceOwner)
	case SharePointSource:
		return siteDrives(ctx, service, resourceOwner)
	default:
		return nil, errors.Errorf("unrecognized drive data source")
	}
}

func siteDrives(ctx context.Context, service graph.Servicer, site string) ([]models.Driveable, error) {
	var (
		drives []models.Driveable

		// TODO(ashmrtn): Pass this in instead of creating it here so it can be
		// mocked for testing.
		drivePager = api.NewSiteDrivePager(
			service,
			site,
			[]string{
				"id",
				"name",
				"weburl",
				"system",
			},
		)
	)

	for {
		page, err := drivePager.GetPage(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to retrieve site drives. site: %s, details: %s",
				site, support.ConnectorStackErrorTrace(err))
		}

		tmp, err := drivePager.ValuesIn(page)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"extracting site drives from response for site %s",
				site,
			)
		}

		drives = append(drives, tmp...)

		nextLink := page.GetOdataNextLink()
		if nextLink == nil || len(*nextLink) == 0 {
			break
		}

		drivePager.SetNext(*nextLink)
	}

	return drives, nil
}

func userDrives(ctx context.Context, service graph.Servicer, user string) ([]models.Driveable, error) {
	var (
		numberOfRetries = 3
		drives          = []models.Driveable{}
	)

	// TODO(ashmrtn): Pass this in instead of creating it here so it can be
	// mocked for testing.
	drivePager := api.NewUserDrivePager(service, user, nil)

	for done := false; !done; {
		// Retry Loop for Drive retrieval. Request can timeout
		for i := 0; i <= numberOfRetries; i++ {
			page, err := drivePager.GetPage(ctx)
			if err == nil {
				// Success path, break out of inner loop at the end.
				tmp, err := drivePager.ValuesIn(page)
				if err != nil {
					return nil, errors.Wrapf(
						err,
						"extracting user drives from response for user %s",
						user,
					)
				}

				drives = append(drives, tmp...)

				nextLink := page.GetOdataNextLink()
				if nextLink == nil || len(*nextLink) == 0 {
					done = true
					break
				}

				drivePager.SetNext(*nextLink)

				break
			}

			// Various error handling. May return an error or perform a retry.
			detailedError := support.ConnectorStackErrorTrace(err)
			if strings.Contains(detailedError, userMysiteURLNotFound) ||
				strings.Contains(detailedError, userMysiteNotFound) {
				logger.Ctx(ctx).Infof("User %s does not have a drive", user)
				return make([]models.Driveable, 0), nil // no license
			}

			if strings.Contains(detailedError, "context deadline exceeded") && i < numberOfRetries {
				time.Sleep(time.Duration(3*(i+1)) * time.Second)
				continue
			}

			return nil, errors.Wrapf(
				err,
				"failed to retrieve user drives. user: %s, details: %s",
				user,
				detailedError,
			)
		}
	}

	logger.Ctx(ctx).Debugf("Found %d drives for user %s", len(drives), user)

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

// collectItems will enumerate all items in the specified drive and hand them to the
// provided `collector` method
func collectItems(
	ctx context.Context,
	service graph.Servicer,
	driveID, driveName string,
	collector itemCollector,
) (string, map[string]string, map[string]struct{}, error) {
	var (
		newDeltaURL = ""
		// TODO(ashmrtn): Eventually this should probably be a parameter so we can
		// take in previous paths.
		oldPaths = map[string]string{}
		newPaths = map[string]string{}
		excluded = map[string]struct{}{}
	)

	maps.Copy(newPaths, oldPaths)

	// TODO: Specify a timestamp in the delta query
	// https://docs.microsoft.com/en-us/graph/api/driveitem-delta?
	// view=graph-rest-1.0&tabs=http#example-4-retrieving-delta-results-using-a-timestamp
	builder := service.Client().DrivesById(driveID).Root().Delta()
	pageCount := int32(999) // max we can do is 999
	requestFields := []string{
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
	}
	requestConfig := &msdrives.ItemRootDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &msdrives.ItemRootDeltaRequestBuilderGetQueryParameters{
			Top:    &pageCount,
			Select: requestFields,
		},
	}

	for {
		r, err := builder.Get(ctx, requestConfig)
		if err != nil {
			return "", nil, nil, errors.Wrapf(
				err,
				"failed to query drive items. details: %s",
				support.ConnectorStackErrorTrace(err),
			)
		}

		err = collector(ctx, driveID, driveName, r.GetValue(), oldPaths, newPaths, excluded)
		if err != nil {
			return "", nil, nil, err
		}

		if r.GetOdataDeltaLink() != nil && len(*r.GetOdataDeltaLink()) > 0 {
			newDeltaURL = *r.GetOdataDeltaLink()
		}

		// Check if there are more items
		nextLink := r.GetOdataNextLink()
		if nextLink == nil {
			break
		}

		logger.Ctx(ctx).Debugf("Found %s nextLink", *nextLink)
		builder = msdrives.NewItemRootDeltaRequestBuilder(*nextLink, service.Adapter())
	}

	return newDeltaURL, newPaths, excluded, nil
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

	foundItem, err := builder.Get(ctx, nil)
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
			gs,
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
