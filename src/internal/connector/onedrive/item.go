package onedrive

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	msup "github.com/microsoftgraph/msgraph-sdk-go/drives/item/items/item/createuploadsession"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
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

	return &details.OneDriveInfo{
		ItemType:     details.OneDriveItem,
		ItemName:     *item.GetName(),
		Created:      *item.GetCreatedDateTime(),
		LastModified: *item.GetLastModifiedDateTime(),
		Size:         *item.GetSize(),
	}, resp.Body, nil
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

	return &itemWriter{id: itemID, contentLength: itemSize, url: url}, nil
}

// itemWriter implements an io.Writer for the OneDrive URL
// it is initialized with
type itemWriter struct {
	// Item ID
	id string
	// Upload URL for this item
	url string
	// Tracks how much data will be written
	contentLength int64
	// Last item offset that was written to
	lastWrittenOffset int64
}

const (
	contentRangeHeaderKey  = "Content-Range"
	contentLengthHeaderKey = "Content-Length"
	// Format for Content-Length is "bytes <start>-<end>/<total>"
	contentLengthHeaderValueFmt = "bytes %d-%d/%d"
)

// Write will upload the provided data to OneDrive. It sets the `Content-Length` and `Content-Range` headers based on
// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession
func (iw *itemWriter) Write(p []byte) (n int, err error) {
	rangeLength := len(p)
	logger.Ctx(context.Background()).Debugf("WRITE for %s. Size:%d, Offset: %d, TotalSize: %d",
		iw.id, rangeLength, iw.lastWrittenOffset, iw.contentLength)

	endOffset := iw.lastWrittenOffset + int64(rangeLength)

	client := resty.New()

	// PUT the request - set headers `Content-Range`to describe total size and `Content-Length` to describe size of
	// data in the current request
	resp, err := client.R().
		SetHeaders(map[string]string{
			contentRangeHeaderKey: fmt.Sprintf(contentLengthHeaderValueFmt,
				iw.lastWrittenOffset,
				endOffset-1,
				iw.contentLength),
			contentLengthHeaderKey: fmt.Sprintf("%d", iw.contentLength),
		}).
		SetBody(bytes.NewReader(p)).Put(iw.url)
	if err != nil {
		return 0, errors.Wrapf(err,
			"failed to upload item %s. Upload failed at Size:%d, Offset: %d, TotalSize: %d ",
			iw.id, rangeLength, iw.lastWrittenOffset, iw.contentLength)
	}

	// Update last offset
	iw.lastWrittenOffset = endOffset

	logger.Ctx(context.Background()).Debugf("Response: %s", resp.String())

	return rangeLength, nil
}
