package uploadsession

import (
	"bytes"
	"context"
	"fmt"

	"github.com/alcionai/clues"
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
	client            *resty.Client
}

func NewWriter(id, url string, size int64) *writer {
	return &writer{id: id, url: url, contentLength: size, client: resty.New()}
}

// Write will upload the provided data to M365. It sets the `Content-Length` and `Content-Range` headers based on
// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession
func (iw *writer) Write(p []byte) (int, error) {
	rangeLength := len(p)
	logger.Ctx(context.Background()).Debugf("WRITE for %s. Size:%d, Offset: %d, TotalSize: %d",
		iw.id, rangeLength, iw.lastWrittenOffset, iw.contentLength)

	endOffset := iw.lastWrittenOffset + int64(rangeLength)

	// PUT the request - set headers `Content-Range`to describe total size and `Content-Length` to describe size of
	// data in the current request
	_, err := iw.client.R().
		SetHeaders(map[string]string{
			contentRangeHeaderKey: fmt.Sprintf(contentRangeHeaderValueFmt,
				iw.lastWrittenOffset,
				endOffset-1,
				iw.contentLength),
			contentLengthHeaderKey: fmt.Sprintf("%d", rangeLength),
		}).
		SetBody(bytes.NewReader(p)).Put(iw.url)
	if err != nil {
		return 0, clues.Wrap(err, "uploading item").WithAll(
			"upload_id", iw.id,
			"upload_chunk_size", rangeLength,
			"upload_offset", iw.lastWrittenOffset,
			"upload_size", iw.contentLength)
	}

	// Update last offset
	iw.lastWrittenOffset = endOffset

	return rangeLength, nil
}
