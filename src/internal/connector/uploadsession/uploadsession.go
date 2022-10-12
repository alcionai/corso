package uploadsession

import (
	"bytes"
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/resty.v1"

	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	contentRangeHeaderKey = "Content-Range"
	// Format for Content-Range is "bytes <start>-<end>/<total>"
	contentRangeHeaderValueFmt = "bytes %d-%d/%d"
	contentLengthHeaderKey     = "Content-Length"
)

// Writer implements an io.Writer for a M365
// UploadSession URL
type writer struct {
	// Identifier
	id string
	// Upload URL for this item
	url string
	// Tracks how much data will be written
	contentLength int64
	// Last item offset that was written to
	lastWrittenOffset int64
}

func NewWriter(id, url string, size int64) *writer {
	return &writer{id: id, url: url, contentLength: size}
}

// Write will upload the provided data to M365. It sets the `Content-Length` and `Content-Range` headers based on
// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession
func (iw *writer) Write(p []byte) (n int, err error) {
	rangeLength := len(p)
	logger.Ctx(context.Background()).Debugf("WRITE for %s. Size:%d, Offset: %d, TotalSize: %d",
		iw.id, rangeLength, iw.lastWrittenOffset, iw.contentLength)

	endOffset := iw.lastWrittenOffset + int64(rangeLength)

	client := resty.New()

	// PUT the request - set headers `Content-Range`to describe total size and `Content-Length` to describe size of
	// data in the current request
	resp, err := client.R().
		SetHeaders(map[string]string{
			contentRangeHeaderKey: fmt.Sprintf(contentRangeHeaderValueFmt,
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
