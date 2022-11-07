package exchange

import (
	"bytes"
	"context"
	"io"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/uploadsession"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	// Use large attachment logic for attachments > 3MB
	// https://learn.microsoft.com/en-us/graph/outlook-large-attachments
	largeAttachmentSize           = int32(3 * 1024 * 1024)
	attachmentChunkSize           = 4 * 1024 * 1024
	fileAttachmentOdataValue      = "#microsoft.graph.fileAttachment"
	itemAttachmentOdataValue      = "#microsoft.graph.itemAttachment"
	referenceAttachmentOdataValue = "#microsoft.graph.referenceAttachment"
)

func attachmentType(attachment models.Attachmentable) models.AttachmentType {
	switch *attachment.GetOdataType() {
	case fileAttachmentOdataValue:
		return models.FILE_ATTACHMENTTYPE
	case itemAttachmentOdataValue:
		return models.ITEM_ATTACHMENTTYPE
	case referenceAttachmentOdataValue:
		return models.REFERENCE_ATTACHMENTTYPE
	default:
		// Should not hit this but default to ITEM_ATTACHMENTTYPE
		// which will pick the default attachment upload mechanism
		return models.ITEM_ATTACHMENTTYPE
	}
}

// uploadAttachment will upload the specified message attachment to M365
func uploadAttachment(
	ctx context.Context,
	uploader attachmentUploadable,
	attachment models.Attachmentable,
) error {
	logger.Ctx(ctx).Debugf("uploading attachment with size %d", *attachment.GetSize())

	// For Item/Reference attachments *or* file attachments < 3MB, use the attachments endpoint
	if attachmentType(attachment) != models.FILE_ATTACHMENTTYPE || *attachment.GetSize() < largeAttachmentSize {
		err := uploader.uploadSmallAttachment(ctx, attachment)

		return err
	}

	return uploadLargeAttachment(ctx, uploader, attachment)
}

// uploadLargeAttachment will upload the specified attachment by creating an upload session and
// doing a chunked upload
func uploadLargeAttachment(ctx context.Context, uploader attachmentUploadable,
	attachment models.Attachmentable,
) error {
	ab := attachmentBytes(attachment)
	size := int64(len(ab))

	session, err := uploader.uploadSession(ctx, *attachment.GetName(), size)
	if err != nil {
		return err
	}

	url := *session.GetUploadUrl()
	aw := uploadsession.NewWriter(uploader.getItemID(), url, size)
	logger.Ctx(ctx).Debugf("Created an upload session for item %s. URL: %s", uploader.getItemID(), url)

	// Upload the stream data
	copyBuffer := make([]byte, attachmentChunkSize)

	_, err = io.CopyBuffer(aw, bytes.NewReader(ab), copyBuffer)
	if err != nil {
		return errors.Wrapf(err, "failed to upload attachment: item %s", uploader.getItemID())
	}

	return nil
}
