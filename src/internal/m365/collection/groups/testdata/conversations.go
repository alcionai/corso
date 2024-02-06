package testdata

import (
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

func StubConversations(ids ...string) []models.Conversationable {
	sl := make([]models.Conversationable, 0, len(ids))

	for _, id := range ids {
		c := models.NewConversation()
		c.SetId(ptr.To(id))

		sl = append(sl, c)
	}

	return sl
}

func StubConversationThreads(ids ...string) []models.ConversationThreadable {
	sl := make([]models.ConversationThreadable, 0, len(ids))

	for _, id := range ids {
		ct := models.NewConversationThread()
		ct.SetId(ptr.To(id))

		sl = append(sl, ct)
	}

	return sl
}

func StubPosts(ids ...string) []models.Postable {
	sl := make([]models.Postable, 0, len(ids))

	for _, id := range ids {
		p := models.NewPost()
		p.SetId(ptr.To(uuid.NewString()))

		body := models.NewItemBody()
		body.SetContent(ptr.To(id))

		p.SetBody(body)

		sl = append(sl, p)
	}

	return sl
}
