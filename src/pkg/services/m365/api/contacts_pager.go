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

var _ itemPager = &contactPager{}

type contactPager struct {
	gs      graph.Servicer
	builder *users.ItemContactFoldersItemContactsRequestBuilder
	options *users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration
}

func (c Contacts) NewContactPager(
	ctx context.Context,
	userID, containerID string,
	immutableIDs bool,
) itemPager {
	config := &users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersItemContactsRequestBuilderGetQueryParameters{
			Select: idAnd(parentFolderID),
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

	return &contactPager{c.Stable, builder, config}
}

func (p *contactPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.Contactable]{PageLinkValuer: resp}, nil
}

func (p *contactPager) setNext(nextLink string) {
	p.builder = users.NewItemContactFoldersItemContactsRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't need reset
func (p *contactPager) reset(context.Context) {}

func (p *contactPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Contactable](pl)
}

// ---------------------------------------------------------------------------
// delta item pager
// ---------------------------------------------------------------------------

var _ itemPager = &contactDeltaPager{}

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
	options *users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration,
) *users.ItemContactFoldersItemContactsDeltaRequestBuilder {
	builder := gs.Client().Users().ByUserId(userID).ContactFolders().ByContactFolderId(containerID).Contacts().Delta()
	return builder
}

func (c Contacts) NewContactDeltaPager(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
) itemPager {
	options := &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetQueryParameters{
			Select: idAnd(parentFolderID),
		},
		Headers: newPreferHeaders(preferPageSize(maxDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	var builder *users.ItemContactFoldersItemContactsDeltaRequestBuilder
	if oldDelta != "" {
		builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(oldDelta, c.Stable.Adapter())
	} else {
		builder = getContactDeltaBuilder(ctx, c.Stable, userID, containerID, options)
	}

	return &contactDeltaPager{c.Stable, userID, containerID, builder, options}
}

func (p *contactDeltaPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
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
	p.builder = getContactDeltaBuilder(ctx, p.gs, p.userID, p.containerID, p.options)
}

func (p *contactDeltaPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
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

	pager := c.NewContactPager(ctx, userID, containerID, immutableIDs)
	deltaPager := c.NewContactDeltaPager(ctx, userID, containerID, oldDelta, immutableIDs)

	return getAddedAndRemovedItemIDs(ctx, c.Stable, pager, deltaPager, oldDelta, canMakeDeltaQueries)
}
