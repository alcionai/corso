package mock

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

type GetChannelMessage struct {
	Err error
}

func (m GetChannelMessage) GetItem(
	_ context.Context,
	_ string,
	_ path.Elements,
	itemID string,
) (models.ChatMessageable, *details.GroupsInfo, error) {
	msg := models.NewChatMessage()
	msg.SetId(ptr.To(itemID))

	return msg, &details.GroupsInfo{}, m.Err
}
