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

func TestChannelsAPIUnitSuite(t *testing.T) {
	suite.Run(t, &ChannelsAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ChannelsAPIUnitSuite) TestChannelMessageInfo() {
	var (
		initial = time.Now().Add(-24 * time.Hour)
		mid     = time.Now().Add(-1 * time.Hour)
		curr    = time.Now()
	)

	var (
		content      = "content"
		body         = models.NewItemBody()
		replyContent = "replycontent"
		replyBody    = models.NewItemBody()
	)

	body.SetContent(ptr.To(content))
	replyBody.SetContent(ptr.To(replyContent))

	var (
		attach1      = models.NewChatMessageAttachment()
		attach2      = models.NewChatMessageAttachment()
		replyAttach1 = models.NewChatMessageAttachment()
		replyAttach2 = models.NewChatMessageAttachment()
	)

	attach1.SetName(ptr.To("attach1.ment"))
	attach2.SetName(ptr.To("attach2.ment"))
	replyAttach1.SetName(ptr.To("replyattach1.ment"))
	replyAttach2.SetName(ptr.To("replyattach2.ment"))

	var (
		attachments            = []models.ChatMessageAttachmentable{attach1, attach2}
		replyAttachments       = []models.ChatMessageAttachmentable{replyAttach1, replyAttach2}
		expectAttachNames      = []string{"attach1.ment", "attach2.ment"}
		expectReplyAttachNames = []string{"replyattach1.ment", "replyattach2.ment"}
	)

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
				msg.SetSubject(ptr.To("subject"))

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				i := &details.GroupsInfo{
					ItemType:  details.GroupsChannelMessage,
					Modified:  initial,
					LastReply: details.ChannelMessageInfo{},
					Message: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       initial,
						Creator:         "user",
						ReplyCount:      0,
						Preview:         "",
						Size:            0,
						Subject:         "subject",
					},
				}

				return msg, i
			},
		},
		{
			name: "No Subject",
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
					ItemType:  details.GroupsChannelMessage,
					Modified:  initial,
					LastReply: details.ChannelMessageInfo{},
					Message: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       initial,
						Creator:         "user",
						ReplyCount:      0,
						Preview:         content,
						Size:            int64(len(content)),
						Subject:         "",
					},
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
				msg.SetSubject(ptr.To("subject"))

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				i := &details.GroupsInfo{
					ItemType:  details.GroupsChannelMessage,
					Modified:  initial,
					LastReply: details.ChannelMessageInfo{},
					Message: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       initial,
						Creator:         "user",
						ReplyCount:      0,
						Preview:         content,
						Size:            int64(len(content)),
						Subject:         "subject",
					},
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
				msg.SetSubject(ptr.To("subject"))

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("app"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetApplication(iden)

				msg.SetFrom(from)

				i := &details.GroupsInfo{
					ItemType:  details.GroupsChannelMessage,
					Modified:  initial,
					LastReply: details.ChannelMessageInfo{},
					Message: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       initial,
						Creator:         "app",
						ReplyCount:      0,
						Preview:         content,
						Size:            int64(len(content)),
						Subject:         "subject",
					},
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
				msg.SetSubject(ptr.To("subject"))

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("device"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetDevice(iden)

				msg.SetFrom(from)

				i := &details.GroupsInfo{
					ItemType:  details.GroupsChannelMessage,
					Modified:  initial,
					LastReply: details.ChannelMessageInfo{},
					Message: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       initial,
						Creator:         "device",
						ReplyCount:      0,
						Preview:         content,
						Size:            int64(len(content)),
						Subject:         "subject",
					},
				}

				return msg, i
			},
		},
		{
			name: "No Replies - with attachments",
			msgAndInfo: func() (models.ChatMessageable, *details.GroupsInfo) {
				msg := models.NewChatMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetBody(body)
				msg.SetSubject(ptr.To("subject"))
				msg.SetAttachments(attachments)

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				i := &details.GroupsInfo{
					ItemType:  details.GroupsChannelMessage,
					Modified:  initial,
					LastReply: details.ChannelMessageInfo{},
					Message: details.ChannelMessageInfo{
						AttachmentNames: expectAttachNames,
						CreatedAt:       initial,
						Creator:         "user",
						ReplyCount:      0,
						Preview:         content,
						Size:            int64(len(content)),
						Subject:         "subject",
					},
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
				msg.SetSubject(ptr.To("subject"))

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				// reply

				iden = models.NewIdentity()
				iden.SetDisplayName(ptr.To("replyuser"))

				from = models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				reply := models.NewChatMessage()
				reply.SetCreatedDateTime(&curr)
				reply.SetLastModifiedDateTime(&curr)
				reply.SetFrom(from)
				reply.SetBody(replyBody)

				msg.SetReplies([]models.ChatMessageable{reply})

				i := &details.GroupsInfo{
					ItemType: details.GroupsChannelMessage,
					Modified: curr,
					LastReply: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       curr,
						Creator:         "replyuser",
						ReplyCount:      0,
						Preview:         replyContent,
						Size:            int64(len(replyContent)),
					},
					Message: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       initial,
						Creator:         "user",
						ReplyCount:      1,
						Preview:         content,
						Size:            int64(len(content)),
						Subject:         "subject",
					},
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
				msg.SetSubject(ptr.To("subject"))

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				// replies

				iden = models.NewIdentity()
				iden.SetDisplayName(ptr.To("reply1user"))

				from = models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				reply1 := models.NewChatMessage()
				reply1.SetCreatedDateTime(&mid)
				reply1.SetLastModifiedDateTime(&mid)
				reply1.SetFrom(from)
				reply1.SetBody(replyBody)

				iden = models.NewIdentity()
				iden.SetDisplayName(ptr.To("reply2user"))

				from = models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				reply2 := models.NewChatMessage()
				reply2.SetCreatedDateTime(&curr)
				reply2.SetLastModifiedDateTime(&curr)
				reply2.SetFrom(from)
				reply2.SetBody(replyBody)

				msg.SetReplies([]models.ChatMessageable{reply1, reply2})

				i := &details.GroupsInfo{
					ItemType: details.GroupsChannelMessage,
					Modified: curr,
					LastReply: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       curr,
						Creator:         "reply2user",
						ReplyCount:      0,
						Preview:         replyContent,
						Size:            int64(len(replyContent)),
					},
					Message: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       initial,
						Creator:         "user",
						ReplyCount:      2,
						Preview:         content,
						Size:            int64(len(content)),
						Subject:         "subject",
					},
				}

				return msg, i
			},
		},
		{
			name: "Many Replies - not last has attachments",
			msgAndInfo: func() (models.ChatMessageable, *details.GroupsInfo) {
				msg := models.NewChatMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetBody(body)
				msg.SetSubject(ptr.To("subject"))

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				// replies

				iden = models.NewIdentity()
				iden.SetDisplayName(ptr.To("reply1user"))

				from = models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				reply1 := models.NewChatMessage()
				reply1.SetCreatedDateTime(&mid)
				reply1.SetLastModifiedDateTime(&mid)
				reply1.SetFrom(from)
				reply1.SetBody(replyBody)
				reply1.SetAttachments(replyAttachments)

				iden = models.NewIdentity()
				iden.SetDisplayName(ptr.To("reply2user"))

				from = models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				reply2 := models.NewChatMessage()
				reply2.SetCreatedDateTime(&curr)
				reply2.SetLastModifiedDateTime(&curr)
				reply2.SetFrom(from)
				reply2.SetBody(replyBody)

				msg.SetReplies([]models.ChatMessageable{reply1, reply2})

				i := &details.GroupsInfo{
					ItemType: details.GroupsChannelMessage,
					Modified: curr,
					LastReply: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       curr,
						Creator:         "reply2user",
						ReplyCount:      0,
						Preview:         replyContent,
						Size:            int64(len(replyContent)),
					},
					Message: details.ChannelMessageInfo{
						AttachmentNames: []string{},
						CreatedAt:       initial,
						Creator:         "user",
						ReplyCount:      2,
						Preview:         content,
						Size:            int64(len(content)),
						Subject:         "subject",
					},
				}

				return msg, i
			},
		},
		{
			name: "Many Replies - last has attachments",
			msgAndInfo: func() (models.ChatMessageable, *details.GroupsInfo) {
				msg := models.NewChatMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetBody(body)
				msg.SetSubject(ptr.To("subject"))
				msg.SetAttachments(attachments)

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				from := models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				msg.SetFrom(from)

				// replies

				iden = models.NewIdentity()
				iden.SetDisplayName(ptr.To("reply1user"))

				from = models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				reply1 := models.NewChatMessage()
				reply1.SetCreatedDateTime(&mid)
				reply1.SetLastModifiedDateTime(&mid)
				reply1.SetFrom(from)
				reply1.SetBody(replyBody)

				iden = models.NewIdentity()
				iden.SetDisplayName(ptr.To("reply2user"))

				from = models.NewChatMessageFromIdentitySet()
				from.SetUser(iden)

				reply2 := models.NewChatMessage()
				reply2.SetCreatedDateTime(&curr)
				reply2.SetLastModifiedDateTime(&curr)
				reply2.SetFrom(from)
				reply2.SetBody(replyBody)
				reply2.SetAttachments(replyAttachments)

				msg.SetReplies([]models.ChatMessageable{reply1, reply2})

				i := &details.GroupsInfo{
					ItemType: details.GroupsChannelMessage,
					Modified: curr,
					LastReply: details.ChannelMessageInfo{
						AttachmentNames: expectReplyAttachNames,
						CreatedAt:       curr,
						Creator:         "reply2user",
						ReplyCount:      0,
						Preview:         replyContent,
						Size:            int64(len(replyContent)),
					},
					Message: details.ChannelMessageInfo{
						AttachmentNames: expectAttachNames,
						CreatedAt:       initial,
						Creator:         "user",
						ReplyCount:      2,
						Preview:         content,
						Size:            int64(len(content)),
						Subject:         "subject",
					},
				}

				return msg, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			chMsg, expected := test.msgAndInfo()
			result := channelMessageInfo(chMsg)

			ma := result.Message.AttachmentNames
			result.Message.AttachmentNames = nil
			ema := expected.Message.AttachmentNames
			expected.Message.AttachmentNames = nil

			lra := result.LastReply.AttachmentNames
			result.LastReply.AttachmentNames = nil
			elra := expected.LastReply.AttachmentNames
			expected.LastReply.AttachmentNames = nil

			assert.Equal(t, expected, result)
			assert.ElementsMatch(t, ema, ma)
			assert.ElementsMatch(t, elra, lra)
		})
	}
}
