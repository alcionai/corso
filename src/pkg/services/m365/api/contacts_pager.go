package api

import (
	"context"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// container pager
// ---------------------------------------------------------------------------

var _ Pager[models.ContactFolderable] = &contactsFoldersPageCtrl{}

type contactsFoldersPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemContactFoldersItemChildFoldersRequestBuilder
	options *users.ItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration
}

func (c Contacts) NewContactFoldersPager(
	userID, baseContainerID string,
	immutableIDs bool,
	selectProps ...string,
) Pager[models.ContactFolderable] {
	options := &users.ItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
		QueryParameters: &users.ItemContactFoldersItemChildFoldersRequestBuilderGetQueryParameters{},
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
		ByContactFolderId(baseContainerID).
		ChildFolders()

	return &contactsFoldersPageCtrl{c.Stable, builder, options}
}

func (p *contactsFoldersPageCtrl) GetPage(
	ctx context.Context,
) (NextLinkValuer[models.ContactFolderable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *contactsFoldersPageCtrl) SetNextLink(nextLink string) {
	p.builder = users.NewItemContactFoldersItemChildFoldersRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *contactsFoldersPageCtrl) ValidModTimes() bool {
	return true
}

// EnumerateContainers iterates through all of the users current
// contacts folders, transforming each to a graph.CacheFolder, and calling
// fn(cf).
// Contact folders are represented in their current state, and do
// not contain historical data.
func (c Contacts) EnumerateContainers(
	ctx context.Context,
	userID, baseContainerID string,
	immutableIDs bool,
	fn func(graph.CachedContainer) error,
	errs *fault.Bus,
) error {
	var (
		el  = errs.Local()
		pgr = c.NewContactFoldersPager(userID, baseContainerID, immutableIDs)
	)

	containers, err := batchEnumerateItems(ctx, pgr)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	for _, c := range containers {
		if el.Failure() != nil {
			break
		}

		gncf := graph.NewCacheFolder(c, nil, nil)

		if err := fn(&gncf); err != nil {
			errs.AddRecoverable(ctx, graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
			continue
		}
	}

	return el.Failure()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ Pager[models.Contactable] = &contactsPageCtrl{}

type contactsPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemContactFoldersItemContactsRequestBuilder
	options *users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration
}

func (c Contacts) NewContactsPager(
	userID, containerID string,
	immutableIDs bool,
	selectProps ...string,
) Pager[models.Contactable] {
	options := &users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
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

func (p *contactsPageCtrl) GetPage(
	ctx context.Context,
) (NextLinkValuer[models.Contactable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *contactsPageCtrl) SetNextLink(nextLink string) {
	p.builder = users.NewItemContactFoldersItemContactsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *contactsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Contacts) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	userID, containerID string,
) (map[string]string, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewContactsPager(userID, containerID, false, contactCollisionKeyProps()...)

	items, err := batchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating contacts")
	}

	m := map[string]string{}

	for _, item := range items {
		m[ContactCollisionKey(item)] = ptr.Val(item.GetId())
	}

	return m, nil
}

func (c Contacts) GetItemIDsInContainer(
	ctx context.Context,
	userID, containerID string,
) (map[string]struct{}, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewContactsPager(userID, containerID, false, idAnd()...)

	items, err := batchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating contacts")
	}

	m := map[string]struct{}{}

	for _, item := range items {
		m[ptr.Val(item.GetId())] = struct{}{}
	}

	return m, nil
}

// ---------------------------------------------------------------------------
// delta item ID pager
// ---------------------------------------------------------------------------

var _ DeltaPager[models.Contactable] = &contactDeltaPager{}

type contactDeltaPager struct {
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
) *users.ItemContactFoldersItemContactsDeltaRequestBuilder {
	builder := gs.Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(containerID).
		Contacts().
		Delta()

	return builder
}

func (c Contacts) NewContactsDeltaPager(
	ctx context.Context,
	userID, containerID, prevDeltaLink string,
	immutableIDs bool,
	selectProps ...string,
) DeltaPager[models.Contactable] {
	options := &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration{
		// do NOT set Top.  It limits the total items received.
		QueryParameters: &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(c.options.DeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	var builder *users.ItemContactFoldersItemContactsDeltaRequestBuilder
	if len(prevDeltaLink) > 0 {
		builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(prevDeltaLink, c.Stable.Adapter())
	} else {
		builder = getContactDeltaBuilder(ctx, c.Stable, userID, containerID)
	}

	return &contactDeltaPager{c.Stable, userID, containerID, builder, options}
}

func (p *contactDeltaPager) GetPage(
	ctx context.Context,
) (DeltaLinkValuer[models.Contactable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *contactDeltaPager) SetNextLink(nextLink string) {
	p.builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *contactDeltaPager) Reset(ctx context.Context) {
	p.builder = getContactDeltaBuilder(ctx, p.gs, p.userID, p.containerID)
}

func (p *contactDeltaPager) ValidModTimes() bool {
	return true
}

func (c Contacts) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, containerID, prevDeltaLink string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) (map[string]time.Time, bool, []string, DeltaUpdate, error) {
	ctx = clues.Add(
		ctx,
		"data_category", path.ContactsCategory,
		"container_id", containerID)

	deltaPager := c.NewContactsDeltaPager(
		ctx,
		userID,
		containerID,
		prevDeltaLink,
		immutableIDs,
		idAnd(lastModifiedDateTime)...)
	pager := c.NewContactsPager(
		userID,
		containerID,
		immutableIDs,
		idAnd(lastModifiedDateTime)...)

	return getAddedAndRemovedItemIDs[models.Contactable](
		ctx,
		pager,
		deltaPager,
		prevDeltaLink,
		canMakeDeltaQueries,
		addedAndRemovedByAddtlData[models.Contactable])
}
