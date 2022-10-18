package onedrive

import (
	"context"
	"fmt"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/items"
	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/items/item"
	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/delta"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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
		"cdd28e44-67e3-425e-be4c-737fab2899d3",
		"b214fe43-f5a3-4703-beeb-fa97188220fc",
		"c2273bd0-dff7-4215-9ef5-2c7bcfb06425",
		"12b8c807-2e20-48fc-b453-542b6ee9d171",
		"c32f9321-a627-406d-a114-1f9c81aaafac",
		"e6778190-713e-4e4f-9119-8b8238de25df",
		"ed01faf2-1d88-4947-ae91-45ca18703a96",
		"ca7f3140-d88c-455b-9a1c-7f0679e31a76",
		"38b434d2-a15e-4cde-9a98-e737c75623e1",
		"4b244418-9658-4451-a2b8-b5e2b364e9bd",
		"c5928f49-12ba-48f7-ada3-0d743a3601d5",
		"4ae99959-6b0f-43b0-b1ce-68146001bdba",
		"afcafa6a-d966-4462-918c-ec0b4e0fe642",
		"c42b9cae-ea4f-4ab7-9717-81576235ccac", // Microsoft 365 E5 Developer
	}
)

const (
	// nextLinkKey is used to find the next link in a paged
	// graph response
	nextLinkKey           = "@odata.nextLink"
	itemChildrenRawURLFmt = "https://graph.microsoft.com/v1.0/drives/%s/items/%s/children"
	itemByPathRawURLFmt   = "https://graph.microsoft.com/v1.0/drives/%s/items/%s:/%s"
	itemNotFoundErrorCode = "itemNotFound"
)

// Enumerates the drives for the specified user
func drives(ctx context.Context, service graph.Service, user string) ([]models.Driveable, error) {
	var hasDrive bool

	resp, err := service.Client().UsersById(user).LicenseDetails().Get(ctx, nil)
	if err != nil {
		return nil, errors.New("First call fails too")
	}

	licenses := resp.GetValue()
	for _, entry := range licenses {
		sku := entry.GetSkuId()
		if sku == nil {
			continue
		}

		if ok := hasLicense(*sku); ok {
			hasDrive = true
			break
		}
	}

	if !hasDrive {
		return make([]models.Driveable, 0), nil // no license
	}

	r, err := service.Client().UsersById(user).Drives().Get(ctx, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve user drives. user: %s, details: %s",
			user, support.ConnectorStackErrorTrace(err))
	}

	logger.Ctx(ctx).Debugf("Found %d drives for user %s", len(r.GetValue()), user)

	return r.GetValue(), nil
}

// itemCollector functions collect the items found in a drive
type itemCollector func(ctx context.Context, driveID string, driveItems []models.DriveItemable) error

// collectItems will enumerate all items in the specified drive and hand them to the
// provided `collector` method
func collectItems(
	ctx context.Context,
	service graph.Service,
	driveID string,
	collector itemCollector,
) error {
	// TODO: Specify a timestamp in the delta query
	// https://docs.microsoft.com/en-us/graph/api/driveitem-delta?
	// view=graph-rest-1.0&tabs=http#example-4-retrieving-delta-results-using-a-timestamp
	builder := service.Client().DrivesById(driveID).Root().Delta()

	for {
		r, err := builder.Get(ctx, nil)
		if err != nil {
			return errors.Wrapf(
				err,
				"failed to query drive items. details: %s",
				support.ConnectorStackErrorTrace(err),
			)
		}

		err = collector(ctx, driveID, r.GetValue())
		if err != nil {
			return err
		}

		// Check if there are more items
		if _, found := r.GetAdditionalData()[nextLinkKey]; !found {
			break
		}

		nextLink := r.GetAdditionalData()[nextLinkKey].(*string)
		logger.Ctx(ctx).Debugf("Found %s nextLink", *nextLink)
		builder = delta.NewDeltaRequestBuilder(*nextLink, service.Adapter())
	}

	return nil
}

// getFolder will lookup the specified folder name under `parentFolderID`
func getFolder(
	ctx context.Context,
	service graph.Service,
	driveID, parentFolderID, folderName string,
) (models.DriveItemable, error) {
	// The `Children().Get()` API doesn't yet support $filter, so using that to find a folder
	// will be sub-optimal.
	// Instead, we leverage OneDrive path-based addressing -
	// https://learn.microsoft.com/en-us/graph/onedrive-addressing-driveitems#path-based-addressing
	// - which allows us to lookup an item by its path relative to the parent ID
	rawURL := fmt.Sprintf(itemByPathRawURLFmt, driveID, parentFolderID, folderName)

	builder := item.NewDriveItemItemRequestBuilder(rawURL, service.Adapter())

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
	service graph.Service,
	driveID, parentFolderID string,
	newItem models.DriveItemable,
) (models.DriveItemable, error) {
	// Graph SDK doesn't yet provide a POST method for `/children` so we set the `rawUrl` ourselves as recommended
	// here: https://github.com/microsoftgraph/msgraph-sdk-go/issues/155#issuecomment-1136254310
	rawURL := fmt.Sprintf(itemChildrenRawURLFmt, driveID, parentFolderID)

	builder := items.NewItemsRequestBuilder(rawURL, service.Adapter())

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
	gs graph.Service,
	userID string,
	prefix string,
) ([]*Displayable, error) {
	drives, err := drives(ctx, gs, userID)
	if err != nil {
		return nil, errors.Wrap(err, "getting OneDrive folders")
	}

	res := []*Displayable{}

	for _, d := range drives {
		err = collectItems(
			ctx,
			gs,
			*d.GetId(),
			func(innerCtx context.Context, driveID string, items []models.DriveItemable) error {
				for _, item := range items {
					// Skip the root item.
					if item.GetRoot() != nil {
						continue
					}

					// Only selecting folders right now, not packages.
					if item.GetFolder() == nil {
						continue
					}

					if !strings.HasPrefix(*item.GetName(), prefix) {
						continue
					}

					// Add the item instead of the folder because the item has more
					// functionality.
					res = append(res, &Displayable{item})
				}

				return nil
			},
		)
		if err != nil {
			return nil, errors.Wrapf(err, "getting items for drive %s", *d.GetName())
		}
	}

	return res, nil
}

func DeleteItem(
	ctx context.Context,
	gs graph.Service,
	driveID string,
	itemID string,
) error {
	err := gs.Client().DrivesById(driveID).ItemsById(itemID).Delete(ctx, nil)
	if err != nil {
		return errors.Wrapf(err, "deleting item with ID %s", itemID)
	}

	return nil
}

func hasLicense(skuID string) bool {
	for _, license := range skuIDs {
		if skuID == license {
			return true
		}
	}

	return false
}
