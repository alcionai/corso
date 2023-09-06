package api

import (
	"context"
	"fmt"

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

func (p *mailFolderPager) getPage(ctx context.Context) (NextLinker, error) {
	page, err := p.builder.Get(ctx, nil)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return page, nil
}

func (p *mailFolderPager) SetNextLink(nextLink string) {
	p.builder = users.NewItemMailFoldersRequestBuilder(nextLink, p.service.Adapter())
}

func (p *mailFolderPager) valuesIn(pl NextLinker) ([]models.MailFolderable, error) {
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

		pgr.SetNextLink(link)
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
		ByUserIdString(userID).
		MailFolders().
		ByMailFolderIdString(containerID).
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
	options *users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration,
) *users.ItemMailFoldersItemMessagesDeltaRequestBuilder {
	builder := gs.Client().
		Users().
		ByUserIdString(userID).
		MailFolders().
		ByMailFolderIdString(containerID).
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
	if len(prevDeltaLink) == 0 {
		builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(prevDeltaLink, c.Stable.Adapter())
	} else {
		builder = getMailDeltaBuilder(ctx, c.Stable, userID, containerID, options)
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
	p.builder = getMailDeltaBuilder(ctx, p.gs, p.userID, p.containerID, p.options)
}

func (c Mail) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, containerID, prevDeltaLink string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
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
		idAnd(additionalData)...)
	pager := c.NewMailPager(
		userID,
		containerID,
		immutableIDs,
		idAnd(additionalData)...)

	return getAddedAndRemovedItemIDs[models.Messageable](
		ctx,
		pager,
		deltaPager,
		prevDeltaLink,
		canMakeDeltaQueries)
}
