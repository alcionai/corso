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
		fields []string,
	) api.ChannelDeltaEnumerator

	GetMessageByID(
		ctx context.Context,
		teamID, channelID, itemID string,
	) (models.ChatMessageable, error)
	NewMessagePager(
		teamID, channelID string,
		fields []string,
	) api.ChannelMessageDeltaEnumerator

	GetMessageReplies(
		ctx context.Context,
		teamID, channelID, messageID string,
	) (serialization.Parsable, error)
}
