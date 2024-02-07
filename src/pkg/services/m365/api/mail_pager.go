package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
	"github.com/alcionai/canario/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// container pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.MailFolderable] = &mailFoldersPageCtrl{}

type mailFoldersPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersRequestBuilder
	options *users.ItemMailFoldersRequestBuilderGetRequestConfiguration
}

func (c Mail) NewMailFoldersPager(
	userID string,
	selectProps ...string,
) pagers.NonDeltaHandler[models.MailFolderable] {
	options := &users.ItemMailFoldersRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(
			preferPageSize(maxNonDeltaPageSize),
			preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
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
) (pagers.NextLinkValuer[models.MailFolderable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, clues.Stack(err).OrNil()
}

func (p *mailFoldersPageCtrl) SetNextLink(nextLink string) {
	p.builder = users.NewItemMailFoldersRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *mailFoldersPageCtrl) ValidModTimes() bool {
	return true
}

// EnumerateContainers retrieves all of the user's current mail folders.
func (c Mail) EnumerateContainers(
	ctx context.Context,
	userID, _ string, // baseContainerID not needed here
) ([]models.MailFolderable, error) {
	containers, err := pagers.BatchEnumerateItems(ctx, c.NewMailFoldersPager(userID))
	return containers, clues.Stack(err).OrNil()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Messageable] = &mailsPageCtrl{}

type mailsPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersItemMessagesRequestBuilder
	options *users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration
}

func (c Mail) NewMailPager(
	userID, containerID string,
	selectProps ...string,
) pagers.NonDeltaHandler[models.Messageable] {
	options := &users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(
			preferPageSize(maxNonDeltaPageSize),
			preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
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
) (pagers.NextLinkValuer[models.Messageable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, clues.Stack(err).OrNil()
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
	pager := c.NewMailPager(userID, containerID, mailCollisionKeyProps()...)

	items, err := pagers.BatchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, clues.Wrap(err, "enumerating mails")
	}

	m := map[string]string{}

	for _, item := range items {
		m[MailCollisionKey(item)] = ptr.Val(item.GetId())
	}

	return m, nil
}

func (c Mail) GetItemsInContainer(
	ctx context.Context,
	userID, containerID string,
) ([]models.Messageable, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewMailPager(userID, containerID)

	items, err := pagers.BatchEnumerateItems(ctx, pager)

	return items, clues.Wrap(err, "enumerating mails").OrNil()
}

func (c Mail) GetItemIDsInContainer(
	ctx context.Context,
	userID, containerID string,
) (map[string]struct{}, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewMailPager(userID, containerID, idAnd()...)

	items, err := pagers.BatchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, clues.Wrap(err, "enumerating mails")
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

var _ pagers.DeltaHandler[models.Messageable] = &mailDeltaPager{}

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
	selectProps ...string,
) pagers.DeltaHandler[models.Messageable] {
	options := &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		// do NOT set Top.  It limits the total items received.
		QueryParameters: &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{},
		Headers: newPreferHeaders(
			preferPageSize(c.options.DeltaPageSize),
			preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
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
) (pagers.DeltaLinkValuer[models.Messageable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, clues.Stack(err).OrNil()
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
	config CallConfig,
) (pagers.AddedAndRemoved, error) {
	ctx = clues.Add(
		ctx,
		"data_category", path.EmailCategory,
		"container_id", containerID)

	deltaPager := c.NewMailDeltaPager(
		ctx,
		userID,
		containerID,
		prevDeltaLink,
		idAnd(lastModifiedDateTime)...)
	pager := c.NewMailPager(
		userID,
		containerID,
		idAnd(lastModifiedDateTime)...)

	return pagers.GetAddedAndRemovedItemIDs[models.Messageable](
		ctx,
		pager,
		deltaPager,
		prevDeltaLink,
		config.CanMakeDeltaQueries,
		config.LimitResults,
		pagers.AddedAndRemovedByAddtlData[models.Messageable])
}
