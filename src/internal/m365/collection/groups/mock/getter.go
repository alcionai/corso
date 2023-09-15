package mock

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type GetChannelMessage struct {
	Err error
}

func (m GetChannelMessage) GetChannelMessage(
	ctx context.Context,
	teamID, channelID, itemID string,
) (models.ChatMessageable, *details.GroupsInfo, error) {
	msg := models.NewChatMessage()
	msg.SetId(ptr.To(itemID))

	return msg, &details.GroupsInfo{}, m.Err
}
