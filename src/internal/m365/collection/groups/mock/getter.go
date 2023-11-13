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

func (m GetChannelMessage) GetItemByID(
	ctx context.Context,
	groupID, channelID, messageID string,
) (models.ChatMessageable, *details.GroupsInfo, error) {
	msg := models.NewChatMessage()
	msg.SetId(ptr.To(messageID))

	return msg, &details.GroupsInfo{}, m.Err
}
