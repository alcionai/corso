package testdata

import (
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

func StubChannels(names ...string) []models.Channelable {
	sl := make([]models.Channelable, 0, len(names))

	for _, name := range names {
		ch := models.NewChannel()
		ch.SetDisplayName(ptr.To(name))
		ch.SetId(ptr.To(uuid.NewString()))

		sl = append(sl, ch)
	}

	return sl
}

func StubChatMessages(names ...string) []models.ChatMessageable {
	sl := make([]models.ChatMessageable, 0, len(names))

	for _, name := range names {
		cm := models.NewChatMessage()
		cm.SetId(ptr.To(uuid.NewString()))

		body := models.NewItemBody()
		body.SetContent(ptr.To(name))

		cm.SetBody(body)

		sl = append(sl, cm)
	}

	return sl
}
