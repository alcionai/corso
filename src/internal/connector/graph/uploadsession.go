package graph

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/alcionai/clues"

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
	client            httpWrapper
}

func NewWriter(id, url string, size int64) *writer {
	return &writer{id: id, url: url, contentLength: size, client: *NewNoTimeoutHTTPWrapper()}
}

// Write will upload the provided data to M365. It sets the `Content-Length` and `Content-Range` headers based on
// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession
func (iw *writer) Write(p []byte) (int, error) {
	rangeLength := len(p)
	logger.Ctx(context.Background()).
		Debugf("WRITE for %s. Size:%d, Offset: %d, TotalSize: %d",
			iw.id, rangeLength, iw.lastWrittenOffset, iw.contentLength)

	endOffset := iw.lastWrittenOffset + int64(rangeLength)

	// PUT the request - set headers `Content-Range`to describe total size and `Content-Length` to describe size of
	// data in the current request
	req, err := http.NewRequest("PUT", iw.url, bytes.NewReader(p))
	if err != nil {
		return 0, clues.Wrap(err, "uploading item").With(
			"upload_id", iw.id,
			"upload_chunk_size", rangeLength,
			"upload_offset", iw.lastWrittenOffset,
			"upload_size", iw.contentLength)
	}

	req.Header = http.Header{
		contentRangeHeaderKey: {fmt.Sprintf(
			contentRangeHeaderValueFmt,
			iw.lastWrittenOffset,
			endOffset-1,
			iw.contentLength)},
		contentLengthHeaderKey: {fmt.Sprintf("%d", rangeLength)},
	}
	_, err = iw.client.client.Do(req)

	if err != nil {
		return 0, clues.Wrap(err, "uploading item").With(
			"upload_id", iw.id,
			"upload_chunk_size", rangeLength,
			"upload_offset", iw.lastWrittenOffset,
			"upload_size", iw.contentLength)
	}

	// Update last offset
	iw.lastWrittenOffset = endOffset

	return rangeLength, nil
}
