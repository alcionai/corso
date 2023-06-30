package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// container pager
// ---------------------------------------------------------------------------

// EnumerateContainers iterates through all of the users current
// contacts folders, converting each to a graph.CacheFolder, and calling
// fn(cf) on each one.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Contacts) EnumerateContainers(
	ctx context.Context,
	userID, baseContainerID string,
	fn func(graph.CachedContainer) error,
	errs *fault.Bus,
) error {
	config := &users.ItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersItemChildFoldersRequestBuilderGetQueryParameters{
			Select: idAnd(displayName, parentFolderID),
		},
	}

	el := errs.Local()
	builder := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(baseContainerID).
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
				errs.AddRecoverable(ctx, graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			fctx := clues.Add(
				ctx,
				"container_id", ptr.Val(fold.GetId()),
				"container_display_name", ptr.Val(fold.GetDisplayName()))

			temp := graph.NewCacheFolder(fold, nil, nil)
			if err := fn(&temp); err != nil {
				errs.AddRecoverable(ctx, graph.Stack(fctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemContactFoldersItemChildFoldersRequestBuilder(link, c.Stable.Adapter())
	}

	return el.Failure()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager[models.Contactable] = &contactsPageCtrl{}

type contactsPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemContactFoldersItemContactsRequestBuilder
	options *users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration
}

func (c Contacts) NewContactsPager(
	userID, containerID string,
	selectProps ...string,
) itemPager[models.Contactable] {
	options := &users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
		QueryParameters: &users.ItemContactFoldersItemContactsRequestBuilderGetQueryParameters{},
		// do NOT set Top.  It limits the total items received.
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	builder := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(containerID).
		Contacts()

	return &contactsPageCtrl{c.Stable, builder, options}
}

//lint:ignore U1000 False Positive
func (p *contactsPageCtrl) getPage(ctx context.Context) (PageLinkValuer[models.Contactable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.Contactable]{PageLinkValuer: resp}, nil
}

//lint:ignore U1000 False Positive
func (p *contactsPageCtrl) setNext(nextLink string) {
	p.builder = users.NewItemContactFoldersItemContactsRequestBuilder(nextLink, p.gs.Adapter())
}

//lint:ignore U1000 False Positive
func (c Contacts) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	userID, containerID string,
) (map[string]string, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewContactsPager(userID, containerID, contactCollisionKeyProps()...)

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating contacts")
	}

	m := map[string]string{}

	for _, item := range items {
		m[ContactCollisionKey(item)] = ptr.Val(item.GetId())
	}

	return m, nil
}

// ---------------------------------------------------------------------------
// item ID pager
// ---------------------------------------------------------------------------

var _ itemIDPager = &contactIDPager{}

type contactIDPager struct {
	gs      graph.Servicer
	builder *users.ItemContactFoldersItemContactsRequestBuilder
	options *users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration
}

func (c Contacts) NewContactIDsPager(
	ctx context.Context,
	userID, containerID string,
	immutableIDs bool,
) itemIDPager {
	config := &users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersItemContactsRequestBuilderGetQueryParameters{
			Select: idAnd(parentFolderID),
			// do NOT set Top.  It limits the total items received.
		},
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	builder := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(containerID).
		Contacts()

	return &contactIDPager{c.Stable, builder, config}
}

func (p *contactIDPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.Contactable]{PageLinkValuer: resp}, nil
}

func (p *contactIDPager) setNext(nextLink string) {
	p.builder = users.NewItemContactFoldersItemContactsRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't need reset
func (p *contactIDPager) reset(context.Context) {}

func (p *contactIDPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Contactable](pl)
}

// ---------------------------------------------------------------------------
// delta item ID pager
// ---------------------------------------------------------------------------

var _ itemIDPager = &contactDeltaIDPager{}

type contactDeltaIDPager struct {
	gs          graph.Servicer
	userID      string
	containerID string
	builder     *users.ItemContactFoldersItemContactsDeltaRequestBuilder
	options     *users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration
}

func getContactDeltaBuilder(
	ctx context.Context,
	gs graph.Servicer,
	userID, containerID string,
	options *users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration,
) *users.ItemContactFoldersItemContactsDeltaRequestBuilder {
	builder := gs.Client().Users().ByUserId(userID).ContactFolders().ByContactFolderId(containerID).Contacts().Delta()
	return builder
}

func (c Contacts) NewContactDeltaIDsPager(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
) itemIDPager {
	options := &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetQueryParameters{
			Select: idAnd(parentFolderID),
			// do NOT set Top.  It limits the total items received.
		},
		Headers: newPreferHeaders(preferPageSize(maxDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	var builder *users.ItemContactFoldersItemContactsDeltaRequestBuilder
	if oldDelta != "" {
		builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(oldDelta, c.Stable.Adapter())
	} else {
		builder = getContactDeltaBuilder(ctx, c.Stable, userID, containerID, options)
	}

	return &contactDeltaIDPager{c.Stable, userID, containerID, builder, options}
}

func (p *contactDeltaIDPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *contactDeltaIDPager) setNext(nextLink string) {
	p.builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *contactDeltaIDPager) reset(ctx context.Context) {
	p.builder = getContactDeltaBuilder(ctx, p.gs, p.userID, p.containerID, p.options)
}

func (p *contactDeltaIDPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Contactable](pl)
}

func (c Contacts) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	ctx = clues.Add(
		ctx,
		"category", selectors.ExchangeContact,
		"container_id", containerID)

	pager := c.NewContactIDsPager(ctx, userID, containerID, immutableIDs)
	deltaPager := c.NewContactDeltaIDsPager(ctx, userID, containerID, oldDelta, immutableIDs)

	return getAddedAndRemovedItemIDs(ctx, c.Stable, pager, deltaPager, oldDelta, canMakeDeltaQueries)
}
