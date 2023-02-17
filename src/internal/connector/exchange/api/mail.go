package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kioser "github.com/microsoft/kiota-serialization-json-go"
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

	mdl, err := c.stable.Client().UsersById(user).MailFolders().Post(ctx, requestBody, nil)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return mdl, nil
}

func (c Mail) CreateMailFolderWithParent(
	ctx context.Context,
	user, folder, parentID string,
) (models.MailFolderable, error) {
	service, err := c.service()
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	isHidden := false
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	requestBody.SetIsHidden(&isHidden)

	mdl, err := service.
		Client().
		UsersById(user).
		MailFoldersById(parentID).
		ChildFolders().
		Post(ctx, requestBody, nil)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return mdl, nil
}

// DeleteContainer removes a mail folder with the corresponding M365 ID  from the user's M365 Exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/mailfolder-delete?view=graph-rest-1.0&tabs=http
func (c Mail) DeleteContainer(
	ctx context.Context,
	user, folderID string,
) error {
	err := c.stable.Client().UsersById(user).MailFoldersById(folderID).Delete(ctx, nil)
	if err != nil {
		return clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return nil
}

func (c Mail) GetContainerByID(
	ctx context.Context,
	userID, dirID string,
) (graph.Container, error) {
	service, err := c.service()
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	ofmf, err := optionsForMailFoldersItem([]string{"displayName", "parentFolderId"})
	if err != nil {
		return nil, clues.Wrap(err, "setting mail folder options").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	resp, err := service.Client().UsersById(userID).MailFoldersById(dirID).Get(ctx, ofmf)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return resp, nil
}

// GetItem retrieves a Messageable item.  If the item contains an attachment, that
// attachment is also downloaded.
func (c Mail) GetItem(
	ctx context.Context,
	user, itemID string,
	errs *fault.Errors,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	mail, err := c.stable.Client().UsersById(user).MessagesById(itemID).Get(ctx, nil)
	if err != nil {
		return nil, nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	if *mail.GetHasAttachments() || HasAttachments(mail.GetBody()) {
		options := &users.ItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMessagesItemAttachmentsRequestBuilderGetQueryParameters{
				Expand: []string{"microsoft.graph.itemattachment/item"},
			},
		}

		attached, err := c.largeItem.
			Client().
			UsersById(user).
			MessagesById(itemID).
			Attachments().
			Get(ctx, options)
		if err != nil {
			return nil, nil, clues.Wrap(err, "mail attachment download").WithClues(ctx).WithAll(graph.ErrData(err)...)
		}

		mail.SetAttachments(attached.GetValue())
	}

	return mail, MailInfo(mail), nil
}

// EnumerateContainers iterates through all of the users current
// mail folders, converting each to a graph.CacheFolder, and calling
// fn(cf) on each one.  If fn(cf) errors, the error is aggregated
// into a multierror that gets returned to the caller.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Mail) EnumerateContainers(
	ctx context.Context,
	userID, baseDirID string,
	fn func(graph.CacheFolder) error,
	errs *fault.Errors,
) error {
	service, err := c.service()
	if err != nil {
		return clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	builder := service.Client().
		UsersById(userID).
		MailFolders().
		Delta()

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
		}

		for _, v := range resp.GetValue() {
			fctx := clues.AddAll(
				ctx,
				"container_id", ptr.Val(v.GetId()),
				"container_name", ptr.Val(v.GetDisplayName()))

			temp := graph.NewCacheFolder(v, nil, nil)
			if err := fn(temp); err != nil {
				errs.Add(clues.Stack(err).WithClues(fctx).WithAll(graph.ErrData(err)...))
				continue
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemMailFoldersDeltaRequestBuilder(link, service.Adapter())
	}

	return errs.Err()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager = &mailPager{}

type mailPager struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersItemMessagesDeltaRequestBuilder
	options *users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func (p *mailPager) getPage(ctx context.Context) (api.DeltaPageLinker, error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return page, nil
}

func (p *mailPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *mailPager) valuesIn(pl api.DeltaPageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Messageable](pl)
}

func (c Mail) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	user, directoryID, oldDelta string,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	var (
		deltaURL   string
		resetDelta bool
	)

	ctx = clues.AddAll(
		ctx,
		"category", selectors.ExchangeMail,
		"folder_id", directoryID)

	options, err := optionsForFolderMessagesDelta([]string{"isRead"})
	if err != nil {
		return nil,
			nil,
			DeltaUpdate{},
			clues.Wrap(err, "setting contact folder options").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	if len(oldDelta) > 0 {
		var (
			builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(oldDelta, service.Adapter())
			pgr     = &mailPager{service, builder, options}
		)

		added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
		// note: happy path, not the error condition
		if err == nil {
			return added, removed, DeltaUpdate{deltaURL, false}, err
		}
		// only return on error if it is NOT a delta issue.
		// on bad deltas we retry the call with the regular builder
		if !graph.IsErrInvalidDelta(err) {
			return nil, nil, DeltaUpdate{}, err
		}

		resetDelta = true
	}

	builder := service.Client().UsersById(user).MailFoldersById(directoryID).Messages().Delta()
	pgr := &mailPager{service, builder, options}

	gri, err := builder.ToGetRequestInformation(ctx, options)
	if err != nil {
		logger.Ctx(ctx).Errorw("getting builder info", "error", err)
	} else {
		uri, err := gri.GetUri()
		if err != nil {
			logger.Ctx(ctx).Errorw("getting builder uri", "error", err)
		} else {
			logger.Ctx(ctx).Infow("mail builder", "user", user, "directoryID", directoryID, "uri", uri)
		}
	}

	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return added, removed, DeltaUpdate{deltaURL, resetDelta}, nil
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
		return nil, clues.Wrap(fmt.Errorf("parseable type: %T", item), "parsable is not a Messageable")
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(msg.GetId()))

	var (
		err    error
		writer = kioser.NewJsonSerializationWriter()
	)

	defer writer.Close()

	if err = writer.WriteObjectValue("", msg); err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, clues.Wrap(err, "serializing email").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return bs, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func MailInfo(msg models.Messageable) *details.ExchangeInfo {
	sender := ""
	subject := ptr.Val(msg.GetSubject())
	received := ptr.Val(msg.GetReceivedDateTime())
	created := ptr.Val(msg.GetCreatedDateTime())

	if msg.GetSender() != nil &&
		msg.GetSender().GetEmailAddress() != nil {
		sender = ptr.Val(msg.GetSender().GetEmailAddress().GetAddress())
	}

	return &details.ExchangeInfo{
		ItemType: details.ExchangeMail,
		Sender:   sender,
		Subject:  subject,
		Received: received,
		Created:  created,
		Modified: ptr.OrNow(msg.GetLastModifiedDateTime()),
	}
}
