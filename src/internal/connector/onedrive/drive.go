package onedrive

import (
	"context"
	"fmt"
	"strings"
	"time"

	msdrive "github.com/microsoftgraph/msgraph-beta-sdk-go/drive"
	msdrives "github.com/microsoftgraph/msgraph-beta-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/sites"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/logger"
)

var (
	errFolderNotFound = errors.New("folder not found")

	// nolint:lll
	// OneDrive associated SKUs located at:
	// https://learn.microsoft.com/en-us/azure/active-directory/enterprise-users/licensing-service-plan-reference
	skuIDs = []string{
		// Microsoft 365 Apps for Business 0365
		"cdd28e44-67e3-425e-be4c-737fab2899d3",
		// Microsoft 365 Apps for Business SMB_Business
		"b214fe43-f5a3-4703-beeb-fa97188220fc",
		// Microsoft 365 Apps for enterprise
		"c2273bd0-dff7-4215-9ef5-2c7bcfb06425",
		// Microsoft 365 Apps for Faculty
		"12b8c807-2e20-48fc-b453-542b6ee9d171",
		// Microsoft 365 Apps for Students
		"c32f9321-a627-406d-a114-1f9c81aaafac",
		// OneDrive for Business (Plan 1)
		"e6778190-713e-4e4f-9119-8b8238de25df",
		// OneDrive for Business (Plan 2)
		"ed01faf2-1d88-4947-ae91-45ca18703a96",
		// Visio Plan 1
		"ca7f3140-d88c-455b-9a1c-7f0679e31a76",
		// Visio Plan 2
		"38b434d2-a15e-4cde-9a98-e737c75623e1",
		// Visio Online Plan 1
		"4b244418-9658-4451-a2b8-b5e2b364e9bd",
		// Visio Online Plan 2
		"c5928f49-12ba-48f7-ada3-0d743a3601d5",
		// Visio Plan 2 for GCC
		"4ae99959-6b0f-43b0-b1ce-68146001bdba",
		// ONEDRIVEENTERPRISE
		"afcafa6a-d966-4462-918c-ec0b4e0fe642",
		// Microsoft 365 E5 Developer
		"c42b9cae-ea4f-4ab7-9717-81576235ccac",
		// Microsoft 365 E5
		"06ebc4ee-1bb5-47dd-8120-11324bc54e06",
		// Office 365 E4
		"1392051d-0cb9-4b7a-88d5-621fee5e8711",
		// Microsoft 365 E3
		"05e9a617-0261-4cee-bb44-138d3ef5d965",
		// Microsoft 365 Business Premium
		"cbdc14ab-d96c-4c30-b9f4-6ada7cdc1d46",
		// Microsoft 365 Business Standard
		"f245ecc8-75af-4f8e-b61f-27d8114de5f3",
		// Microsoft 365 Business Basic
		"3b555118-da6a-4418-894f-7df1e2096870",
	}
)

const (
	// nextLinkKey is used to find the next link in a paged
	// graph response
	nextLinkKey           = "@odata.nextLink"
	itemChildrenRawURLFmt = "https://graph.microsoft.com/v1.0/drives/%s/items/%s/children"
	itemByPathRawURLFmt   = "https://graph.microsoft.com/v1.0/drives/%s/items/%s:/%s"
	itemNotFoundErrorCode = "itemNotFound"
	userMysiteURLNotFound = "BadRequest Unable to retrieve user's mysite URL"
	userMysiteNotFound    = "ResourceNotFound User's mysite not found"
)

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
	options := &sites.ItemDrivesRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemDrivesRequestBuilderGetQueryParameters{
			Select: []string{"id", "name", "weburl", "system"},
		},
	}

	r, err := service.Client().SitesById(site).Drives().Get(ctx, options)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve site drives. site: %s, details: %s",
			site, support.ConnectorStackErrorTrace(err))
	}

	return r.GetValue(), nil
}

func userDrives(ctx context.Context, service graph.Servicer, user string) ([]models.Driveable, error) {
	var (
		hasDrive        bool
		numberOfRetries = 3
		r               models.DriveCollectionResponseable
	)

	hasDrive, err := hasDriveLicense(ctx, service, user)
	if err != nil {
		return nil, errors.Wrap(err, user)
	}

	if !hasDrive {
		logger.Ctx(ctx).Debugf("User %s does not have a license for OneDrive", user)
		return make([]models.Driveable, 0), nil // no license
	}

	// Retry Loop for Drive retrieval. Request can timeout
	for i := 0; i <= numberOfRetries; i++ {
		r, err = service.Client().UsersById(user).Drives().Get(ctx, nil)
		if err != nil {
			detailedError := support.ConnectorStackErrorTrace(err)
			if strings.Contains(detailedError, userMysiteURLNotFound) ||
				strings.Contains(detailedError, userMysiteNotFound) {
				logger.Ctx(ctx).Debugf("User %s does not have a drive", user)
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

		break
	}

	logger.Ctx(ctx).Debugf("Found %d drives for user %s", len(r.GetValue()), user)

	return r.GetValue(), nil
}

// itemCollector functions collect the items found in a drive
type itemCollector func(
	ctx context.Context,
	driveID, driveName string,
	driveItems []models.DriveItemable,
	paths map[string]string,
) error

// collectItems will enumerate all items in the specified drive and hand them to the
// provided `collector` method
func collectItems(
	ctx context.Context,
	service graph.Servicer,
	driveID, driveName string,
	collector itemCollector,
) (string, map[string]string, error) {
	var (
		newDeltaURL = ""
		// TODO(ashmrtn): Eventually this should probably be a parameter so we can
		// take in previous paths.
		paths = map[string]string{}
	)

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
			return "", nil, errors.Wrapf(
				err,
				"failed to query drive items. details: %s",
				support.ConnectorStackErrorTrace(err),
			)
		}

		err = collector(ctx, driveID, driveName, r.GetValue(), paths)
		if err != nil {
			return "", nil, err
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

	return newDeltaURL, paths, nil
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
	userID string,
	prefix string,
) ([]*Displayable, error) {
	drives, err := drives(ctx, gs, userID, OneDriveSource)
	if err != nil {
		return nil, errors.Wrap(err, "getting OneDrive folders")
	}

	folders := map[string]*Displayable{}

	for _, d := range drives {
		_, _, err = collectItems(
			ctx,
			gs,
			*d.GetId(),
			*d.GetName(),
			func(
				innerCtx context.Context,
				driveID, driveName string,
				items []models.DriveItemable,
				paths map[string]string,
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

// hasDriveLicense utility function that queries M365 server
// to investigate the user's includes access to OneDrive.
func hasDriveLicense(
	ctx context.Context,
	service graph.Servicer,
	user string,
) (bool, error) {
	var hasDrive bool

	resp, err := service.Client().UsersById(user).LicenseDetails().Get(ctx, nil)
	if err != nil {
		return false,
			errors.Wrap(err, "failure obtaining license details for user")
	}

	iter, err := msgraphgocore.NewPageIterator(
		resp, service.Adapter(),
		models.CreateLicenseDetailsCollectionResponseFromDiscriminatorValue,
	)
	if err != nil {
		return false, err
	}

	cb := func(pageItem any) bool {
		entry, ok := pageItem.(models.LicenseDetailsable)
		if !ok {
			err = errors.New("casting item to models.LicenseDetailsable")
			return false
		}

		sku := entry.GetSkuId()
		if sku == nil {
			return true
		}

		for _, license := range skuIDs {
			if sku.String() == license {
				hasDrive = true
				return false
			}
		}

		return true
	}

	if err := iter.Iterate(ctx, cb); err != nil {
		return false,
			errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	return hasDrive, nil
}
