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

func (c Client) Contacts() Contacts {
	return Contacts{c}
}

// Contacts is an interface-compliant provider of the client.
type Contacts struct {
	Client
}

// ---------------------------------------------------------------------------
// methods
// ---------------------------------------------------------------------------

// CreateContactFolder makes a contact folder with the displayName of folderName.
// If successful, returns the created folder object.
func (c Contacts) CreateContactFolder(
	ctx context.Context,
	user, folderName string,
) (models.ContactFolderable, error) {
	requestBody := models.NewContactFolder()
	temp := folderName
	requestBody.SetDisplayName(&temp)

	mdl, err := c.Stable.Client().UsersById(user).ContactFolders().Post(ctx, requestBody, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating contact folder")
	}

	return mdl, nil
}

// DeleteContainer deletes the ContactFolder associated with the M365 ID if permissions are valid.
func (c Contacts) DeleteContainer(
	ctx context.Context,
	user, folderID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.Client().UsersById(user).ContactFoldersById(folderID).Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

// GetItem retrieves a Contactable item.
func (c Contacts) GetItem(
	ctx context.Context,
	user, itemID string,
	immutableIDs bool,
	_ *fault.Bus, // no attachments to iterate over, so this goes unused
) (serialization.Parsable, *details.ExchangeInfo, error) {
	options := &users.ItemContactsContactItemRequestBuilderGetRequestConfiguration{
		Headers: buildPreferHeaders(false, immutableIDs),
	}

	cont, err := c.Stable.Client().UsersById(user).ContactsById(itemID).Get(ctx, options)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	return cont, ContactInfo(cont), nil
}

func (c Contacts) GetContainerByID(
	ctx context.Context,
	userID, dirID string,
) (graph.Container, error) {
	ofcf, err := optionsForContactFolderByID([]string{"displayName", "parentFolderId"})
	if err != nil {
		return nil, graph.Wrap(ctx, err, "setting contact folder options")
	}

	resp, err := c.Stable.Client().UsersById(userID).ContactFoldersById(dirID).Get(ctx, ofcf)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

// EnumerateContainers iterates through all of the users current
// contacts folders, converting each to a graph.CacheFolder, and calling
// fn(cf) on each one.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Contacts) EnumerateContainers(
	ctx context.Context,
	userID, baseDirID string,
	fn func(graph.CacheFolder) error,
	errs *fault.Bus,
) error {
	service, err := c.service()
	if err != nil {
		return graph.Stack(ctx, err)
	}

	fields := []string{"displayName", "parentFolderId"}

	ofcf, err := optionsForContactChildFolders(fields)
	if err != nil {
		return graph.Wrap(ctx, err, "setting contact child folder options")
	}

	el := errs.Local()
	builder := service.Client().
		UsersById(userID).
		ContactFoldersById(baseDirID).
		ChildFolders()

	for {
		if el.Failure() != nil {
			break
		}

		resp, err := builder.Get(ctx, ofcf)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		for _, fold := range resp.GetValue() {
			if el.Failure() != nil {
				return el.Failure()
			}

			if err := checkIDAndName(fold); err != nil {
				errs.AddRecoverable(graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			fctx := clues.Add(
				ctx,
				"container_id", ptr.Val(fold.GetId()),
				"container_display_name", ptr.Val(fold.GetDisplayName()))

			temp := graph.NewCacheFolder(fold, nil, nil)
			if err := fn(temp); err != nil {
				errs.AddRecoverable(graph.Stack(fctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemContactFoldersItemChildFoldersRequestBuilder(link, service.Adapter())
	}

	return el.Failure()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager = &contactPager{}

type contactPager struct {
	gs          graph.Servicer
	user        string
	directoryID string

	// switch between using delta vs non-delta fetch
	nonDelta bool

	builder *users.ItemContactFoldersItemContactsRequestBuilder
	options *users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration

	deltaBuilder *users.ItemContactFoldersItemContactsDeltaRequestBuilder
	deltaOptions *users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration
}

func NewContactPager(
	gs graph.Servicer,
	user, directoryID, deltaURL string,
	nonDelta bool,
	immutableIDs bool,
) (contactPager, error) {
	options, err := optionsForContactFoldersItem([]string{"parentFolderId"}, immutableIDs)
	if err != nil {
		return contactPager{}, err
	}

	deltaOptions, err := optionsForContactFoldersItemDelta([]string{"parentFolderId"}, immutableIDs)
	if err != nil {
		return contactPager{}, err
	}

	deltaBuilder := gs.Client().UsersById(user).ContactFoldersById(directoryID).Contacts().Delta()
	if deltaURL != "" {
		deltaBuilder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(deltaURL, gs.Adapter())
	}

	return contactPager{
		gs:          gs,
		user:        user,
		directoryID: directoryID,
		nonDelta:    nonDelta,

		// for non-delta based pagination
		builder: gs.Client().UsersById(user).ContactFoldersById(directoryID).Contacts(),
		options: options,

		// for delta token based pagination
		deltaBuilder: deltaBuilder,
		deltaOptions: deltaOptions,
	}, nil
}

func (p *contactPager) getNextPageDelta(ctx context.Context) ([]getIDAndAddtler, bool, string, error) {
	page, err := p.deltaBuilder.Get(ctx, p.deltaOptions)
	if err != nil {
		return nil, false, "", graph.Stack(ctx, err)
	}

	nextLink, deltaLink := api.NextAndDeltaLink(page)
	if len(os.Getenv("CORSO_URL_LOGGING")) > 0 {
		if !api.IsNextLinkValid(nextLink) || api.IsNextLinkValid(deltaLink) {
			logger.Ctx(ctx).Infof("Received invalid link from M365:\nNext Link: %s\nDelta Link: %s\n", nextLink, deltaLink)
		}
	}

	p.deltaBuilder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(nextLink, p.gs.Adapter())

	items, err := toValues[models.Contactable](page)
	if err != nil {
		return nil, false, "", err
	}

	return items, nextLink == "", deltaLink, nil
}

func (p *contactPager) getNextPageNonDelta(ctx context.Context) ([]getIDAndAddtler, bool, string, error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, false, "", graph.Stack(ctx, err)
	}

	nextLink := api.NextLink(page)
	if len(os.Getenv("CORSO_URL_LOGGING")) > 0 {
		if !api.IsNextLinkValid(nextLink) {
			logger.Ctx(ctx).Infof("Received invalid link from M365:\nNext Link: %s\n", nextLink)
		}
	}

	p.builder = users.NewItemContactFoldersItemContactsRequestBuilder(nextLink, p.gs.Adapter())

	items, err := toValues[models.Contactable](page)
	if err != nil {
		return nil, false, "", err
	}

	return items, nextLink == "", "", nil
}

func (p *contactPager) getNextPage(ctx context.Context) ([]getIDAndAddtler, bool, string, error) {
	if p.nonDelta {
		return p.getNextPageNonDelta(ctx)
	}

	return p.getNextPageDelta(ctx)
}

func (p *contactPager) valuesIn(pl api.DeltaPageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Messageable](pl)
}

func (p *contactPager) reset(nonDelta bool) {
	if nonDelta {
		p.nonDelta = true
		return
	}

	p.deltaBuilder = p.gs.Client().UsersById(p.user).ContactFoldersById(p.directoryID).Contacts().Delta()
}

func (c Contacts) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	user, directoryID, oldDelta string,
	immutableIDs bool,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Stack(ctx, err)
	}

	ctx = clues.Add(
		ctx,
		"category", selectors.ExchangeContact,
		"container_id", directoryID)

	// TODO(meain): Check if exchange if full here and start with non-delta if possible
	pgr, err := NewContactPager(service, user, directoryID, oldDelta, false, immutableIDs)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return GetAddedAndRemovedItemIDsFromPager(ctx, oldDelta, &pgr)
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

// Serialize rserializes the item into a byte slice.
func (c Contacts) Serialize(
	ctx context.Context,
	item serialization.Parsable,
	user, itemID string,
) ([]byte, error) {
	contact, ok := item.(models.Contactable)
	if !ok {
		return nil, clues.New(fmt.Sprintf("item is not a Contactable: %T", item))
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(contact.GetId()))

	var (
		err    error
		writer = kjson.NewJsonSerializationWriter()
	)

	defer writer.Close()

	if err = writer.WriteObjectValue("", contact); err != nil {
		return nil, graph.Stack(ctx, err)
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, graph.Wrap(ctx, err, "serializing contact")
	}

	return bs, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func ContactInfo(contact models.Contactable) *details.ExchangeInfo {
	name := ptr.Val(contact.GetDisplayName())
	created := ptr.Val(contact.GetCreatedDateTime())

	return &details.ExchangeInfo{
		ItemType:    details.ExchangeContact,
		ContactName: name,
		Created:     created,
		Modified:    ptr.OrNow(contact.GetLastModifiedDateTime()),
	}
}
