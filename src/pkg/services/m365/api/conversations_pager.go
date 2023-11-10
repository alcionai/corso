package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// conversation pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Conversationable] = &conversationsPageCtrl{}

type conversationsPageCtrl struct {
	resourceID string
	gs         graph.Servicer
	builder    *groups.ItemConversationsRequestBuilder
	options    *groups.ItemConversationsRequestBuilderGetRequestConfiguration
}

func (p *conversationsPageCtrl) SetNextLink(nextLink string) {
	p.builder = groups.NewItemConversationsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *conversationsPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.Conversationable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *conversationsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Conversations) NewConversationsPager(
	groupID string,
	cc CallConfig,
) *conversationsPageCtrl {
	builder := c.Stable.
		Client().
		Groups().
		ByGroupId(groupID).
		Conversations()

	options := &groups.ItemConversationsRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.ItemConversationsRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	return &conversationsPageCtrl{
		resourceID: groupID,
		builder:    builder,
		gs:         c.Stable,
		options:    options,
	}
}

// GetConversations fetches all conversations in the group.
func (c Conversations) GetConversations(
	ctx context.Context,
	groupID string,
	cc CallConfig,
) ([]models.Conversationable, error) {
	pager := c.NewConversationsPager(groupID, cc)
	items, err := pagers.BatchEnumerateItems[models.Conversationable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// conversation thread pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.ConversationThreadable] = &conversationThreadsPageCtrl{}

type conversationThreadsPageCtrl struct {
	resourceID, conversationID string
	gs                         graph.Servicer
	builder                    *groups.ItemConversationsItemThreadsRequestBuilder
	options                    *groups.ItemConversationsItemThreadsRequestBuilderGetRequestConfiguration
}

func (p *conversationThreadsPageCtrl) SetNextLink(nextLink string) {
	p.builder = groups.NewItemConversationsItemThreadsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *conversationThreadsPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.ConversationThreadable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *conversationThreadsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Conversations) NewConversationThreadsPager(
	groupID, conversationID string,
	cc CallConfig,
) *conversationThreadsPageCtrl {
	builder := c.Stable.
		Client().
		Groups().
		ByGroupId(groupID).
		Conversations().
		ByConversationId(conversationID).
		Threads()

	options := &groups.ItemConversationsItemThreadsRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.ItemConversationsItemThreadsRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	if len(cc.Expand) > 0 {
		options.QueryParameters.Expand = cc.Expand
	}

	return &conversationThreadsPageCtrl{
		resourceID:     groupID,
		conversationID: conversationID,
		builder:        builder,
		gs:             c.Stable,
		options:        options,
	}
}

// GetConversations fetches all conversation threads in the group.
func (c Conversations) GetConversationThreads(
	ctx context.Context,
	groupID, conversationID string,
	cc CallConfig,
) ([]models.ConversationThreadable, error) {
	ctx = clues.Add(ctx, "conversation_id", conversationID)
	pager := c.NewConversationThreadsPager(groupID, conversationID, cc)
	items, err := pagers.BatchEnumerateItems[models.ConversationThreadable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// conversation thread post pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Postable] = &conversationThreadPostsPageCtrl{}

type conversationThreadPostsPageCtrl struct {
	resourceID, conversationID, threadID string
	gs                                   graph.Servicer
	builder                              *groups.ItemConversationsItemThreadsItemPostsRequestBuilder
	options                              *groups.ItemConversationsItemThreadsItemPostsRequestBuilderGetRequestConfiguration
}

func (p *conversationThreadPostsPageCtrl) SetNextLink(nextLink string) {
	p.builder = groups.NewItemConversationsItemThreadsItemPostsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *conversationThreadPostsPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.Postable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *conversationThreadPostsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Conversations) NewConversationThreadPostsPager(
	groupID, conversationID, threadID string,
	cc CallConfig,
) *conversationThreadPostsPageCtrl {
	builder := c.Stable.
		Client().
		Groups().
		ByGroupId(groupID).
		Conversations().
		ByConversationId(conversationID).
		Threads().
		ByConversationThreadId(threadID).
		Posts()

	options := &groups.ItemConversationsItemThreadsItemPostsRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.ItemConversationsItemThreadsItemPostsRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	if len(cc.Expand) > 0 {
		options.QueryParameters.Expand = cc.Expand
	}

	return &conversationThreadPostsPageCtrl{
		resourceID:     groupID,
		conversationID: conversationID,
		threadID:       threadID,
		builder:        builder,
		gs:             c.Stable,
		options:        options,
	}
}

// GetConversations fetches all conversation posts in the group.
func (c Conversations) GetConversationThreadPosts(
	ctx context.Context,
	groupID, conversationID, threadID string,
	cc CallConfig,
) ([]models.Postable, error) {
	ctx = clues.Add(ctx, "conversation_id", conversationID)
	pager := c.NewConversationThreadPostsPager(groupID, conversationID, threadID, cc)
	items, err := pagers.BatchEnumerateItems[models.Postable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// GetConversations fetches all added and deleted conversation posts in the group.
func (c Conversations) GetConversationThreadPostIDs(
	ctx context.Context,
	groupID, conversationID, threadID string,
	cc CallConfig,
) (pagers.AddedAndRemoved, error) {
	canMakeDeltaQueries := false

	aarh, err := pagers.GetAddedAndRemovedItemIDs[models.Postable](
		ctx,
		c.NewConversationThreadPostsPager(groupID, conversationID, threadID, CallConfig{}),
		nil,
		"",
		canMakeDeltaQueries,
		pagers.AddedAndRemovedAddAll[models.Postable],
		pagers.FilterIncludeAll[models.Postable])

	return aarh, clues.Stack(err).OrNil()
}
