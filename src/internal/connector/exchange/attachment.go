package exchange

import (
	"bytes"
	"context"
	"io"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/support"
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

	var (
		attachmentType = attachmentType(attachment)
		err            error
	)
	// Reference attachments that are inline() do not need to be recreated. The contents are part of the body.
	if attachmentType == models.REFERENCE_ATTACHMENTTYPE &&
		attachment.GetIsInline() != nil && *attachment.GetIsInline() {
		logger.Ctx(ctx).Debugf("skip uploading inline reference attachment: ", *attachment.GetName())
		return nil
	}

	// item Attachments to be skipped until the completion of Issue #2353
	if attachmentType == models.ITEM_ATTACHMENTTYPE {
		prev := attachment

		attachment, err = support.ToItemAttachment(attachment)
		if err != nil {
			name := ""
			if prev.GetName() != nil {
				name = *prev.GetName()
			}

			// TODO: (rkeepers) Update to support PII protection
			msg := "item attachment restore not supported for this type. skipping upload."
			logger.Ctx(ctx).Infow(msg,
				"err", err,
				"attachment_name", name,
				"attachment_type", attachmentType,
				"internal_item_type", getItemAttachmentItemType(prev),
				"attachment_id", *prev.GetId(),
			)

			return nil
		}
	}

	// For Item/Reference attachments *or* file attachments < 3MB, use the attachments endpoint
	if attachmentType != models.FILE_ATTACHMENTTYPE || *attachment.GetSize() < largeAttachmentSize {
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

func getItemAttachmentItemType(query models.Attachmentable) string {
	empty := ""
	attachment, ok := query.(models.ItemAttachmentable)

	if !ok {
		return empty
	}

	item := attachment.GetItem()
	if item.GetOdataType() == nil {
		return empty
	}

	return *item.GetOdataType()
}
