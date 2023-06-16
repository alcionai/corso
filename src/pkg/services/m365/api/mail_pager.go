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
	"github.com/alcionai/corso/src/pkg/selectors"
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
// item pager
// ---------------------------------------------------------------------------

var _ itemPager[models.Messageable] = &mailPageCtrl{}

type mailPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersItemMessagesRequestBuilder
	options *users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration
}

func (c Mail) NewMailPager(
	userID, containerID string,
	selectProps ...string,
) itemPager[models.Messageable] {
	options := &users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
		QueryParameters: &users.ItemMailFoldersItemMessagesRequestBuilderGetQueryParameters{
			Top: ptr.To(maxNonDeltaPageSize),
		},
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

	return &mailPageCtrl{c.Stable, builder, options}
}

//lint:ignore U1000 False Positive
func (p *mailPageCtrl) getPage(ctx context.Context) (PageLinkValuer[models.Messageable], error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.Messageable]{PageLinkValuer: page}, nil
}

//lint:ignore U1000 False Positive
func (p *mailPageCtrl) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

// ---------------------------------------------------------------------------
// item ID pager
// ---------------------------------------------------------------------------

var _ itemIDPager = &mailIDPager{}

type mailIDPager struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersItemMessagesRequestBuilder
	options *users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration
}

func (c Mail) NewMailIDsPager(
	ctx context.Context,
	userID, containerID string,
	immutableIDs bool,
) itemIDPager {
	config := &users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersItemMessagesRequestBuilderGetQueryParameters{
			Select: idAnd("isRead"),
			Top:    ptr.To(maxNonDeltaPageSize),
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

	return &mailIDPager{c.Stable, builder, config}
}

func (p *mailIDPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.Messageable]{PageLinkValuer: page}, nil
}

func (p *mailIDPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't have reset
func (p *mailIDPager) reset(context.Context) {}

func (p *mailIDPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Messageable](pl)
}

func (c Mail) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	userID, containerID string,
) (map[string]string, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewMailPager(userID, containerID, mailCollisionKeyProps()...)

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating mail")
	}

	m := map[string]string{}

	for _, item := range items {
		m[MailCollisionKey(item)] = ptr.Val(item.GetId())
	}

	return m, nil
}

// ---------------------------------------------------------------------------
// delta item ID pager
// ---------------------------------------------------------------------------

var _ itemIDPager = &mailDeltaIDPager{}

type mailDeltaIDPager struct {
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

func (c Mail) NewMailDeltaIDsPager(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
) itemIDPager {
	config := &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{
			Select: idAnd("isRead"),
			Top:    ptr.To(maxDeltaPageSize),
		},
		Headers: newPreferHeaders(preferPageSize(maxDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	var builder *users.ItemMailFoldersItemMessagesDeltaRequestBuilder

	if len(oldDelta) > 0 {
		builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(oldDelta, c.Stable.Adapter())
	} else {
		builder = getMailDeltaBuilder(ctx, c.Stable, userID, containerID, config)
	}

	return &mailDeltaIDPager{c.Stable, userID, containerID, builder, config}
}

func (p *mailDeltaIDPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return page, nil
}

func (p *mailDeltaIDPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *mailDeltaIDPager) reset(ctx context.Context) {
	p.builder = p.gs.
		Client().
		Users().
		ByUserId(p.userID).
		MailFolders().
		ByMailFolderId(p.containerID).
		Messages().
		Delta()
}

func (p *mailDeltaIDPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
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

	pager := c.NewMailIDsPager(ctx, userID, containerID, immutableIDs)
	deltaPager := c.NewMailDeltaIDsPager(ctx, userID, containerID, oldDelta, immutableIDs)

	return getAddedAndRemovedItemIDs(ctx, c.Stable, pager, deltaPager, oldDelta, canMakeDeltaQueries)
}
