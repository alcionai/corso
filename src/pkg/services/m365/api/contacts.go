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

	mdl, err := c.Stable.Client().Users().ByUserId(user).ContactFolders().Post(ctx, requestBody, nil)
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

	err = srv.Client().Users().ByUserId(user).ContactFolders().ByContactFolderId(folderID).Delete(ctx, nil)
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
		Headers: newPreferHeaders(preferImmutableIDs(immutableIDs)),
	}

	cont, err := c.Stable.Client().Users().ByUserId(user).Contacts().ByContactId(itemID).Get(ctx, options)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	return cont, ContactInfo(cont), nil
}

func (c Contacts) GetContainerByID(
	ctx context.Context,
	userID, dirID string,
) (graph.Container, error) {
	config := &users.ItemContactFoldersContactFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersContactFolderItemRequestBuilderGetQueryParameters{
			Select: idAnd(displayName, parentFolderID),
		},
	}

	resp, err := c.Stable.Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(dirID).
		Get(ctx, config)
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
	fn func(graph.CachedContainer) error,
	errs *fault.Bus,
) error {
	service, err := c.Service()
	if err != nil {
		return graph.Stack(ctx, err)
	}

	config := &users.ItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersItemChildFoldersRequestBuilderGetQueryParameters{
			Select: idAnd(displayName, parentFolderID),
		},
	}

	el := errs.Local()
	builder := service.Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(baseDirID).
		ChildFolders()

	for {
		if el.Failure() != nil {
			break
		}

		resp, err := builder.Get(ctx, config)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		for _, fold := range resp.GetValue() {
			if el.Failure() != nil {
				return el.Failure()
			}

			if err := graph.CheckIDNameAndParentFolderID(fold); err != nil {
				errs.AddRecoverable(graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			fctx := clues.Add(
				ctx,
				"container_id", ptr.Val(fold.GetId()),
				"container_display_name", ptr.Val(fold.GetDisplayName()))

			temp := graph.NewCacheFolder(fold, nil, nil)
			if err := fn(&temp); err != nil {
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
	gs      graph.Servicer
	builder *users.ItemContactFoldersItemContactsRequestBuilder
	options *users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration
}

func NewContactPager(
	ctx context.Context,
	gs graph.Servicer,
	user, directoryID string,
	immutableIDs bool,
) itemPager {
	config := &users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersItemContactsRequestBuilderGetQueryParameters{
			Select: idAnd(parentFolderID),
		},
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	builder := gs.Client().
		Users().
		ByUserId(user).
		ContactFolders().
		ByContactFolderId(directoryID).
		Contacts()

	return &contactPager{gs, builder, config}
}

func (p *contactPager) getPage(ctx context.Context) (api.DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return api.EmptyDeltaLinker[models.Contactable]{PageLinkValuer: resp}, nil
}

func (p *contactPager) setNext(nextLink string) {
	p.builder = users.NewItemContactFoldersItemContactsRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't need reset
func (p *contactPager) reset(context.Context) {}

func (p *contactPager) valuesIn(pl api.PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Contactable](pl)
}

// ---------------------------------------------------------------------------
// delta item pager
// ---------------------------------------------------------------------------

var _ itemPager = &contactDeltaPager{}

type contactDeltaPager struct {
	gs          graph.Servicer
	user        string
	directoryID string
	builder     *users.ItemContactFoldersItemContactsDeltaRequestBuilder
	options     *users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration
}

func getContactDeltaBuilder(
	ctx context.Context,
	gs graph.Servicer,
	user string,
	directoryID string,
	options *users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration,
) *users.ItemContactFoldersItemContactsDeltaRequestBuilder {
	builder := gs.Client().Users().ByUserId(user).ContactFolders().ByContactFolderId(directoryID).Contacts().Delta()
	return builder
}

func NewContactDeltaPager(
	ctx context.Context,
	gs graph.Servicer,
	user, directoryID, deltaURL string,
	immutableIDs bool,
) itemPager {
	options := &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetQueryParameters{
			Select: idAnd(parentFolderID),
		},
		Headers: newPreferHeaders(preferPageSize(maxDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	var builder *users.ItemContactFoldersItemContactsDeltaRequestBuilder
	if deltaURL != "" {
		builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(deltaURL, gs.Adapter())
	} else {
		builder = getContactDeltaBuilder(ctx, gs, user, directoryID, options)
	}

	return &contactDeltaPager{gs, user, directoryID, builder, options}
}

func (p *contactDeltaPager) getPage(ctx context.Context) (api.DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *contactDeltaPager) setNext(nextLink string) {
	p.builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *contactDeltaPager) reset(ctx context.Context) {
	p.builder = getContactDeltaBuilder(ctx, p.gs, p.user, p.directoryID, p.options)
}

func (p *contactDeltaPager) valuesIn(pl api.PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Contactable](pl)
}

func (c Contacts) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	user, directoryID, oldDelta string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.Service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Stack(ctx, err)
	}

	ctx = clues.Add(
		ctx,
		"category", selectors.ExchangeContact,
		"container_id", directoryID)

	pager := NewContactPager(ctx, service, user, directoryID, immutableIDs)
	deltaPager := NewContactDeltaPager(ctx, service, user, directoryID, oldDelta, immutableIDs)

	return getAddedAndRemovedItemIDs(ctx, service, pager, deltaPager, oldDelta, canMakeDeltaQueries)
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
