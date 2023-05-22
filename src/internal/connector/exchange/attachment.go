package exchange

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/logger"
)

type attachmentPoster interface {
	PostSmallAttachment(
		ctx context.Context,
		userID, containerID, itemID string,
		body models.Attachmentable,
	) error
	PostLargeAttachment(
		ctx context.Context,
		userID, containerID, itemID, name string,
		size int64,
		body models.Attachmentable,
	) (models.UploadSessionable, error)
}

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
	cli attachmentPoster,
	userID, containerID, itemID, name string,
	size int32,
	attachment models.Attachmentable,
) error {
	attachmentType := attachmentType(attachment)

	ctx = clues.Add(
		ctx,
		"attachment_size", size,
		"attachment_id", itemID,
		"attachment_name", clues.Hide(name),
		"attachment_type", attachmentType,
		"internal_item_type", getItemAttachmentItemType(attachment))

	logger.Ctx(ctx).Debug("uploading attachment")

	// reference attachments that are inline() do not need to be recreated. The contents are part of the body.
	if attachmentType == models.REFERENCE_ATTACHMENTTYPE && ptr.Val(attachment.GetIsInline()) {
		logger.Ctx(ctx).Debug("skip uploading inline reference attachment")
		return nil
	}

	// item Attachments to be skipped until the completion of Issue #2353
	if attachmentType == models.ITEM_ATTACHMENTTYPE {
		a, err := support.ToItemAttachment(attachment)
		if err != nil {
			logger.CtxErr(ctx, err).Info(fmt.Sprintf("item attachment type not supported: %v", attachmentType))
			return nil
		}

		attachment = a
	}

	// for Item/Reference attachments *or* file attachments < 3MB
	if attachmentType != models.FILE_ATTACHMENTTYPE || ptr.Val(attachment.GetSize()) < largeAttachmentSize {
		return cli.PostSmallAttachment(ctx, userID, containerID, itemID, attachment)
	}

	// for all other attachments
	_, err := cli.PostLargeAttachment(ctx, userID, containerID, itemID, name, int64(size), attachment)

	return err
}

func getItemAttachmentItemType(query models.Attachmentable) string {
	empty := ""
	attachment, ok := query.(models.ItemAttachmentable)

	if !ok {
		return empty
	}

	item := attachment.GetItem()
	if item == nil {
		return empty
	}

	return ptr.Val(item.GetOdataType())
}
