package api

import (
	"regexp"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
)

type ChatsPagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestChatsPagerIntgSuite(t *testing.T) {
	suite.Run(t, &ChatsPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ChatsPagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *ChatsPagerIntgSuite) TestEnumerateChats() {
	var (
		t  = suite.T()
		ac = suite.its.ac.Chats()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	cc := CallConfig{
		Expand: []string{"lastMessagePreview"},
	}

	chats, err := ac.GetChats(ctx, suite.its.user.id, cc)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, chats)

	for _, chat := range chats {
		chatID := ptr.Val(chat.GetId())

		suite.Run("chat_"+chatID, func() {
			testGetChatByID(suite.T(), ac, chatID)
		})

		suite.Run("chat_messages_"+chatID, func() {
			testEnumerateChatMessages(
				suite.T(),
				ac,
				chatID,
				chat.GetLastMessagePreview())
		})
	}
}

func testGetChatByID(
	t *testing.T,
	ac Chats,
	chatID string,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	cc := CallConfig{}

	chat, _, err := ac.GetChatByID(ctx, chatID, cc)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, chat)
}

var attachmentHtmlRegexp = regexp.MustCompile("<attachment id=\"[a-zA-Z0-9].*\"></attachment>")

func testEnumerateChatMessages(
	t *testing.T,
	ac Chats,
	chatID string,
	lastMessagePreview models.ChatMessageInfoable,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	cc := CallConfig{}

	messages, err := ac.GetChatMessages(ctx, chatID, cc)
	require.NoError(t, err, clues.ToCore(err))

	var lastID string
	if lastMessagePreview != nil {
		lastID = ptr.Val(lastMessagePreview.GetId())
	}

	for _, msg := range messages {
		msgID := ptr.Val(msg.GetId())

		assert.Equal(
			t,
			chatID,
			ptr.Val(msg.GetChatId()),
			"message:",
			msgID)

		if msgID == lastID {
			previewContent := ptr.Val(lastMessagePreview.GetBody().GetContent())
			msgContent := ptr.Val(msg.GetBody().GetContent())

			previewContent = replaceAttachmentMarkup(previewContent, nil)
			msgContent = replaceAttachmentMarkup(msgContent, nil)

			assert.Equal(
				t,
				previewContent,
				msgContent)
		}
	}
}
