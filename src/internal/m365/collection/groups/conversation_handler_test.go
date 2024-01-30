package groups

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/collection/groups/metadata"
	"github.com/alcionai/corso/src/internal/tester"
)

const (
	resourceEmail = "test@example.com"
)

type ConversationHandlerUnitSuite struct {
	tester.Suite
}

func TestConversationHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &ConversationHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// Basic test to ensure metadata is serialized and deserialized correctly.
func (suite *ConversationHandlerUnitSuite) TestGetItemMetadata() {
	var (
		t  = suite.T()
		bh = conversationsBackupHandler{
			resourceEmail: resourceEmail,
		}

		topic = "test topic"
		conv  = models.NewConversation()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	conv.SetTopic(&topic)

	rc, size, err := bh.getItemMetadata(ctx, conv)
	assert.NoError(t, err, clues.ToCore(err))

	require.NotNil(t, rc, "nil read closer")
	assert.Greater(t, size, 0, "incorrect size")

	defer rc.Close()

	m, err := io.ReadAll(rc)
	assert.NoError(t, err, "reading metadata")

	var meta metadata.ConversationPostMetadata

	err = json.Unmarshal(m, &meta)
	assert.NoError(t, err, "deserializing metadata")

	assert.Equal(t, []string{resourceEmail}, meta.Recipients, "incorrect recipients")
	assert.Equal(t, ptr.Val(conv.GetTopic()), meta.Topic, "incorrect topic")
}
