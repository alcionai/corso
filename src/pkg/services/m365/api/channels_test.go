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

type ChannelsAPIUnitSuite struct {
	tester.Suite
}

func TestChannelsAPIUnitSuitee(t *testing.T) {
	suite.Run(t, &ChannelsAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ChannelsAPIUnitSuite) TestChannelMessageInfo() {
	var (
		initial = time.Now().Add(-24 * time.Hour)
		mid     = time.Now().Add(-1 * time.Hour)
		curr    = time.Now()

		content = "content"
		body    = models.NewItemBody()
	)

	body.SetContent(ptr.To(content))

	tests := []struct {
		name       string
		msgAndInfo func() (models.ChatMessageable, *details.GroupsInfo)
	}{
		{
			name: "No body",
			msgAndInfo: func() (models.ChatMessageable, *details.GroupsInfo) {
				msg := models.NewChatMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				i := &details.GroupsInfo{
					ItemType:       details.GroupsChannelMessage,
					Created:        initial,
					Modified:       initial,
					LastReplyAt:    time.Time{},
					ReplyCount:     0,
					MessageCreator: "user",
				}

				return msg, i
			},
		},
		{
			name: "No Replies - created by user",
			msgAndInfo: func() (models.ChatMessageable, *details.GroupsInfo) {
				msg := models.NewChatMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetBody(body)

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				i := &details.GroupsInfo{
					ItemType:       details.GroupsChannelMessage,
					Created:        initial,
					Modified:       initial,
					LastReplyAt:    time.Time{},
					ReplyCount:     0,
					MessageCreator: "user",
					Size:           int64(len(content)),
					MessagePreview: content,
				}

				return msg, i
			},
		},
		{
			name: "No Replies - created by application",
			msgAndInfo: func() (models.ChatMessageable, *details.GroupsInfo) {
				msg := models.NewChatMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetBody(body)

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("app"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetApplication(iden)

				msg.SetFrom(from)

				i := &details.GroupsInfo{
					ItemType:       details.GroupsChannelMessage,
					Created:        initial,
					Modified:       initial,
					LastReplyAt:    time.Time{},
					ReplyCount:     0,
					MessageCreator: "app",
					Size:           int64(len(content)),
					MessagePreview: content,
				}

				return msg, i
			},
		},
		{
			name: "No Replies - created by device",
			msgAndInfo: func() (models.ChatMessageable, *details.GroupsInfo) {
				msg := models.NewChatMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetBody(body)

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("device"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetDevice(iden)

				msg.SetFrom(from)

				i := &details.GroupsInfo{
					ItemType:       details.GroupsChannelMessage,
					Created:        initial,
					Modified:       initial,
					LastReplyAt:    time.Time{},
					ReplyCount:     0,
					MessageCreator: "device",
					Size:           int64(len(content)),
					MessagePreview: content,
				}

				return msg, i
			},
		},
		{
			name: "One Reply",
			msgAndInfo: func() (models.ChatMessageable, *details.GroupsInfo) {
				msg := models.NewChatMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetBody(body)

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				reply := models.NewChatMessage()
				reply.SetCreatedDateTime(&curr)
				reply.SetLastModifiedDateTime(&curr)

				msg.SetReplies([]models.ChatMessageable{reply})

				i := &details.GroupsInfo{
					ItemType:       details.GroupsChannelMessage,
					Created:        initial,
					Modified:       curr,
					LastReplyAt:    curr,
					ReplyCount:     1,
					MessageCreator: "user",
					Size:           int64(len(content)),
					MessagePreview: content,
				}

				return msg, i
			},
		},
		{
			name: "Many Replies",
			msgAndInfo: func() (models.ChatMessageable, *details.GroupsInfo) {
				msg := models.NewChatMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetBody(body)

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				reply1 := models.NewChatMessage()
				reply1.SetCreatedDateTime(&mid)
				reply1.SetLastModifiedDateTime(&mid)

				reply2 := models.NewChatMessage()
				reply2.SetCreatedDateTime(&curr)
				reply2.SetLastModifiedDateTime(&curr)

				msg.SetReplies([]models.ChatMessageable{reply1, reply2})

				i := &details.GroupsInfo{
					ItemType:       details.GroupsChannelMessage,
					Created:        initial,
					Modified:       curr,
					LastReplyAt:    curr,
					ReplyCount:     2,
					MessageCreator: "user",
					Size:           int64(len(content)),
					MessagePreview: content,
				}

				return msg, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			chMsg, expected := test.msgAndInfo()
			assert.Equal(suite.T(), expected, ChannelMessageInfo(chMsg))
		})
	}
}
