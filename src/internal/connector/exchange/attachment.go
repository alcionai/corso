package exchange

import (
	"bytes"
	"context"
	"io"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	ups "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/attachments/createuploadsession"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item/attachments/createuploadsession"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
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
func uploadAttachment(ctx context.Context, service graph.Service, userID, folderID, messageID string,
	attachment models.Attachmentable,
) error {
	logger.Ctx(ctx).Debugf("uploading attachment with size %d", *attachment.GetSize())

	// For Item/Reference attachments *or* file attachments < 3MB, use the attachments endpoint
	if attachmentType(attachment) != models.FILE_ATTACHMENTTYPE || *attachment.GetSize() < largeAttachmentSize {
		_, err := service.Client().
			UsersById(userID).
			MailFoldersById(folderID).
			MessagesById(messageID).
			Attachments().
			Post(ctx, attachment, nil)

		return err
	}

	return uploadLargeAttachment(ctx, service, userID, folderID, messageID, attachment)
}

func uploadEventAttachment(ctx context.Context, service graph.Service, userID, calendardID, eventID string,
	attachment models.Attachmentable,
) error {
	logger.Ctx(ctx).Debugf("uploading event attachment with size %d", *attachment.GetSize())

	if attachmentType(attachment) != models.FILE_ATTACHMENTTYPE || *attachment.GetSize() < largeAttachmentSize {
		_, err := service.Client().
			UsersById(userID).
			CalendarsById(calendardID).
			EventsById(eventID).
			Attachments().
			Post(ctx, attachment, nil)

		return err
	}

	return uploadLargeEventAttachment(ctx, service, userID, calendardID, eventID, attachment)
}

// uploadLargeAttachment will upload the specified attachment for an event using an upload session
func uploadLargeEventAttachment(ctx context.Context, service graph.Service, userID, calendarID, eventID string,
	attachment models.Attachmentable,
) error {
	ab := attachmentBytes(attachment)

	aw, err := attachmentEventWriter(ctx, service, userID, calendarID, eventID, attachment, int64(len(ab)))
	if err != nil {
		return err
	}

	copyBuffer := make([]byte, attachmentChunkSize)

	_, err = io.CopyBuffer(aw, bytes.NewReader(ab), copyBuffer)
	if err != nil {
		return errors.Wrapf(err, "failed to upload attachment: item %s", eventID)
	}

	return nil
}

// uploadLargeAttachment will upload the specified attachment by creating an upload session and
// doing a chunked upload
func uploadLargeAttachment(ctx context.Context, service graph.Service, userID, folderID, messageID string,
	attachment models.Attachmentable,
) error {
	ab := attachmentBytes(attachment)

	aw, err := attachmentWriter(ctx, service, userID, folderID, messageID, attachment, int64(len(ab)))
	if err != nil {
		return err
	}

	// Upload the stream data
	copyBuffer := make([]byte, attachmentChunkSize)

	_, err = io.CopyBuffer(aw, bytes.NewReader(ab), copyBuffer)
	if err != nil {
		return errors.Wrapf(err, "failed to upload attachment: item %s", messageID)
	}

	return nil
}

// attachmentWriter is used to initialize and return an io.Writer to upload data for the specified attachment
// It does so by creating an upload session and using that URL to initialize an `itemWriter`
func attachmentWriter(ctx context.Context, service graph.Service, userID, folderID, messageID string,
	attachment models.Attachmentable, size int64,
) (io.Writer, error) {
	session := createuploadsession.NewCreateUploadSessionPostRequestBody()

	attItem := models.NewAttachmentItem()
	attType := models.FILE_ATTACHMENTTYPE
	attItem.SetAttachmentType(&attType)
	attItem.SetName(attachment.GetName())
	attItem.SetSize(&size)
	session.SetAttachmentItem(attItem)

	r, err := service.Client().UsersById(userID).MailFoldersById(folderID).
		MessagesById(messageID).Attachments().CreateUploadSession().Post(ctx, session, nil)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to create attachment upload session for item %s. details: %s",
			messageID,
			support.ConnectorStackErrorTrace(err),
		)
	}

	url := *r.GetUploadUrl()

	logger.Ctx(ctx).Debugf("Created an upload session for item %s. URL: %s", messageID, url)

	return uploadsession.NewWriter(messageID, url, size), nil
}

func attachmentEventWriter(ctx context.Context, service graph.Service, userID, calendarID, eventID string,
	attachment models.Attachmentable, size int64,
) (io.Writer, error) {
	session := ups.NewCreateUploadSessionPostRequestBody()

	attItem := models.NewAttachmentItem()
	attType := models.FILE_ATTACHMENTTYPE
	attItem.SetAttachmentType(&attType)
	attItem.SetName(attachment.GetName())
	attItem.SetSize(&size)
	session.SetAttachmentItem(attItem)

	r, err := service.Client().
		UsersById(userID).
		CalendarsById(calendarID).
		EventsById(eventID).
		Attachments().
		CreateUploadSession().
		Post(ctx, session, nil)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to create attachment upload session for event item %s. details: %s",
			eventID, support.ConnectorStackErrorTrace(err),
		)
	}

	url := *r.GetUploadUrl()

	logger.Ctx(ctx).Debugf("Created an upload session for item %s. URL: %s", eventID, url)

	return uploadsession.NewWriter(eventID, url, size), nil
}
