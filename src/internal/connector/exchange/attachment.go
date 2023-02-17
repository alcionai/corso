package exchange

import (
	"bytes"
	"context"
	"io"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
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
	attachmentType := ptr.Val(attachment.GetOdataType())
	switch attachmentType {
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
	attachmentType := attachmentType(attachment)

	ctx = clues.AddAll(
		ctx,
		"attachment_size", ptr.Val(attachment.GetSize()),
		"attachment_id", ptr.Val(attachment.GetId()),
		"attachment_name", ptr.Val(attachment.GetName()), // TODO: pii
		"attachment_type", attachmentType,
		"internal_item_type", getItemAttachmentItemType(attachment),
		"uploader_item_id", uploader.getItemID())

	logger.Ctx(ctx).Debug("uploading attachment")

	// Reference attachments that are inline() do not need to be recreated. The contents are part of the body.
	if attachmentType == models.REFERENCE_ATTACHMENTTYPE && ptr.Val(attachment.GetIsInline()) {
		logger.Ctx(ctx).Debug("skip uploading inline reference attachment")
		return nil
	}

	// item Attachments to be skipped until the completion of Issue #2353
	if attachmentType == models.ITEM_ATTACHMENTTYPE {
		a, err := support.ToItemAttachment(attachment)
		if err != nil {
			logger.Ctx(ctx).
				With("err", err).
				Infow("item attachment restore not supported for this type. skipping upload.", clues.InErr(err).Slice()...)

			return nil
		}

		attachment = a
	}

	// For Item/Reference attachments *or* file attachments < 3MB, use the attachments endpoint
	if attachmentType != models.FILE_ATTACHMENTTYPE || ptr.Val(attachment.GetSize()) < largeAttachmentSize {
		return uploader.uploadSmallAttachment(ctx, attachment)
	}

	return uploadLargeAttachment(ctx, uploader, attachment)
}

// uploadLargeAttachment will upload the specified attachment by creating an upload session and
// doing a chunked upload
func uploadLargeAttachment(
	ctx context.Context,
	uploader attachmentUploadable,
	attachment models.Attachmentable,
) error {
	var (
		bs   = attachmentBytes(attachment)
		size = int64(len(bs))
	)

	session, err := uploader.uploadSession(ctx, ptr.Val(attachment.GetName()), size)
	if err != nil {
		return clues.Stack(err).WithClues(ctx)
	}

	url := ptr.Val(session.GetUploadUrl())
	aw := uploadsession.NewWriter(uploader.getItemID(), url, size)
	logger.Ctx(ctx).Debugw("uploading large attachment", "attachment_url", url) // TODO: url pii

	// Upload the stream data
	copyBuffer := make([]byte, attachmentChunkSize)

	_, err = io.CopyBuffer(aw, bytes.NewReader(bs), copyBuffer)
	if err != nil {
		return clues.Wrap(err, "uploading large attachment").WithClues(ctx)
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

	return ptr.Val(item.GetOdataType())
}
