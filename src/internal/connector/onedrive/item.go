package onedrive

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/pkg/logger"
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
) (string, io.ReadCloser, error) {
	logger.Ctx(ctx).Debugf("Reading Item %s at %s", itemID, time.Now())

	item, err := service.Client().DrivesById(driveID).ItemsById(itemID).Get()
	if err != nil {
		return "", nil, errors.Wrapf(err, "failed to get item %s", itemID)
	}

	// Get the download URL - https://docs.microsoft.com/en-us/graph/api/driveitem-get-content
	// These URLs are pre-authenticated and can be used to download the data using the standard
	// http client
	if _, found := item.GetAdditionalData()[downloadURLKey]; !found {
		return "", nil, errors.Errorf("file does not have a download URL. ID: %s, %#v",
			itemID, item.GetAdditionalData())
	}
	downloadURL := item.GetAdditionalData()[downloadURLKey].(*string)

	// TODO: We should use the `msgraphgocore` http client which has the right
	// middleware/options configured
	resp, err := http.Get(*downloadURL)
	if err != nil {
		return "", nil, errors.Wrapf(err, "failed to download file from %s", *downloadURL)
	}

	return *item.GetName(), resp.Body, nil
}
