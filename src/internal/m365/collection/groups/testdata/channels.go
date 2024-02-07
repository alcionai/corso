package testdata

import (
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/internal/common/ptr"
)

func StubChannels(ids ...string) []models.Channelable {
	sl := make([]models.Channelable, 0, len(ids))

	for _, id := range ids {
		ch := models.NewChannel()
		ch.SetDisplayName(ptr.To(id))
		ch.SetId(ptr.To(id))

		sl = append(sl, ch)
	}

	return sl
}

func StubChatMessages(ids ...string) []models.ChatMessageable {
	sl := make([]models.ChatMessageable, 0, len(ids))

	for _, id := range ids {
		cm := models.NewChatMessage()
		cm.SetId(ptr.To(uuid.NewString()))

		body := models.NewItemBody()
		body.SetContent(ptr.To(id))

		cm.SetBody(body)

		sl = append(sl, cm)
	}

	return sl
}
