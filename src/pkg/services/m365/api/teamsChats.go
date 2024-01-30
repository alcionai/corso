package api

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/chats"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Chats() Chats {
	return Chats{c}
}

// Chats is an interface-compliant provider of the client.
type Chats struct {
	Client
}

// ---------------------------------------------------------------------------
// Chats
// ---------------------------------------------------------------------------

func (c Chats) GetChatByID(
	ctx context.Context,
	chatID string,
	cc CallConfig,
) (models.Chatable, *details.TeamsChatsInfo, error) {
	config := &chats.ChatItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &chats.ChatItemRequestBuilderGetQueryParameters{},
	}

	if len(cc.Select) > 0 {
		config.QueryParameters.Select = cc.Select
	}

	if len(cc.Expand) > 0 {
		config.QueryParameters.Expand = cc.Expand
	}

	resp, err := c.Stable.
		Client().
		Chats().
		ByChatId(chatID).
		Get(ctx, config)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	return resp, TeamsChatInfo(resp), nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func TeamsChatInfo(chat models.Chatable) *details.TeamsChatsInfo {
	return &details.TeamsChatsInfo{
		ItemType: details.TeamsChat,
		Modified: ptr.OrNow(chat.GetLastUpdatedDateTime()),

		Chat: details.ChatInfo{
			CreatedAt: ptr.OrNow(chat.GetCreatedDateTime()),
			Name:      ptr.Val(chat.GetTopic()),
		},
	}
}
