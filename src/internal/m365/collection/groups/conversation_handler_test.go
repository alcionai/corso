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

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/internal/m365/collection/groups/metadata"
	"github.com/alcionai/canario/src/internal/tester"
	deltaPath "github.com/alcionai/canario/src/pkg/backup/metadata"
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

func (suite *ConversationHandlerUnitSuite) TestMakeTombstones() {
	table := []struct {
		name        string
		dps         deltaPath.DeltaPaths
		expected    map[string]string
		expectedErr require.ErrorAssertionFunc
	}{
		{
			name: "valid",
			dps: deltaPath.DeltaPaths{
				"c1/t1": deltaPath.DeltaPath{
					Path: "p1",
				},
				"c2/t2": deltaPath.DeltaPath{
					Path: "p2",
				},
			},
			expected: map[string]string{
				"c1": "p1",
				"c2": "p2",
			},
			expectedErr: require.NoError,
		},
		{
			name: "invalid prev path",
			dps: deltaPath.DeltaPaths{
				"c1": deltaPath.DeltaPath{
					Path: "p1",
				},
			},
			expected:    nil,
			expectedErr: require.Error,
		},
		{
			name: "invalid prev path 2",
			dps: deltaPath.DeltaPaths{
				"c1/t1/a1": deltaPath.DeltaPath{
					Path: "p1",
				},
			},
			expected:    nil,
			expectedErr: require.Error,
		},
		{
			name: "multiple threads exist for a conversation",
			dps: deltaPath.DeltaPaths{
				"c1/t1": deltaPath.DeltaPath{
					Path: "p1",
				},
				"c1/t2": deltaPath.DeltaPath{
					Path: "p2",
				},
			},
			expected:    nil,
			expectedErr: require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			bh := conversationsBackupHandler{}

			result, err := bh.makeTombstones(test.dps)
			test.expectedErr(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}
