package testdata

import (
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

func StubChats(ids ...string) []models.Chatable {
	sl := make([]models.Chatable, 0, len(ids))

	for _, id := range ids {
		chat := models.NewChat()
		chat.SetTopic(ptr.To(id))
		chat.SetId(ptr.To(id))

		// we should expect to get the latest message preview by default
		lastMsgPrv := models.NewChatMessageInfo()
		lastMsgPrv.SetId(ptr.To(uuid.NewString()))

		body := models.NewItemBody()
		body.SetContent(ptr.To(id))
		lastMsgPrv.SetBody(body)

		chat.SetLastMessagePreview(lastMsgPrv)

		sl = append(sl, chat)
	}

	return sl
}

func StubChatMessages(ids ...string) []models.ChatMessageable {
	sl := make([]models.ChatMessageable, 0, len(ids))

	var lastMsg models.ChatMessageable

	for _, id := range ids {
		msg := models.NewChatMessage()
		msg.SetId(ptr.To(uuid.NewString()))

		body := models.NewItemBody()
		body.SetContent(ptr.To(id))

		msg.SetBody(body)

		sl = append(sl, msg)
		lastMsg = msg
	}

	lastMsgPrv := models.NewChatMessageInfo()
	lastMsgPrv.SetId(lastMsg.GetId())
	lastMsgPrv.SetBody(lastMsg.GetBody())

	return sl
}
