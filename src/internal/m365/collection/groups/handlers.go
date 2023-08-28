package groups

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type BackupHandler interface {
	GetChannelByID(
		ctx context.Context,
		teamID, channelID string,
	) (models.Channelable, error)
	NewChannelsPager(
		teamID string,
	) api.DeltaPager[models.Channelable]

	GetMessageByID(
		ctx context.Context,
		teamID, channelID, itemID string,
	) (models.ChatMessageable, error)
	NewMessagePager(
		teamID, channelID string,
	) api.DeltaPager[models.ChatMessageable]

	GetMessageReplies(
		ctx context.Context,
		teamID, channelID, messageID string,
	) (serialization.Parsable, error)
}

type BackupMessagesHandler interface {
	GetMessage(ctx context.Context, teamID, channelID, itemID string) (models.ChatMessageable, error)
	NewMessagePager(teamID, channelID string) api.DeltaPager[models.ChatMessageable]
	GetChannel(ctx context.Context, teamID, channelID string) (models.Channelable, error)
	GetReply(ctx context.Context, teamID, channelID, messageID string) (serialization.Parsable, error)
}
