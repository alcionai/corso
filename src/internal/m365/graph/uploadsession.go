package graph

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

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
	// ID is the id of the item created.
	// Will be available after the upload is complete
	ID string
	// Identifier
	parentID string
	// Upload URL for this item
	url string
	// Tracks how much data will be written
	contentLength int64
	// Last item offset that was written to
	lastWrittenOffset int64
	client            httpWrapper
}

func NewLargeItemWriter(parentID, url string, size int64) *largeItemWriter {
	return &largeItemWriter{
		parentID:      parentID,
		url:           url,
		contentLength: size,
		client:        *NewNoTimeoutHTTPWrapper(),
	}
}

// Write will upload the provided data to M365. It sets the `Content-Length` and `Content-Range` headers based on
// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession
func (iw *largeItemWriter) Write(p []byte) (int, error) {
	rangeLength := len(p)
	ctx := context.Background()

	logger.Ctx(ctx).
		Debugf("WRITE for %s. Size:%d, Offset: %d, TotalSize: %d",
			iw.parentID, rangeLength, iw.lastWrittenOffset, iw.contentLength)

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

	resp, err := iw.client.Request(
		ctx,
		http.MethodPut,
		iw.url,
		bytes.NewReader(p),
		headers)
	if err != nil {
		return 0, clues.Wrap(err, "uploading item").With(
			"upload_id", iw.parentID,
			"upload_chunk_size", rangeLength,
			"upload_offset", iw.lastWrittenOffset,
			"upload_size", iw.contentLength)
	}

	// Update last offset
	iw.lastWrittenOffset = endOffset

	// Once the upload is complete, we get a Location header in the
	// below format from which we can get the id of the uploaded
	// item. This will only be available after we have uploaded the
	// entire content(based on the size in the req header).
	// https://outlook.office.com/api/v2.0/Users('<user-id>')/Messages('<message-id>')/Attachments('<attachment-id>')
	// Ref: https://learn.microsoft.com/en-us/graph/outlook-large-attachments?tabs=http
	loc := resp.Header.Get("Location")
	if loc != "" {
		splits := strings.Split(loc, "'")
		if len(splits) != 7 || splits[4] != ")/Attachments(" || len(splits[5]) == 0 {
			return 0, clues.New("invalid format for upload completion url").
				With("location", loc)
		}

		iw.ID = splits[5]
	}

	return rangeLength, nil
}
