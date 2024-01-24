package api

import (
	"context"
	"time"

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
	var (
		// in case of an empty chat, we want to use Val instead of OrNow
		lastModTime      = ptr.Val(chat.GetLastUpdatedDateTime())
		lastMsgPreview   = chat.GetLastMessagePreview()
		lastMsgCreatedAt time.Time
		members          = chat.GetMembers()
		memberNames      = []string{}
		msgs             = chat.GetMessages()
		preview          string
		err              error
	)

	if lastMsgPreview != nil {
		preview, _, err = getChatMessageContentPreview(lastMsgPreview, noAttachments{})
		if err != nil {
			preview = "malformed or unparseable html" + preview
		}

		// in case of an empty mod time, we want to use the chat's mod time
		// therefore Val instaed of OrNow
		lastMsgCreatedAt = ptr.Val(lastMsgPreview.GetCreatedDateTime())
		if lastModTime.Before(lastMsgCreatedAt) {
			lastModTime = lastMsgCreatedAt
		}
	}

	for _, m := range members {
		memberNames = append(memberNames, ptr.Val(m.GetDisplayName()))
	}

	return &details.TeamsChatsInfo{
		ItemType: details.TeamsChat,
		Modified: lastModTime,

		Chat: details.ChatInfo{
			CreatedAt:          ptr.OrNow(chat.GetCreatedDateTime()),
			LastMessageAt:      lastMsgCreatedAt,
			LastMessagePreview: preview,
			Members:            memberNames,
			MessageCount:       len(msgs),
			Name:               ptr.Val(chat.GetTopic()),
		},
	}
}
