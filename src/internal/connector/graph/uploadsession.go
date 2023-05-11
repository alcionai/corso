package graph

import (
	"bytes"
	"context"
	"fmt"

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
type largeItemWriter struct {
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

func NewLargeItemWriter(id, url string, size int64) *largeItemWriter {
	return &largeItemWriter{id: id, url: url, contentLength: size, client: *NewNoTimeoutHTTPWrapper()}
}

// Write will upload the provided data to M365. It sets the `Content-Length` and `Content-Range` headers based on
// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession
func (iw *largeItemWriter) Write(p []byte) (int, error) {
	rangeLength := len(p)
	ctx := context.Background()

	logger.Ctx(ctx).
		Debugf("WRITE for %s. Size:%d, Offset: %d, TotalSize: %d",
			iw.id, rangeLength, iw.lastWrittenOffset, iw.contentLength)

	endOffset := iw.lastWrittenOffset + int64(rangeLength)

	// PUT the request - set headers `Content-Range`to describe total size and `Content-Length` to describe size of
	// data in the current request
	headers := make(map[string]string)
	headers[contentRangeHeaderKey] = fmt.Sprintf(
		contentRangeHeaderValueFmt,
		iw.lastWrittenOffset,
		endOffset-1,
		iw.contentLength)
	headers[contentLengthHeaderKey] = fmt.Sprintf("%d", rangeLength)

	_, err := iw.client.Request(ctx, "PUT", iw.url, bytes.NewReader(p), headers)
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
