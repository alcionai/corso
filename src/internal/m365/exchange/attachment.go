package exchange

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
		content []byte,
	) (string, error)
}

const (
	// Use large attachment logic for attachments > 3MB
	// https://learn.microsoft.com/en-us/graph/outlook-large-attachments
	largeAttachmentSize           = 3 * 1024 * 1024
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
	userID, containerID, parentItemID string,
	attachment models.Attachmentable,
) error {
	var (
		attachmentType = attachmentType(attachment)
		id             = ptr.Val(attachment.GetId())
		name           = ptr.Val(attachment.GetName())
		size           = ptr.Val(attachment.GetSize())
	)

	ctx = clues.Add(
		ctx,
		"attachment_size", size,
		"attachment_id", id,
		"attachment_name", clues.Hide(name),
		"attachment_type", attachmentType,
		"attachment_odata_type", ptr.Val(attachment.GetOdataType()),
		"attachment_outlook_odata_type", getOutlookOdataType(attachment),
		"parent_item_id", parentItemID)

	logger.Ctx(ctx).Debug("uploading attachment")

	// reference attachments that are inline() do not need to be recreated. The contents are part of the body.
	if attachmentType == models.REFERENCE_ATTACHMENTTYPE && ptr.Val(attachment.GetIsInline()) {
		logger.Ctx(ctx).Debug("skip uploading inline reference attachment")
		return nil
	}

	// item Attachments to be skipped until the completion of Issue #2353
	if attachmentType == models.ITEM_ATTACHMENTTYPE {
		a, err := toItemAttachment(attachment)
		if err != nil {
			logger.CtxErr(ctx, err).Info(fmt.Sprintf("item attachment type not supported: %v", attachmentType))
			return nil
		}

		attachment = a
	}

	// for file attachments sized >= 3MB
	if attachmentType == models.FILE_ATTACHMENTTYPE && size >= largeAttachmentSize {
		content, err := api.GetAttachmentContent(attachment)
		if err != nil {
			return clues.Wrap(err, "serializing attachment content").WithClues(ctx)
		}

		_, err = cli.PostLargeAttachment(ctx, userID, containerID, parentItemID, name, content)

		return err
	}

	// for all other attachments
	return cli.PostSmallAttachment(ctx, userID, containerID, parentItemID, attachment)
}

func getOutlookOdataType(query models.Attachmentable) string {
	attachment, ok := query.(models.ItemAttachmentable)
	if !ok {
		return ""
	}

	item := attachment.GetItem()
	if item == nil {
		return ""
	}

	return ptr.Val(item.GetOdataType())
}
