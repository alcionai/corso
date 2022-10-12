package onedrive

import (
	"context"
	"io"
	"net/http"
	"time"

	msup "github.com/microsoftgraph/msgraph-sdk-go/drives/item/items/item/createuploadsession"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
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

// itemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func driveItemReader(
	ctx context.Context,
	service graph.Service,
	driveID, itemID string,
) (*details.OneDriveInfo, io.ReadCloser, error) {
	logger.Ctx(ctx).Debugf("Reading Item %s at %s", itemID, time.Now())

	item, err := service.Client().DrivesById(driveID).ItemsById(itemID).Get(ctx, nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to get item %s", itemID)
	}

	logger.Ctx(ctx).Debugw("reading item", "name", *item.GetName(), "time", common.Now())

	// Get the download URL - https://docs.microsoft.com/en-us/graph/api/driveitem-get-content
	// These URLs are pre-authenticated and can be used to download the data using the standard
	// http client
	if _, found := item.GetAdditionalData()[downloadURLKey]; !found {
		return nil, nil, errors.Errorf("file does not have a download URL. ID: %s, %#v",
			itemID, item.GetAdditionalData())
	}

	downloadURL := item.GetAdditionalData()[downloadURLKey].(*string)

	// TODO: We should use the `msgraphgocore` http client which has the right
	// middleware/options configured
	resp, err := http.Get(*downloadURL)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to download file from %s", *downloadURL)
	}

	return driveItemInfo(item), resp.Body, nil
}

// driveItemInfo will populate a details.OneDriveInfo struct
// with properties from the drive item.
func driveItemInfo(di models.DriveItemable) *details.OneDriveInfo {
	return &details.OneDriveInfo{
		ItemType: details.OneDriveItem,
		ItemName: *di.GetName(),
		Created:  *di.GetCreatedDateTime(),
		Modified: *di.GetLastModifiedDateTime(),
		Size:     *di.GetSize(),
	}
}

// driveItemWriter is used to initialize and return an io.Writer to upload data for the specified item
// It does so by creating an upload session and using that URL to initialize an `itemWriter`
func driveItemWriter(ctx context.Context, service graph.Service, driveID, itemID string, itemSize int64,
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
