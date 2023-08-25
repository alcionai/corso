package groups

import (
	"context"

	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

var _ BackupHandler = &groupBackupHandler{}

type groupBackupHandler struct {
	ac      api.Channels
	groupID string
	scope   selectors.GroupsScope
}

func NewGroupBackupHandler(groupID string, ac api.Channels, scope selectors.GroupsScope) groupBackupHandler {
	return groupBackupHandler{
		ac:      ac,
		groupID: groupID,
		scope:   scope,
	}
}

func (gHandler groupBackupHandler) GetChannelByID(ctx context.Context, teamID, channelID string) (models.Channelable, error) {
	return gHandler.ac.Client.Channels().GetChannel(ctx, teamID, channelID)
}

func (gHandler groupBackupHandler) NewChannelsPager(
	teamID string,
	fields []string,
) api.ChannelDeltaEnumerator {
	return gHandler.ac.NewChannelPager(teamID, fields)
}

func (gHandler groupBackupHandler) GetMessageByID(
	ctx context.Context,
	teamID, channelID, itemID string,
) (models.ChatMessageable, error) {
	return gHandler.ac.GetMessageByID(ctx, teamID, channelID, itemID)
}

func (gHandler groupBackupHandler) NewMessagePager(
	teamID, channelID string,
	fields []string,
) api.ChannelMessageDeltaEnumerator {
	return gHandler.ac.NewMessagePager(teamID, channelID, fields)
}

func (gHandler groupBackupHandler) GetMessageReplies(
	ctx context.Context,
	teamID, channelID, messageID string,
) (serialization.Parsable, error) {
	return gHandler.ac.GetReplies(ctx, teamID, channelID, messageID)
}
