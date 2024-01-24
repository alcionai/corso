package api

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/collection/teamschats/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type ChatsAPIUnitSuite struct {
	tester.Suite
}

func TestChatsAPIUnitSuite(t *testing.T) {
	suite.Run(t, &ChatsAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ChatsAPIUnitSuite) TestChatsInfo() {
	start := time.Now()

	tests := []struct {
		name     string
		expected func() (models.Chatable, *details.TeamsChatsInfo)
	}{
		{
			name: "Empty chat",
			expected: func() (models.Chatable, *details.TeamsChatsInfo) {
				chat := models.NewChat()

				i := &details.TeamsChatsInfo{
					ItemType: details.TeamsChat,
					Modified: ptr.Val(chat.GetLastUpdatedDateTime()),
					Chat:     details.ChatInfo{},
				}

				return chat, i
			},
		},
		{
			name: "All fields",
			expected: func() (models.Chatable, *details.TeamsChatsInfo) {
				now := time.Now()
				then := now.Add(1 * time.Hour)
				id := uuid.NewString()

				chat := testdata.StubChats(id)[0]
				chat.SetTopic(ptr.To("Hello world"))
				chat.SetCreatedDateTime(&now)
				chat.SetLastUpdatedDateTime(&now)
				chat.GetLastMessagePreview().SetCreatedDateTime(&then)

				msgs := testdata.StubChatMessages(ptr.Val(chat.GetLastMessagePreview().GetId()))
				chat.SetMessages(msgs)

				i := &details.TeamsChatsInfo{
					ItemType: details.TeamsChat,
					Modified: then,
					Chat: details.ChatInfo{
						Name:               "Hello world",
						LastMessageAt:      then,
						LastMessagePreview: id,
						Members:            []string{},
						MessageCount:       1,
					},
				}

				return chat, i
			},
		},
		{
			name: "last message preview, but no messages",
			expected: func() (models.Chatable, *details.TeamsChatsInfo) {
				now := time.Now()
				then := now.Add(1 * time.Hour)
				id := uuid.NewString()

				chat := testdata.StubChats(id)[0]
				chat.SetTopic(ptr.To("Hello world"))
				chat.SetCreatedDateTime(&now)
				chat.SetLastUpdatedDateTime(&now)
				chat.GetLastMessagePreview().SetCreatedDateTime(&then)

				i := &details.TeamsChatsInfo{
					ItemType: details.TeamsChat,
					Modified: then,
					Chat: details.ChatInfo{
						Name:               "Hello world",
						LastMessageAt:      then,
						LastMessagePreview: id,
						Members:            []string{},
						MessageCount:       0,
					},
				}

				return chat, i
			},
		},
		{
			name: "chat only, no messages",
			expected: func() (models.Chatable, *details.TeamsChatsInfo) {
				now := time.Now()
				then := now.Add(1 * time.Hour)

				chat := testdata.StubChats(uuid.NewString())[0]
				chat.SetTopic(ptr.To("Hello world"))
				chat.SetCreatedDateTime(&now)
				chat.SetLastUpdatedDateTime(&then)
				chat.SetLastMessagePreview(nil)
				chat.SetMessages(nil)

				i := &details.TeamsChatsInfo{
					ItemType: details.TeamsChat,
					Modified: then,
					Chat: details.ChatInfo{
						Name:               "Hello world",
						LastMessageAt:      time.Time{},
						LastMessagePreview: "",
						Members:            []string{},
						MessageCount:       0,
					},
				}

				return chat, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()
			chat, expected := test.expected()
			result := TeamsChatInfo(chat)

			assert.Equal(t, expected.Chat.Name, result.Chat.Name)

			expectCreated := chat.GetCreatedDateTime()
			if expectCreated != nil {
				assert.Equal(t, ptr.Val(expectCreated), result.Chat.CreatedAt)
			} else {
				assert.True(t, result.Chat.CreatedAt.After(start))
			}

			assert.Truef(
				t,
				expected.Modified.Equal(result.Modified),
				"modified time doesn't match\nexpected %v\ngot %v",
				expected.Modified,
				result.Modified)
		})
	}
}
