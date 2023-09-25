package api

import (
	"context"
	"fmt"
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

var _ Pager[models.MailFolderable] = &mailFoldersPageCtrl{}

type mailFoldersPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersRequestBuilder
	options *users.ItemMailFoldersRequestBuilderGetRequestConfiguration
}

func (c Mail) NewMailFoldersPager(
	userID string,
	immutableIDs bool,
	selectProps ...string,
) Pager[models.MailFolderable] {
	options := &users.ItemMailFoldersRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
		QueryParameters: &users.ItemMailFoldersRequestBuilderGetQueryParameters{},
		// do NOT set Top.  It limits the total items received.
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	// v1.0 non delta /mailFolders endpoint does not return any of the nested folders
	rawURL := fmt.Sprintf(mailFoldersBetaURLTemplate, userID)
	builder := users.NewItemMailFoldersRequestBuilder(rawURL, c.Stable.Adapter())

	return &mailFoldersPageCtrl{c.Stable, builder, options}
}

func (p *mailFoldersPageCtrl) GetPage(
	ctx context.Context,
) (NextLinkValuer[models.MailFolderable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *mailFoldersPageCtrl) SetNextLink(nextLink string) {
	p.builder = users.NewItemMailFoldersRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *mailFoldersPageCtrl) ValidModTimes() bool {
	return true
}

// EnumerateContainers iterates through all of the users current
// mail folders, transforming each to a graph.CacheFolder, and calling
// fn(cf).
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Mail) EnumerateContainers(
	ctx context.Context,
	userID, _ string, // baseContainerID not needed here
	immutableIDs bool,
	fn func(graph.CachedContainer) error,
	errs *fault.Bus,
) error {
	var (
		el  = errs.Local()
		pgr = c.NewMailFoldersPager(userID, immutableIDs)
	)

	containers, err := enumerateItems(ctx, pgr)
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

var _ Pager[models.Messageable] = &mailsPageCtrl{}

type mailsPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersItemMessagesRequestBuilder
	options *users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration
}

func (c Mail) NewMailPager(
	userID, containerID string,
	immutableIDs bool,
	selectProps ...string,
) Pager[models.Messageable] {
	options := &users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
		QueryParameters: &users.ItemMailFoldersItemMessagesRequestBuilderGetQueryParameters{},
		// do NOT set Top.  It limits the total items received.
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	builder := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Messages()

	return &mailsPageCtrl{c.Stable, builder, options}
}

func (p *mailsPageCtrl) GetPage(
	ctx context.Context,
) (NextLinkValuer[models.Messageable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *mailsPageCtrl) SetNextLink(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *mailsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Mail) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	userID, containerID string,
) (map[string]string, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewMailPager(userID, containerID, false, mailCollisionKeyProps()...)

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating mails")
	}

	m := map[string]string{}

	for _, item := range items {
		m[MailCollisionKey(item)] = ptr.Val(item.GetId())
	}

	return m, nil
}

func (c Mail) GetItemIDsInContainer(
	ctx context.Context,
	userID, containerID string,
) (map[string]struct{}, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewMailPager(userID, containerID, false, idAnd()...)

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating mails")
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

var _ DeltaPager[models.Messageable] = &mailDeltaPager{}

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
	userID, containerID string,
) *users.ItemMailFoldersItemMessagesDeltaRequestBuilder {
	builder := gs.Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(containerID).
		Messages().
		Delta()

	return builder
}

func (c Mail) NewMailDeltaPager(
	ctx context.Context,
	userID, containerID, prevDeltaLink string,
	immutableIDs bool,
	selectProps ...string,
) DeltaPager[models.Messageable] {
	options := &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		// do NOT set Top.  It limits the total items received.
		QueryParameters: &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(c.options.DeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	var builder *users.ItemMailFoldersItemMessagesDeltaRequestBuilder
	if len(prevDeltaLink) > 0 {
		builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(prevDeltaLink, c.Stable.Adapter())
	} else {
		builder = getMailDeltaBuilder(ctx, c.Stable, userID, containerID)
	}

	return &mailDeltaPager{c.Stable, userID, containerID, builder, options}
}

func (p *mailDeltaPager) GetPage(
	ctx context.Context,
) (DeltaLinkValuer[models.Messageable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *mailDeltaPager) SetNextLink(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *mailDeltaPager) Reset(ctx context.Context) {
	p.builder = getMailDeltaBuilder(ctx, p.gs, p.userID, p.containerID)
}

func (p *mailDeltaPager) ValidModTimes() bool {
	return true
}

func (c Mail) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, containerID, prevDeltaLink string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) (map[string]time.Time, bool, []string, DeltaUpdate, error) {
	ctx = clues.Add(
		ctx,
		"data_category", path.EmailCategory,
		"container_id", containerID)

	deltaPager := c.NewMailDeltaPager(
		ctx,
		userID,
		containerID,
		prevDeltaLink,
		immutableIDs,
		idAnd(lastModifiedDateTime)...)
	pager := c.NewMailPager(
		userID,
		containerID,
		immutableIDs,
		idAnd(lastModifiedDateTime)...)

	return getAddedAndRemovedItemIDs[models.Messageable](
		ctx,
		pager,
		deltaPager,
		prevDeltaLink,
		canMakeDeltaQueries,
		addedAndRemovedByAddtlData[models.Messageable])
}
