package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msusers "github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/connector/graph"
)

// attachementUploadable represents structs that are able to upload small attachments directly to an item or use an
// upload session to connect large attachments to their corresponding M365 item.
type attachmentUploadable interface {
	uploadSmallAttachment(ctx context.Context, attachment models.Attachmentable) error
	uploadSession(ctx context.Context, attachName string, attachSize int64) (models.UploadSessionable, error)
	// getItemID returns the M365ID of the item associated with  the attachment
	getItemID() string
}

var (
	_ attachmentUploadable = &mailAttachmentUploader{}
	_ attachmentUploadable = &eventAttachmentUploader{}
)

// mailAttachmentUploader is a struct that is able to upload attachments for exchange.Mail objects
type mailAttachmentUploader struct {
	userID   string
	folderID string
	itemID   string
	service  graph.Servicer
}

func (mau *mailAttachmentUploader) getItemID() string {
	return mau.itemID
}

func (mau *mailAttachmentUploader) uploadSmallAttachment(ctx context.Context, attach models.Attachmentable) error {
	_, err := mau.service.Client().
		UsersById(mau.userID).
		MailFoldersById(mau.folderID).
		MessagesById(mau.itemID).
		Attachments().
		Post(ctx, attach, nil)
	if err != nil {
		return clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return nil
}

func (mau *mailAttachmentUploader) uploadSession(
	ctx context.Context,
	attachmentName string,
	attachmentSize int64,
) (models.UploadSessionable, error) {
	session := msusers.NewItemMailFoldersItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody()
	session.SetAttachmentItem(makeSessionAttachment(attachmentName, attachmentSize))

	r, err := mau.
		service.
		Client().
		UsersById(mau.userID).
		MailFoldersById(mau.folderID).
		MessagesById(mau.itemID).
		Attachments().
		CreateUploadSession().
		Post(ctx, session, nil)
	if err != nil {
		return nil, clues.Wrap(err, "uploading mail attachment").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return r, nil
}

// eventAttachmentUploader is a struct capable of uploading attachments for exchange.Event objects
type eventAttachmentUploader struct {
	userID     string
	calendarID string
	itemID     string
	service    graph.Servicer
}

func (eau *eventAttachmentUploader) getItemID() string {
	return eau.itemID
}

func (eau *eventAttachmentUploader) uploadSmallAttachment(ctx context.Context, attach models.Attachmentable) error {
	_, err := eau.service.Client().
		UsersById(eau.userID).
		CalendarsById(eau.calendarID).
		EventsById(eau.itemID).
		Attachments().
		Post(ctx, attach, nil)
	if err != nil {
		return clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return nil
}

func (eau *eventAttachmentUploader) uploadSession(
	ctx context.Context,
	attachmentName string,
	attachmentSize int64,
) (models.UploadSessionable, error) {
	session := msusers.NewItemCalendarEventsItemAttachmentsCreateUploadSessionPostRequestBody()
	session.SetAttachmentItem(makeSessionAttachment(attachmentName, attachmentSize))

	r, err := eau.service.Client().
		UsersById(eau.userID).
		CalendarsById(eau.calendarID).
		EventsById(eau.itemID).
		Attachments().
		CreateUploadSession().
		Post(ctx, session, nil)
	if err != nil {
		return nil, clues.Wrap(err, "uploading event attachment").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return r, nil
}

func makeSessionAttachment(name string, size int64) *models.AttachmentItem {
	attItem := models.NewAttachmentItem()
	attType := models.FILE_ATTACHMENTTYPE
	attItem.SetAttachmentType(&attType)
	attItem.SetName(&name)
	attItem.SetSize(&size)

	return attItem
}
