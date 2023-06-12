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

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
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

// CreateMailFolder makes a mail folder iff a folder of the same name does not exist
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-mailfolders?view=graph-rest-1.0&tabs=http
func (c Mail) CreateMailFolder(
	ctx context.Context,
	userID, containerName string,
) (models.MailFolderable, error) {
	isHidden := false
	body := models.NewMailFolder()
	body.SetDisplayName(&containerName)
	body.SetIsHidden(&isHidden)

	mdl, err := c.Stable.Client().
		Users().
		ByUserId(userID).
		MailFolders().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating mail folder")
	}

	return mdl, nil
}

func (c Mail) CreateContainer(
	ctx context.Context,
	userID, containerName, parentContainerID string,
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

// prefer GetContainerByID where possible.
// use this only in cases where the models.MailFolderable
// is required.
func (c Mail) GetFolder(
	ctx context.Context,
	userID, containerID string,
) (models.MailFolderable, error) {
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

// interface-compliant wrapper of GetFolder
func (c Mail) GetContainerByID(
	ctx context.Context,
	userID, containerID string,
) (graph.Container, error) {
	return c.GetFolder(ctx, userID, containerID)
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
// container pager
// ---------------------------------------------------------------------------

type mailFolderPager struct {
	service graph.Servicer
	builder *users.ItemMailFoldersRequestBuilder
}

func (c Mail) NewMailFolderPager(userID string) mailFolderPager {
	// v1.0 non delta /mailFolders endpoint does not return any of the nested folders
	rawURL := fmt.Sprintf(mailFoldersBetaURLTemplate, userID)
	builder := users.NewItemMailFoldersRequestBuilder(rawURL, c.Stable.Adapter())

	return mailFolderPager{c.Stable, builder}
}

func (p *mailFolderPager) getPage(ctx context.Context) (PageLinker, error) {
	page, err := p.builder.Get(ctx, nil)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return page, nil
}

func (p *mailFolderPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersRequestBuilder(nextLink, p.service.Adapter())
}

func (p *mailFolderPager) valuesIn(pl PageLinker) ([]models.MailFolderable, error) {
	// Ideally this should be `users.ItemMailFoldersResponseable`, but
	// that is not a thing as stable returns different result
	page, ok := pl.(models.MailFolderCollectionResponseable)
	if !ok {
		return nil, clues.New("converting to ItemMailFoldersResponseable")
	}

	return page.GetValue(), nil
}

// EnumerateContainers iterates through all of the users current
// mail folders, converting each to a graph.CacheFolder, and calling
// fn(cf) on each one.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Mail) EnumerateContainers(
	ctx context.Context,
	userID, baseContainerID string,
	fn func(graph.CachedContainer) error,
	errs *fault.Bus,
) error {
	el := errs.Local()
	pgr := c.NewMailFolderPager(userID)

	for {
		if el.Failure() != nil {
			break
		}

		page, err := pgr.getPage(ctx)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		resp, err := pgr.valuesIn(page)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		for _, fold := range resp {
			if el.Failure() != nil {
				break
			}

			if err := graph.CheckIDNameAndParentFolderID(fold); err != nil {
				errs.AddRecoverable(ctx, graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			fctx := clues.Add(
				ctx,
				"container_id", ptr.Val(fold.GetId()),
				"container_name", ptr.Val(fold.GetDisplayName()))

			temp := graph.NewCacheFolder(fold, nil, nil)
			if err := fn(&temp); err != nil {
				errs.AddRecoverable(ctx, graph.Stack(fctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}
		}

		link, ok := ptr.ValOK(page.GetOdataNextLink())
		if !ok {
			break
		}

		pgr.setNext(link)
	}

	return el.Failure()
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
	size int64,
	body models.Attachmentable,
) (models.UploadSessionable, error) {
	bs, err := GetAttachmentContent(body)
	if err != nil {
		return nil, clues.Wrap(err, "serializing attachment content").WithClues(ctx)
	}

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
		return nil, graph.Wrap(ctx, err, "uploading large mail attachment")
	}

	url := ptr.Val(us.GetUploadUrl())
	w := graph.NewLargeItemWriter(parentItemID, url, size)
	copyBuffer := make([]byte, graph.AttachmentChunkSize)

	_, err = io.CopyBuffer(w, bytes.NewReader(bs), copyBuffer)
	if err != nil {
		return nil, clues.Wrap(err, "buffering large attachment content").WithClues(ctx)
	}

	return us, nil
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager = &mailPager{}

type mailPager struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersItemMessagesRequestBuilder
	options *users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration
}

func (c Mail) NewMailPager(
	ctx context.Context,
	userID, containerID string,
	immutableIDs bool,
) itemPager {
	config := &users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersItemMessagesRequestBuilderGetQueryParameters{
			Select: idAnd("isRead"),
		},
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	builder := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Messages()

	return &mailPager{c.Stable, builder, config}
}

func (p *mailPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.Messageable]{PageLinkValuer: page}, nil
}

func (p *mailPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't have reset
func (p *mailPager) reset(context.Context) {}

func (p *mailPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Messageable](pl)
}

// ---------------------------------------------------------------------------
// delta item pager
// ---------------------------------------------------------------------------

var _ itemPager = &mailDeltaPager{}

type mailDeltaPager struct {
	gs          graph.Servicer
	userID      string
	containerID string
	builder     *users.ItemMailFoldersItemMessagesDeltaRequestBuilder
	options     *users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func getMailDeltaBuilder(
	ctx context.Context,
	gs graph.Servicer,
	user, containerID string,
	options *users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration,
) *users.ItemMailFoldersItemMessagesDeltaRequestBuilder {
	builder := gs.
		Client().
		Users().
		ByUserId(user).
		MailFolders().
		ByMailFolderId(containerID).
		Messages().
		Delta()

	return builder
}

func (c Mail) NewMailDeltaPager(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
) itemPager {
	config := &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{
			Select: idAnd("isRead"),
		},
		Headers: newPreferHeaders(preferPageSize(maxDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	var builder *users.ItemMailFoldersItemMessagesDeltaRequestBuilder

	if len(oldDelta) > 0 {
		builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(oldDelta, c.Stable.Adapter())
	} else {
		builder = getMailDeltaBuilder(ctx, c.Stable, userID, containerID, config)
	}

	return &mailDeltaPager{c.Stable, userID, containerID, builder, config}
}

func (p *mailDeltaPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return page, nil
}

func (p *mailDeltaPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *mailDeltaPager) reset(ctx context.Context) {
	p.builder = p.gs.
		Client().
		Users().
		ByUserId(p.userID).
		MailFolders().
		ByMailFolderId(p.containerID).
		Messages().
		Delta()
}

func (p *mailDeltaPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Messageable](pl)
}

func (c Mail) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	ctx = clues.Add(
		ctx,
		"category", selectors.ExchangeMail,
		"container_id", containerID)

	pager := c.NewMailPager(ctx, userID, containerID, immutableIDs)
	deltaPager := c.NewMailDeltaPager(ctx, userID, containerID, oldDelta, immutableIDs)

	return getAddedAndRemovedItemIDs(ctx, c.Stable, pager, deltaPager, oldDelta, canMakeDeltaQueries)
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
