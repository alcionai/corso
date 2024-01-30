package api

import (
	"testing"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
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
		name            string
		chatAndExpected func() (models.Chatable, *details.TeamsChatsInfo)
	}{
		{
			name: "Empty chat",
			chatAndExpected: func() (models.Chatable, *details.TeamsChatsInfo) {
				chat := models.NewChat()

				i := &details.TeamsChatsInfo{
					ItemType: details.TeamsChat,
					Chat:     details.ChatInfo{},
				}

				return chat, i
			},
		},
		{
			name: "All fields",
			chatAndExpected: func() (models.Chatable, *details.TeamsChatsInfo) {
				now := time.Now()
				then := now.Add(1 * time.Hour)

				chat := models.NewChat()
				chat.SetTopic(ptr.To("Hello world"))
				chat.SetCreatedDateTime(&now)
				chat.SetLastUpdatedDateTime(&then)

				i := &details.TeamsChatsInfo{
					ItemType: details.TeamsChat,
					Chat: details.ChatInfo{
						Name: "Hello world",
					},
				}

				return chat, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()
			chat, expected := test.chatAndExpected()
			result := TeamsChatInfo(chat)

			assert.Equal(t, expected.Chat.Name, result.Chat.Name)

			expectLastUpdated := chat.GetLastUpdatedDateTime()
			if expectLastUpdated != nil {
				assert.Equal(t, ptr.Val(expectLastUpdated), result.Modified)
			} else {
				assert.True(t, result.Modified.After(start))
			}

			expectCreated := chat.GetCreatedDateTime()
			if expectCreated != nil {
				assert.Equal(t, ptr.Val(expectCreated), result.Chat.CreatedAt)
			} else {
				assert.True(t, result.Chat.CreatedAt.After(start))
			}
		})
	}
}
