package groups

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ BackupHandler = &groupBackupHandler{}

type groupBackupHandler struct {
	ac      api.Channels
	groupID string
}

func NewGroupBackupHandler(groupID string, ac api.Channels, scope selectors.GroupsScope) groupBackupHandler {
	return groupBackupHandler{
		ac:      ac,
		groupID: groupID,
	}
}

func (gHandler groupBackupHandler) GetChannelByID(
	ctx context.Context,
	teamID, channelID string,
) (models.Channelable, error) {
	return gHandler.ac.Client.Channels().GetChannel(ctx, teamID, channelID)
}

func (gHandler groupBackupHandler) NewChannelsPager(
	teamID string,
) api.ChannelDeltaEnumerator {
	return gHandler.ac.NewChannelPager(teamID)
}

func (gHandler groupBackupHandler) GetMessageByID(
	ctx context.Context,
	teamID, channelID, itemID string,
) (models.ChatMessageable, error) {
	chatMessage, _, err := gHandler.ac.GetMessage(ctx, teamID, channelID, itemID)
	return chatMessage, err
}

func (gHandler groupBackupHandler) NewMessagePager(
	teamID, channelID string,
) api.ChannelMessageDeltaEnumerator {
	return gHandler.ac.NewMessagePager(teamID, channelID)
}

func (gHandler groupBackupHandler) GetMessageReplies(
	ctx context.Context,
	teamID, channelID, messageID string,
) (serialization.Parsable, error) {
	return gHandler.ac.GetReplies(ctx, teamID, channelID, messageID)
}
