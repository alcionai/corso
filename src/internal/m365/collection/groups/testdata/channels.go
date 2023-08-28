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
		ch.SetDisplayName(&name)
		ch.SetId(ptr.To(uuid.NewString()))
	}

	return sl
}
