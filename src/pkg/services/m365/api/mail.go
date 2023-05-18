package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
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
// methods
// ---------------------------------------------------------------------------

// CreateMailFolder makes a mail folder iff a folder of the same name does not exist
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-mailfolders?view=graph-rest-1.0&tabs=http
func (c Mail) CreateMailFolder(
	ctx context.Context,
	user, folder string,
) (models.MailFolderable, error) {
	isHidden := false
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	requestBody.SetIsHidden(&isHidden)

	mdl, err := c.Stable.Client().Users().ByUserId(user).MailFolders().Post(ctx, requestBody, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating mail folder")
	}

	return mdl, nil
}

func (c Mail) CreateMailFolderWithParent(
	ctx context.Context,
	user, folder, parentID string,
) (models.MailFolderable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	isHidden := false
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	requestBody.SetIsHidden(&isHidden)

	mdl, err := service.
		Client().
		Users().
		ByUserId(user).
		MailFolders().
		ByMailFolderId(parentID).
		ChildFolders().
		Post(ctx, requestBody, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating nested mail folder")
	}

	return mdl, nil
}

// DeleteContainer removes a mail folder with the corresponding M365 ID  from the user's M365 Exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/mailfolder-delete?view=graph-rest-1.0&tabs=http
func (c Mail) DeleteContainer(
	ctx context.Context,
	user, folderID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.Client().
		Users().
		ByUserId(user).
		MailFolders().
		ByMailFolderId(folderID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

func (c Mail) GetContainerByID(
	ctx context.Context,
	userID, dirID string,
) (graph.Container, error) {
	service, err := c.Service()
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	queryParams := &users.ItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersMailFolderItemRequestBuilderGetQueryParameters{
			Select: []string{"id", "displayName", "parentFolderId"},
		},
	}

	resp, err := service.Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(dirID).
		Get(ctx, queryParams)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

// GetItem retrieves a Messageable item.  If the item contains an attachment, that
// attachment is also downloaded.
func (c Mail) GetItem(
	ctx context.Context,
	user, itemID string,
	immutableIDs bool,
	errs *fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	var (
		size     int64
		mailBody models.ItemBodyable
	)
	// Will need adjusted if attachments start allowing paging.
	headers := buildPreferHeaders(false, immutableIDs)
	itemOpts := &users.ItemMessagesMessageItemRequestBuilderGetRequestConfiguration{
		Headers: headers,
	}

	mail, err := c.Stable.Client().Users().ByUserId(user).Messages().ByMessageId(itemID).Get(ctx, itemOpts)
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

	options := &users.ItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMessagesItemAttachmentsRequestBuilderGetQueryParameters{
			Expand: []string{"microsoft.graph.itemattachment/item"},
		},
		Headers: headers,
	}

	attached, err := c.LargeItem.
		Client().
		Users().
		ByUserId(user).
		Messages().
		ByMessageId(itemID).
		Attachments().
		Get(ctx, options)
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
	options.QueryParameters.Select = []string{"id", "size"}

	attachments, err := c.LargeItem.
		Client().
		Users().
		ByUserId(user).
		Messages().
		ByMessageId(itemID).
		Attachments().
		Get(ctx, options)
	if err != nil {
		return nil, nil, graph.Wrap(ctx, err, "getting mail attachment ids")
	}

	atts := []models.Attachmentable{}

	for _, a := range attachments.GetValue() {
		options := &users.ItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters{
				Expand: []string{"microsoft.graph.itemattachment/item"},
			},
			Headers: headers,
		}

		att, err := c.Stable.
			Client().
			Users().
			ByUserId(user).
			Messages().
			ByMessageId(itemID).
			Attachments().
			ByAttachmentId(ptr.Val(a.GetId())).
			Get(ctx, options)
		if err != nil {
			return nil, nil,
				graph.Wrap(ctx, err, "getting mail attachment").
					With("attachment_id", ptr.Val(a.GetId()), "attachment_size", ptr.Val(a.GetSize()))
		}

		atts = append(atts, att)
		attachSize := ptr.Val(a.GetSize())
		size = +int64(attachSize)
	}

	mail.SetAttachments(atts)

	return mail, MailInfo(mail, size), nil
}

type mailFolderPager struct {
	service graph.Servicer
	builder *users.ItemMailFoldersRequestBuilder
}

func NewMailFolderPager(service graph.Servicer, user string) mailFolderPager {
	// v1.0 non delta /mailFolders endpoint does not return any of the nested folders
	rawURL := fmt.Sprintf(mailFoldersBetaURLTemplate, user)
	builder := users.NewItemMailFoldersRequestBuilder(rawURL, service.Adapter())

	return mailFolderPager{service, builder}
}

func (p *mailFolderPager) getPage(ctx context.Context) (api.PageLinker, error) {
	page, err := p.builder.Get(ctx, nil)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return page, nil
}

func (p *mailFolderPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersRequestBuilder(nextLink, p.service.Adapter())
}

func (p *mailFolderPager) valuesIn(pl api.PageLinker) ([]models.MailFolderable, error) {
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
	userID, baseDirID string,
	fn func(graph.CacheFolder) error,
	errs *fault.Bus,
) error {
	service, err := c.Service()
	if err != nil {
		return graph.Stack(ctx, err)
	}

	el := errs.Local()

	pgr := NewMailFolderPager(service, userID)

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

		for _, v := range resp {
			if el.Failure() != nil {
				break
			}

			fctx := clues.Add(
				ctx,
				"container_id", ptr.Val(v.GetId()),
				"container_name", ptr.Val(v.GetDisplayName()))

			temp := graph.NewCacheFolder(v, nil, nil)
			if err := fn(temp); err != nil {
				errs.AddRecoverable(graph.Stack(fctx, err).Label(fault.LabelForceNoBackupCreation))
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
// item pager
// ---------------------------------------------------------------------------

var _ itemPager = &mailPager{}

type mailPager struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersItemMessagesRequestBuilder
	options *users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration
}

func NewMailPager(
	ctx context.Context,
	gs graph.Servicer,
	user, directoryID string,
	immutableIDs bool,
) itemPager {
	queryParams := &users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersItemMessagesRequestBuilderGetQueryParameters{
			Select: []string{"id", "isRead"},
		},
		Headers: buildPreferHeaders(true, immutableIDs),
	}

	builder := gs.Client().
		Users().
		ByUserId(user).
		MailFolders().
		ByMailFolderId(directoryID).
		Messages()

	return &mailPager{gs, builder, queryParams}
}

func (p *mailPager) getPage(ctx context.Context) (api.DeltaPageLinker, error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return api.EmptyDeltaLinker[models.Messageable]{PageLinkValuer: page}, nil
}

func (p *mailPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't have reset
func (p *mailPager) reset(context.Context) {}

func (p *mailPager) valuesIn(pl api.PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Messageable](pl)
}

// ---------------------------------------------------------------------------
// delta item pager
// ---------------------------------------------------------------------------

var _ itemPager = &mailDeltaPager{}

type mailDeltaPager struct {
	gs          graph.Servicer
	user        string
	directoryID string
	builder     *users.ItemMailFoldersItemMessagesDeltaRequestBuilder
	options     *users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func getMailDeltaBuilder(
	ctx context.Context,
	gs graph.Servicer,
	user string,
	directoryID string,
	options *users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration,
) *users.ItemMailFoldersItemMessagesDeltaRequestBuilder {
	builder := gs.Client().
		Users().
		ByUserId(user).
		MailFolders().
		ByMailFolderId(directoryID).
		Messages().
		Delta()

	return builder
}

func NewMailDeltaPager(
	ctx context.Context,
	gs graph.Servicer,
	user, directoryID, oldDelta string,
	immutableIDs bool,
) itemPager {
	queryParams := &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{
			Select: []string{"id", "isRead"},
		},
		Headers: buildPreferHeaders(true, immutableIDs),
	}

	var builder *users.ItemMailFoldersItemMessagesDeltaRequestBuilder

	if len(oldDelta) > 0 {
		builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(oldDelta, gs.Adapter())
	} else {
		builder = getMailDeltaBuilder(ctx, gs, user, directoryID, queryParams)
	}

	return &mailDeltaPager{gs, user, directoryID, builder, queryParams}
}

func (p *mailDeltaPager) getPage(ctx context.Context) (api.DeltaPageLinker, error) {
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
	p.builder = p.gs.Client().
		Users().
		ByUserId(p.user).
		MailFolders().
		ByMailFolderId(p.directoryID).
		Messages().
		Delta()
}

func (p *mailDeltaPager) valuesIn(pl api.PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Messageable](pl)
}

func (c Mail) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	user, directoryID, oldDelta string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.Service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	ctx = clues.Add(
		ctx,
		"category", selectors.ExchangeMail,
		"container_id", directoryID)

	pager := NewMailPager(ctx, service, user, directoryID, immutableIDs)
	deltaPager := NewMailDeltaPager(ctx, service, user, directoryID, oldDelta, immutableIDs)

	return getAddedAndRemovedItemIDs(ctx, service, pager, deltaPager, oldDelta, canMakeDeltaQueries)
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

// Serialize transforms the mail item into a byte slice.
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

	var (
		err    error
		writer = kjson.NewJsonSerializationWriter()
	)

	defer writer.Close()

	if err = writer.WriteObjectValue("", msg); err != nil {
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
