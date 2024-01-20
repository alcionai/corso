package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/chats"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// chat message pager
// ---------------------------------------------------------------------------

// delta queries are not supported
var _ pagers.NonDeltaHandler[models.ChatMessageable] = &chatMessagePageCtrl{}

type chatMessagePageCtrl struct {
	chatID  string
	gs      graph.Servicer
	builder *chats.ItemMessagesRequestBuilder
	options *chats.ItemMessagesRequestBuilderGetRequestConfiguration
}

func (p *chatMessagePageCtrl) SetNextLink(nextLink string) {
	p.builder = chats.NewItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *chatMessagePageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.ChatMessageable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *chatMessagePageCtrl) ValidModTimes() bool {
	return true
}

func (c Chats) NewChatMessagePager(
	chatID string,
	cc CallConfig,
) *chatMessagePageCtrl {
	builder := c.Stable.
		Client().
		Chats().
		ByChatId(chatID).
		Messages()

	options := &chats.ItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &chats.ItemMessagesRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	if len(cc.Expand) > 0 {
		options.QueryParameters.Expand = cc.Expand
	}

	return &chatMessagePageCtrl{
		chatID:  chatID,
		builder: builder,
		gs:      c.Stable,
		options: options,
	}
}

// GetChatMessages fetches a delta of all messages in the chat.
func (c Chats) GetChatMessages(
	ctx context.Context,
	chatID string,
	cc CallConfig,
) ([]models.ChatMessageable, error) {
	ctx = clues.Add(ctx, "chat_id", chatID)
	pager := c.NewChatMessagePager(chatID, cc)
	items, err := pagers.BatchEnumerateItems[models.ChatMessageable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// GetChatMessageIDs fetches a delta of all messages in the chat.
// returns two maps: addedItems, deletedItems
func (c Chats) GetChatMessageIDs(
	ctx context.Context,
	chatID string,
	cc CallConfig,
) (pagers.AddedAndRemoved, error) {
	aar, err := pagers.GetAddedAndRemovedItemIDs[models.ChatMessageable](
		ctx,
		c.NewChatMessagePager(chatID, CallConfig{}),
		nil,
		"",
		false, // delta queries are not supported
		0,
		pagers.AddedAndRemovedByDeletedDateTime[models.ChatMessageable],
		IsNotSystemMessage)

	return aar, clues.Stack(err).OrNil()
}

// ---------------------------------------------------------------------------
// chat pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Chatable] = &chatPageCtrl{}

type chatPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemChatsRequestBuilder
	options *users.ItemChatsRequestBuilderGetRequestConfiguration
}

func (p *chatPageCtrl) SetNextLink(nextLink string) {
	p.builder = users.NewItemChatsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *chatPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.Chatable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *chatPageCtrl) ValidModTimes() bool {
	return false
}

func (c Chats) NewChatPager(
	userID string,
	cc CallConfig,
) *chatPageCtrl {
	options := &users.ItemChatsRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemChatsRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	if len(cc.Expand) > 0 {
		options.QueryParameters.Expand = cc.Expand
	}

	res := &chatPageCtrl{
		gs:      c.Stable,
		options: options,
		builder: c.Stable.
			Client().
			Users().
			ByUserId(userID).
			Chats(),
	}

	return res
}

// GetChats fetches all chats in the team.
func (c Chats) GetChats(
	ctx context.Context,
	userID string,
	cc CallConfig,
) ([]models.Chatable, error) {
	return pagers.BatchEnumerateItems[models.Chatable](ctx, c.NewChatPager(userID, cc))
}
