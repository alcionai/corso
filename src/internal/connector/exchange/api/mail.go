package api

import (
	"context"
	"fmt"
	"os"

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
		return nil, graph.Wrap(ctx, err, "creating mail folder")
	}

	return mdl, nil
}

func (c Mail) CreateMailFolderWithParent(
	ctx context.Context,
	user, folder, parentID string,
) (models.MailFolderable, error) {
	service, err := c.service()
	if err != nil {
		return nil, graph.Stack(ctx, err)
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
	srv, err := newService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.Client().UsersById(user).MailFoldersById(folderID).Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

func (c Mail) GetContainerByID(
	ctx context.Context,
	userID, dirID string,
) (graph.Container, error) {
	service, err := c.service()
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	ofmf, err := optionsForMailFoldersItem([]string{"displayName", "parentFolderId"})
	if err != nil {
		return nil, graph.Wrap(ctx, err, "setting mail folder options")
	}

	resp, err := service.Client().UsersById(userID).MailFoldersById(dirID).Get(ctx, ofmf)
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
	errs *fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	mail, err := c.stable.Client().UsersById(user).MessagesById(itemID).Get(ctx, nil)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	if ptr.Val(mail.GetHasAttachments()) || HasAttachments(mail.GetBody()) {
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
			return nil, nil, graph.Wrap(ctx, err, "mail attachment download")
		}

		mail.SetAttachments(attached.GetValue())
	}

	return mail, MailInfo(mail), nil
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
	service, err := c.service()
	if err != nil {
		return graph.Stack(ctx, err)
	}

	el := errs.Local()
	builder := service.Client().
		UsersById(userID).
		MailFolders().
		Delta()

	for {
		if el.Failure() != nil {
			break
		}

		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		for _, v := range resp.GetValue() {
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

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemMailFoldersDeltaRequestBuilder(link, service.Adapter())
	}

	return el.Failure()
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
		return nil, graph.Stack(ctx, err)
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

	ctx = clues.Add(
		ctx,
		"category", selectors.ExchangeMail,
		"container_id", directoryID)

	options, err := optionsForFolderMessagesDelta([]string{"isRead"})
	if err != nil {
		return nil,
			nil,
			DeltaUpdate{},
			graph.Wrap(ctx, err, "setting contact folder options")
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

	if len(os.Getenv("CORSO_URL_LOGGING")) > 0 {
		gri, err := builder.ToGetRequestInformation(ctx, options)
		if err != nil {
			logger.CtxErr(ctx, err).Error("getting builder info")
		} else {
			logger.Ctx(ctx).
				Infow("builder path-parameters", "path_parameters", gri.PathParameters)
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

func MailInfo(msg models.Messageable) *details.ExchangeInfo {
	var (
		sender     = graph.UnwrapEmailAddress(msg.GetSender())
		subject    = ptr.Val(msg.GetSubject())
		received   = ptr.Val(msg.GetReceivedDateTime())
		created    = ptr.Val(msg.GetCreatedDateTime())
		recipients = make([]string, 0)
	)

	if msg.GetToRecipients() != nil {
		ppl := msg.GetToRecipients()
		for _, entry := range ppl {
			temp := graph.UnwrapEmailAddress(entry)
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
		Created:   created,
		Modified:  ptr.OrNow(msg.GetLastModifiedDateTime()),
	}
}
