package onedrive

import (
	"context"
	"io"
	"time"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	msup "github.com/microsoftgraph/msgraph-sdk-go/drives/item/items/item/createuploadsession"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/connector/uploadsession"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	// downloadUrlKey is used to find the download URL in a
	// DriveItem response
	downloadURLKey = "@microsoft.graph.downloadUrl"
)

// sharePointItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func sharePointItemReader(
	ctx context.Context,
	service graph.Service,
	driveID, itemID string,
) (details.ItemInfo, io.ReadCloser, error) {
	item, rc, err := driveItemReader(ctx, service, driveID, itemID)
	if err != nil {
		return details.ItemInfo{}, nil, err
	}

	dii := details.ItemInfo{
		SharePoint: sharePointItemInfo(item, *item.GetSize()),
	}

	return dii, rc, nil
}

// oneDriveItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func oneDriveItemReader(
	ctx context.Context,
	service graph.Service,
	driveID, itemID string,
) (details.ItemInfo, io.ReadCloser, error) {
	item, rc, err := driveItemReader(ctx, service, driveID, itemID)
	if err != nil {
		return details.ItemInfo{}, nil, err
	}

	dii := details.ItemInfo{
		OneDrive: oneDriveItemInfo(item, *item.GetSize()),
	}

	return dii, rc, nil
}

// driveItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func driveItemReader(
	ctx context.Context,
	service graph.Service,
	driveID, itemID string,
) (models.DriveItemable, io.ReadCloser, error) {
	logger.Ctx(ctx).Debugw("Reading Item", "id", itemID, "time", time.Now())

	item, err := service.Client().DrivesById(driveID).ItemsById(itemID).Get(ctx, nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to get item %s", itemID)
	}

	// Get the download URL - https://docs.microsoft.com/en-us/graph/api/driveitem-get-content
	// These URLs are pre-authenticated and can be used to download the data using the standard
	// http client
	if _, found := item.GetAdditionalData()[downloadURLKey]; !found {
		return nil, nil, errors.Errorf("file does not have a download URL. ID: %s, %#v",
			itemID, item.GetAdditionalData())
	}

	downloadURL := item.GetAdditionalData()[downloadURLKey].(*string)

	clientOptions := msgraphsdk.GetDefaultClientOptions()
	middlewares := msgraphgocore.GetDefaultMiddlewaresWithOptions(&clientOptions)

	httpClient := msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
	httpClient.Timeout = 0 // need infinite timeout for pulling large files

	resp, err := httpClient.Get(*downloadURL)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to download file from %s", *downloadURL)
	}

	return item, resp.Body, nil
}

// oneDriveItemInfo will populate a details.OneDriveInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
func oneDriveItemInfo(di models.DriveItemable, itemSize int64) *details.OneDriveInfo {
	ed, ok := di.GetCreatedBy().GetUser().GetAdditionalData()["email"]

	email := ""
	if ok {
		email = *ed.(*string)
	}

	return &details.OneDriveInfo{
		ItemType: details.OneDriveItem,
		ItemName: *di.GetName(),
		Created:  *di.GetCreatedDateTime(),
		Modified: *di.GetLastModifiedDateTime(),
		Size:     itemSize,
		Owner:    email,
	}
}

// sharePointItemInfo will populate a details.SharePointInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
func sharePointItemInfo(di models.DriveItemable, itemSize int64) *details.SharePointInfo {
	var (
		id   string
		url  string
		idp  = di.GetSharepointIds().GetSiteId()
		urlp = di.GetSharepointIds().GetSiteUrl()
	)

	if idp != nil {
		id = *idp
	}

	if urlp != nil {
		url = *urlp
	}

	return &details.SharePointInfo{
		ItemType: details.OneDriveItem,
		ItemName: *di.GetName(),
		Created:  *di.GetCreatedDateTime(),
		Modified: *di.GetLastModifiedDateTime(),
		Size:     itemSize,
		Owner:    id,
		WebURL:   url,
	}
}

// driveItemWriter is used to initialize and return an io.Writer to upload data for the specified item
// It does so by creating an upload session and using that URL to initialize an `itemWriter`
func driveItemWriter(
	ctx context.Context,
	service graph.Service,
	driveID, itemID string,
	itemSize int64,
) (io.Writer, error) {
	// TODO: @vkamra verify if var session is the desired input
	session := msup.NewCreateUploadSessionPostRequestBody()

	r, err := service.Client().DrivesById(driveID).ItemsById(itemID).CreateUploadSession().Post(ctx, session, nil)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to create upload session for item %s. details: %s",
			itemID,
			support.ConnectorStackErrorTrace(err),
		)
	}

	url := *r.GetUploadUrl()

	logger.Ctx(ctx).Debugf("Created an upload session for item %s. URL: %s", itemID, url)

	return uploadsession.NewWriter(itemID, url, itemSize), nil
}
