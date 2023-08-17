package groups

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type BackupChannelHandler interface {
	GetChannel(ctx context.Context, teamID, channelID string) (models.Channelable, error)
	NewChannelPager(teamID, channelID string) api.ChannelItemDeltaEnumerator
}

type BackupMessagesHandler interface {
	GetItem(ctx context.Context, teamID, channelID, itemID string) (models.ChatMessageable, error)
	NewItemPager(teamID, channelID string) api.MessageItemDeltaEnumerator
}

type BackupReplyHandler interface {
	GetItem(ctx context.Context, teamID, channelID, messageID string) (serialization.Parsable, error)
}
