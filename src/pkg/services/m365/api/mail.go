package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/sanitize"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
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
	srv, err := NewService(c.Credentials, c.counter)
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
		return nil, graph.Stack(ctx, err)
	}

	gv := resp.GetValue()

	if len(gv) == 0 {
		return nil, clues.NewWC(ctx, "container not found")
	}

	// We only allow the api to match one container with the provided name.
	// Return an error if multiple container exist (unlikely) or if no container
	// is found.
	if len(gv) != 1 {
		return nil, clues.StackWC(ctx, graph.ErrMultipleResultsMatchIdentifier).
			With("returned_container_count", len(gv))
	}

	// Sanity check ID and name
	container := gv[0]

	if err := graph.CheckIDAndName(container); err != nil {
		return nil, clues.StackWC(ctx, err)
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

// TODO: needs pager implementation for completion
func (c Mail) GetContainerChildren(
	ctx context.Context,
	userID, containerID string,
) ([]models.MailFolderable, error) {
	resp, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		ChildFolders().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting container child folders")
	}

	return resp.GetValue(), nil
}

// ---------------------------------------------------------------------------
// items
// ---------------------------------------------------------------------------

// GetItem retrieves a Messageable item.  If the item contains an attachment, that
// attachment is also downloaded.
func (c Mail) GetItem(
	ctx context.Context,
	userID, mailID string,
	errs *fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	var (
		// ends up as len(mail.Body) + sum([]attachment.size)
		size     int64
		mailBody models.ItemBodyable
		config   = &users.ItemMessagesMessageItemRequestBuilderGetRequestConfiguration{
			Headers: newPreferHeaders(preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
		}
	)

	mail, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Messages().
		ByMessageId(mailID).
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

	attachments, totalSize, err := c.getAttachments(ctx, userID, mailID)
	if err != nil {
		// A failure can be caused by having a lot of attachments.
		// If that happens, we can progres with a two-step approach of:
		// 1. getting all attachment IDs.
		// 2. fetching each attachment individually.
		logger.CtxErr(ctx, err).Info("falling back to fetching attachments by id")

		attachments, totalSize, err = c.getAttachmentsIterated(
			ctx,
			userID,
			mailID,
			errs)
		if err != nil {
			return nil, nil, clues.Stack(err)
		}
	}

	size += totalSize

	mail.SetAttachments(attachments)

	return mail, MailInfo(mail, size), nil
}

// getAttachments attempts to get all attachments, including their content, in a singe query.
func (c Mail) getAttachments(
	ctx context.Context,
	userID, mailID string,
) ([]models.Attachmentable, int64, error) {
	var (
		result    = []models.Attachmentable{}
		totalSize int64
		cfg       = &users.ItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMessagesItemAttachmentsRequestBuilderGetQueryParameters{
				Expand: []string{"microsoft.graph.itemattachment/item"},
			},
			Headers: newPreferHeaders(
				preferPageSize(maxNonDeltaPageSize),
				preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
		}
	)

	attachments, err := c.LargeItem.
		Client().
		Users().
		ByUserId(userID).
		Messages().
		ByMessageId(mailID).
		Attachments().
		Get(ctx, cfg)
	if err != nil {
		return nil, 0, graph.Stack(ctx, err)
	}

	for _, a := range attachments.GetValue() {
		totalSize += int64(ptr.Val(a.GetSize()))
		result = append(result, a)
	}

	return result, totalSize, nil
}

// getAttachmentsIterated runs a two step fetch: one bulk query to get all attachment IDs,
// and then another lookup to fetch the content of each attachment.
// TODO: Once MS Graph fixes pagination for this, we can swap to a pager.
// https://learn.microsoft.com/en-us/answers/questions/1227026/pagination-not-working-when-fetching-message-attac
func (c Mail) getAttachmentsIterated(
	ctx context.Context,
	userID, mailID string,
	errs *fault.Bus,
) ([]models.Attachmentable, int64, error) {
	var (
		result    = []models.Attachmentable{}
		totalSize int64
		cfg       = &users.ItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMessagesItemAttachmentsRequestBuilderGetQueryParameters{
				Select: idAnd(),
			},
			Headers: newPreferHeaders(
				preferPageSize(maxNonDeltaPageSize),
				preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
		}
	)

	attachments, err := c.LargeItem.
		Client().
		Users().
		ByUserId(userID).
		Messages().
		ByMessageId(mailID).
		Attachments().
		Get(ctx, cfg)
	if err != nil {
		return nil, 0, graph.Wrap(ctx, err, "getting mail attachment ids")
	}

	for _, a := range attachments.GetValue() {
		var (
			aID              = ptr.Val(a.GetId())
			aODataType       = ptr.Val(a.GetOdataType())
			isItemAttachment = aODataType == "#microsoft.graph.itemAttachment"
		)

		ictx := clues.Add(
			ctx,
			"attachment_id", aID,
			"attachment_odatatype", aODataType)

		attachment, err := c.getAttachmentByID(
			ictx,
			userID,
			mailID,
			aID,
			isItemAttachment,
			errs)
		if err != nil {
			return nil, 0, clues.Stack(err)
		}

		if attachment != nil {
			result = append(result, attachment)
			totalSize += int64(ptr.Val(attachment.GetSize()))
		}
	}

	return result, totalSize, nil
}

func (c Mail) getAttachmentByID(
	ctx context.Context,
	userID, mailID, attachmentID string,
	isItemAttachment bool,
	errs *fault.Bus,
) (models.Attachmentable, error) {
	cfg := &users.ItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters{
			Expand: []string{"microsoft.graph.itemattachment/item"},
		},
		Headers: newPreferHeaders(preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
	}

	attachment, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Messages().
		ByMessageId(mailID).
		Attachments().
		ByAttachmentId(attachmentID).
		Get(ctx, cfg)
	if err != nil {
		// CannotOpenFileAttachment errors are not transient and
		// happens possibly from the original item somehow getting
		// deleted from M365 and so we can skip these
		if graph.IsErrCannotOpenFileAttachment(err) {
			logger.CtxErr(ctx, err).Info("attachment not found")
			errs.AddAlert(ctx, fault.NewAlert(
				"cannot open attached file",
				"", // no namespace
				mailID,
				"mailAttachment",
				map[string]any{
					"attachment_id":      attachmentID,
					"user_id":            userID,
					"is_item_attachment": isItemAttachment,
				}))
			// TODO This should use a `AddSkip` once we have
			// figured out the semantics for skipping
			// subcomponents of an item

			return nil, nil
		}

		return nil, graph.Wrap(ctx, err, "getting mail attachment by id")
	}

	return attachment, nil
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
		return nil, clues.NewWC(ctx, "nil response mail message creation")
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
	srv, err := NewService(c.Credentials, c.counter)
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
	w := graph.NewLargeItemWriter(parentItemID, url, size, c.counter)
	copyBuffer := make([]byte, graph.AttachmentChunkSize)

	_, err = io.CopyBuffer(w, bytes.NewReader(content), copyBuffer)
	if err != nil {
		return "", clues.WrapWC(ctx, err, "buffering large attachment content")
	}

	return w.ID, nil
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

func bytesToMessageable(body []byte) (serialization.Parsable, error) {
	v, err := CreateFromBytes(body, models.CreateMessageFromDiscriminatorValue)
	if err != nil {
		if !strings.Contains(err.Error(), invalidJSON) {
			return nil, clues.Wrap(err, "deserializing bytes to message")
		}

		// If the JSON was invalid try sanitizing and deserializing again.
		// Sanitizing should transform characters < 0x20 according to the spec where
		// possible. The resulting JSON may still be invalid though.
		body = sanitize.JSONBytes(body)
		v, err = CreateFromBytes(body, models.CreateMessageFromDiscriminatorValue)
	}

	return v, clues.Stack(err).OrNil()
}

func BytesToMessageable(body []byte) (models.Messageable, error) {
	v, err := bytesToMessageable(body)
	if err != nil {
		return nil, clues.Stack(err)
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
		sender     = unwrapEmailAddress(msg.GetSender())
		subject    = ptr.Val(msg.GetSubject())
		received   = ptr.Val(msg.GetReceivedDateTime())
		created    = ptr.Val(msg.GetCreatedDateTime())
		recipients = make([]string, 0)
	)

	if msg.GetToRecipients() != nil {
		ppl := msg.GetToRecipients()
		for _, entry := range ppl {
			temp := unwrapEmailAddress(entry)
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

func unwrapEmailAddress(contact models.Recipientable) string {
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
