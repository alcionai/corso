package api

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	mailFoldersBetaURLTemplate = "https://graph.microsoft.com/beta/users/%s/mailFolders"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Mail() Mail {
	return Mail{c}
}

// Mail is an interface-compliant provider of the client.
type Mail struct {
	Client
}

// ---------------------------------------------------------------------------
// containers
// ---------------------------------------------------------------------------

func (c Mail) CreateContainer(
	ctx context.Context,
	userID, parentContainerID, containerName string,
) (graph.Container, error) {
	isHidden := false
	body := models.NewMailFolder()
	body.SetDisplayName(&containerName)
	body.SetIsHidden(&isHidden)

	mdl, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(parentContainerID).
		ChildFolders().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating nested mail folder")
	}

	return mdl, nil
}

// DeleteContainer removes a mail folder with the corresponding M365 ID  from the user's M365 Exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/mailfolder-delete?view=graph-rest-1.0&tabs=http
func (c Mail) DeleteContainer(
	ctx context.Context,
	userID, containerID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

func (c Mail) GetContainerByID(
	ctx context.Context,
	userID, containerID string,
) (graph.Container, error) {
	config := &users.ItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersMailFolderItemRequestBuilderGetQueryParameters{
			Select: idAnd(displayName, parentFolderID),
		},
	}

	resp, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Get(ctx, config)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

// GetContainerByName fetches a folder by name
func (c Mail) GetContainerByName(
	ctx context.Context,
	userID, parentContainerID, containerName string,
) (graph.Container, error) {
	filter := fmt.Sprintf("displayName eq '%s'", containerName)

	ctx = clues.Add(ctx, "container_name", containerName)

	var (
		builder = c.Stable.
			Client().
			Users().
			ByUserId(userID).
			MailFolders()
		resp models.MailFolderCollectionResponseable
		err  error
	)

	if len(parentContainerID) > 0 {
		options := &users.ItemMailFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMailFoldersItemChildFoldersRequestBuilderGetQueryParameters{
				Filter: &filter,
			},
		}

		resp, err = builder.
			ByMailFolderId(parentContainerID).
			ChildFolders().
			Get(ctx, options)
	} else {
		options := &users.ItemMailFoldersRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMailFoldersRequestBuilderGetQueryParameters{
				Filter: &filter,
			},
		}

		resp, err = builder.Get(ctx, options)
	}

	if err != nil {
		return nil, graph.Stack(ctx, err).WithClues(ctx)
	}

	gv := resp.GetValue()

	if len(gv) == 0 {
		return nil, clues.New("container not found").WithClues(ctx)
	}

	// We only allow the api to match one container with the provided name.
	// Return an error if multiple container exist (unlikely) or if no container
	// is found.
	if len(gv) != 1 {
		return nil, clues.New("unexpected number of folders returned").
			With("returned_container_count", len(gv)).
			WithClues(ctx)
	}

	// Sanity check ID and name
	container := gv[0]

	if err := graph.CheckIDAndName(container); err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return container, nil
}

func (c Mail) MoveContainer(
	ctx context.Context,
	userID, containerID string,
	body users.ItemMailFoldersItemMovePostRequestBodyable,
) error {
	_, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Move().
		Post(ctx, body, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "moving mail folder")
	}

	return nil
}

func (c Mail) PatchFolder(
	ctx context.Context,
	userID, containerID string,
	body models.MailFolderable,
) error {
	_, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Patch(ctx, body, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "patching mail folder")
	}

	return nil
}

// ---------------------------------------------------------------------------
// items
// ---------------------------------------------------------------------------

// GetItem retrieves a Messageable item.  If the item contains an attachment, that
// attachment is also downloaded.
func (c Mail) GetItem(
	ctx context.Context,
	userID, itemID string,
	immutableIDs bool,
	errs *fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	var (
		size     int64
		mailBody models.ItemBodyable
		config   = &users.ItemMessagesMessageItemRequestBuilderGetRequestConfiguration{
			Headers: newPreferHeaders(preferImmutableIDs(immutableIDs)),
		}
	)

	mail, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Messages().
		ByMessageId(itemID).
		Get(ctx, config)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	mailBody = mail.GetBody()
	if mailBody != nil {
		content := ptr.Val(mailBody.GetContent())
		if len(content) > 0 {
			size = int64(len(content))
		}
	}

	if !ptr.Val(mail.GetHasAttachments()) && !HasAttachments(mailBody) {
		return mail, MailInfo(mail, size), nil
	}

	attachConfig := &users.ItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMessagesItemAttachmentsRequestBuilderGetQueryParameters{
			Expand: []string{"microsoft.graph.itemattachment/item"},
		},
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	attached, err := c.LargeItem.
		Client().
		Users().
		ByUserId(userID).
		Messages().
		ByMessageId(itemID).
		Attachments().
		Get(ctx, attachConfig)
	if err == nil {
		for _, a := range attached.GetValue() {
			attachSize := ptr.Val(a.GetSize())
			size = +int64(attachSize)
		}

		mail.SetAttachments(attached.GetValue())

		return mail, MailInfo(mail, size), nil
	}

	// A failure can be caused by having a lot of attachments as
	// we are trying to fetch the data within the attachments as
	// well in the request. We instead fetch all the attachment
	// ids and fetch each item individually.
	// NOTE: Maybe filter for specific error:
	// graph.IsErrTimeout(err) || graph.IsServiceUnavailable(err)
	// TODO: Once MS Graph fixes pagination for this, we can
	// probably paginate and fetch items.
	// https://learn.microsoft.com/en-us/answers/questions/1227026/pagination-not-working-when-fetching-message-attac
	logger.CtxErr(ctx, err).Info("fetching all attachments by id")

	// Getting size just to log in case of error
	attachConfig.QueryParameters.Select = []string{"id", "size"}

	attachments, err := c.LargeItem.
		Client().
		Users().
		ByUserId(userID).
		Messages().
		ByMessageId(itemID).
		Attachments().
		Get(ctx, attachConfig)
	if err != nil {
		return nil, nil, graph.Wrap(ctx, err, "getting mail attachment ids")
	}

	atts := []models.Attachmentable{}

	for _, a := range attachments.GetValue() {
		attachConfig := &users.ItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters{
				Expand: []string{"microsoft.graph.itemattachment/item"},
			},
			Headers: newPreferHeaders(preferImmutableIDs(immutableIDs)),
		}

		att, err := c.Stable.
			Client().
			Users().
			ByUserId(userID).
			Messages().
			ByMessageId(itemID).
			Attachments().
			ByAttachmentId(ptr.Val(a.GetId())).
			Get(ctx, attachConfig)
		if err != nil {
			// CannotOpenFileAttachment errors are not transient and
			// happens possibly from the original item somehow getting
			// deleted from M365 and so we can skip these
			if graph.IsErrCannotOpenFileAttachment(err) {
				logger.CtxErr(ctx, err).
					With(
						"skipped_reason", fault.SkipNotFound,
						"attachment_id", ptr.Val(a.GetId()),
						"attachment_size", ptr.Val(a.GetSize()),
					).Info("attachment not found")
				// TODO This should use a `AddSkip` once we have
				// figured out the semantics for skipping
				// subcomponents of an item

				continue
			}

			return nil, nil, graph.Wrap(ctx, err, "getting mail attachment").
				With("attachment_id", ptr.Val(a.GetId()), "attachment_size", ptr.Val(a.GetSize()))
		}

		atts = append(atts, att)
		attachSize := ptr.Val(a.GetSize())
		size = +int64(attachSize)
	}

	mail.SetAttachments(atts)

	return mail, MailInfo(mail, size), nil
}

func (c Mail) PostItem(
	ctx context.Context,
	userID, containerID string,
	body models.Messageable,
) (models.Messageable, error) {
	itm, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Messages().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating mail message")
	}

	if itm == nil {
		return nil, clues.New("nil response mail message creation").WithClues(ctx)
	}

	return itm, nil
}

func (c Mail) MoveItem(
	ctx context.Context,
	userID, oldContainerID, newContainerID, itemID string,
) (string, error) {
	body := users.NewItemMailFoldersItemMessagesItemMovePostRequestBody()
	body.SetDestinationId(ptr.To(newContainerID))

	resp, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(oldContainerID).
		Messages().
		ByMessageId(itemID).
		Move().
		Post(ctx, body, nil)
	if err != nil {
		return "", graph.Wrap(ctx, err, "moving message")
	}

	return ptr.Val(resp.GetId()), nil
}

func (c Mail) DeleteItem(
	ctx context.Context,
	userID, itemID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.
		Client().
		Users().
		ByUserId(userID).
		Messages().
		ByMessageId(itemID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting mail message")
	}

	return nil
}

func (c Mail) PostSmallAttachment(
	ctx context.Context,
	userID, containerID, parentItemID string,
	body models.Attachmentable,
) error {
	_, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Messages().
		ByMessageId(parentItemID).
		Attachments().
		Post(ctx, body, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "uploading small mail attachment")
	}

	return nil
}

func (c Mail) PostLargeAttachment(
	ctx context.Context,
	userID, containerID, parentItemID, itemName string,
	content []byte,
) (string, error) {
	size := int64(len(content))
	session := users.NewItemMailFoldersItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody()
	session.SetAttachmentItem(makeSessionAttachment(itemName, size))

	us, err := c.LargeItem.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Messages().
		ByMessageId(parentItemID).
		Attachments().
		CreateUploadSession().
		Post(ctx, session, nil)
	if err != nil {
		return "", graph.Wrap(ctx, err, "uploading large mail attachment")
	}

	url := ptr.Val(us.GetUploadUrl())
	w := graph.NewLargeItemWriter(parentItemID, url, size)
	copyBuffer := make([]byte, graph.AttachmentChunkSize)

	_, err = io.CopyBuffer(w, bytes.NewReader(content), copyBuffer)
	if err != nil {
		return "", clues.Wrap(err, "buffering large attachment content").WithClues(ctx)
	}

	return w.ID, nil
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

func BytesToMessageable(body []byte) (models.Messageable, error) {
	v, err := createFromBytes(body, models.CreateMessageFromDiscriminatorValue)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes to message")
	}

	return v.(models.Messageable), nil
}

func (c Mail) Serialize(
	ctx context.Context,
	item serialization.Parsable,
	user, itemID string,
) ([]byte, error) {
	msg, ok := item.(models.Messageable)
	if !ok {
		return nil, clues.New(fmt.Sprintf("item is not a Messageable: %T", item))
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(msg.GetId()))
	writer := kjson.NewJsonSerializationWriter()

	defer writer.Close()

	if err := writer.WriteObjectValue("", msg); err != nil {
		return nil, graph.Stack(ctx, err)
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, graph.Wrap(ctx, err, "serializing email")
	}

	return bs, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func MailInfo(msg models.Messageable, size int64) *details.ExchangeInfo {
	var (
		sender     = UnwrapEmailAddress(msg.GetSender())
		subject    = ptr.Val(msg.GetSubject())
		received   = ptr.Val(msg.GetReceivedDateTime())
		created    = ptr.Val(msg.GetCreatedDateTime())
		recipients = make([]string, 0)
	)

	if msg.GetToRecipients() != nil {
		ppl := msg.GetToRecipients()
		for _, entry := range ppl {
			temp := UnwrapEmailAddress(entry)
			if len(temp) > 0 {
				recipients = append(recipients, temp)
			}
		}
	}

	return &details.ExchangeInfo{
		ItemType:  details.ExchangeMail,
		Sender:    sender,
		Recipient: recipients,
		Subject:   subject,
		Received:  received,
		Size:      size,
		Created:   created,
		Modified:  ptr.OrNow(msg.GetLastModifiedDateTime()),
	}
}

func UnwrapEmailAddress(contact models.Recipientable) string {
	var empty string
	if contact == nil || contact.GetEmailAddress() == nil {
		return empty
	}

	return ptr.Val(contact.GetEmailAddress().GetAddress())
}

func mailCollisionKeyProps() []string {
	return idAnd("subject", sentDateTime, receivedDateTime)
}

// MailCollisionKey constructs a key from the messageable's subject, sender, and recipients (to, cc, bcc).
// collision keys are used to identify duplicate item conflicts for handling advanced restoration config.
func MailCollisionKey(item models.Messageable) string {
	if item == nil {
		return ""
	}

	var (
		subject  = ptr.Val(item.GetSubject())
		sent     = ptr.Val(item.GetSentDateTime())
		received = ptr.Val(item.GetReceivedDateTime())
	)

	return subject + dttm.Format(sent) + dttm.Format(received)
}
