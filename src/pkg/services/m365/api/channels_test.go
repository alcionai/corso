package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/alcionai/clues"
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

func (suite *ChannelsAPIUnitSuite) TestStripChatMessageContent() {
	attach1 := models.NewChatMessageAttachment()
	attach1.SetId(ptr.To("id1"))
	attach1.SetName(ptr.To("a1"))

	attach2 := models.NewChatMessageAttachment()
	attach2.SetId(ptr.To("id2"))
	attach2.SetName(ptr.To("a2"))

	attachments := []models.ChatMessageAttachmentable{attach1, attach2}

	realAttach := models.NewChatMessageAttachment()
	realAttach.SetId(ptr.To("1693946862569"))
	realAttach.SetName(ptr.To("meeting"))

	realAttach2 := models.NewChatMessageAttachment()
	realAttach2.SetId(ptr.To("c285cc11-a1ba-473f-ba49-433696e93ef5"))
	realAttach2.SetName(ptr.To("thankyoublue.jpg"))

	realAttachments := []models.ChatMessageAttachmentable{realAttach, realAttach2}

	attachML := func(id string) string {
		return fmt.Sprintf(`<attachment id="%s"></attachment>`, id)
	}

	tests := []struct {
		name        string
		content     string
		attachments []models.ChatMessageAttachmentable
		expect      string
		expectErr   assert.ErrorAssertionFunc
	}{
		{
			name:        "empty content",
			content:     "",
			attachments: attachments,
			expect:      "",
			expectErr:   assert.NoError,
		},
		{
			name:        "only attachment",
			content:     attachML("id1"),
			attachments: attachments,
			expect:      "[attachment:a1]",
			expectErr:   assert.NoError,
		},
		{
			name:        "unknown attachment",
			content:     attachML("idX"),
			attachments: attachments,
			expect:      "[attachment]",
			expectErr:   assert.NoError,
		},
		{
			name:        "text and attachment",
			content:     "some text" + attachML("id1") + "other text",
			attachments: attachments,
			expect:      "some text[attachment:a1]other text",
			expectErr:   assert.NoError,
		},
		{
			name:        "multiple attachments",
			content:     attachML("id1") + attachML("id2"),
			attachments: attachments,
			expect:      "[attachment:a1][attachment:a2]",
			expectErr:   assert.NoError,
		},
		{
			name:        "multiple attachments with unidentified",
			content:     attachML("id1") + attachML("id2") + attachML("idX"),
			attachments: attachments,
			expect:      "[attachment:a1][attachment:a2][attachment]",
			expectErr:   assert.NoError,
		},
		{
			name:        "with empty html",
			content:     "<body><div></div></body>",
			attachments: attachments,
			expect:      "",
			expectErr:   assert.NoError,
		},
		{
			name:        "with malformed div",
			content:     "<body>body<div/>end</body>",
			attachments: attachments,
			expect:      "body\nend",
			expectErr:   assert.NoError,
		},
		{
			name:        "with malformed html 2",
			content:     "<body>body<p>inner</div>end</body>",
			attachments: attachments,
			expect:      "body\n\ninnerend",
			expectErr:   assert.NoError,
		},
		{
			name:        "with html",
			content:     "<body>body<div>in the div</div>end</body>",
			attachments: attachments,
			expect:      "body\nin the div\nend",
			expectErr:   assert.NoError,
		},
		{
			name:        "with html and attachments",
			content:     "<body>body<div>" + attachML("id1") + attachML("id2") + attachML("idX") + "</div>end</body>",
			attachments: attachments,
			expect:      "body\n[attachment:a1][attachment:a2][attachment]\nend",
			expectErr:   assert.NoError,
		},
		{
			name:        "real world example 1",
			content:     "Scheduled a meeting<attachment id=\"1693946862569\"></attachment>",
			attachments: realAttachments,
			expect:      "Scheduled a meeting[attachment:meeting]",
			expectErr:   assert.NoError,
		},
		{
			name: "real world example 2",
			//nolint:lll
			content:     "<p>gifzz</p>\n<img alt=\"The Pumpkin Dance Dancing GIF by Halloween (GIF Image)\" height=\"240\" width=\"320\" src=\"https://media1.giphy.com/media/3oz8xwooUvMqNB1zEs/giphy.gif?cid=ae9bf95er6q1rr4j2btels96zvkiv5q7ylj6spnoekvbhe4p&amp;ep=v1_gifs_trending&amp;rid=giphy.gif&amp;ct=g\" style=\"padding-top:5px\">\n<p>&nbsp;</p>",
			attachments: attachments,
			expect:      "gifzz",
			expectErr:   assert.NoError,
		},
		{
			name: "real world example 3",
			//nolint:lll
			content:     "<p>one</p>\n<img alt=\"Cracking Up Lol GIF by Rodney Dangerfield (GIF Image)\" height=\"250\" width=\"250\" src=\"https://media4.giphy.com/media/l3fQf1OEAq0iri9RC/giphy.gif?cid=de9bf95exzvk1yspjkyg8vh903jyi5yuio8cm7hx4bm1062q&amp;ep=v1_gifs_trending&amp;rid=giphy.gif&amp;ct=g\" style=\"padding-top:5px\">\n<p>attach</p>\n<img alt=\"Sticker image, OH YEAH!\" height=\"250\" width=\"334\" src=\"https://graph.microsoft.com/v1.0/teams/1e8f7352-5393-4b37-b8ab-86ccded848ae/channels/19:vXykALk6NNnBYOkIAqJ6Shs9NuiW3qAuC-1cjprE_X41@thread.tacv2/messages/1698360469770/hostedContents/aWQ9eF8wLWN1cy1kMTItMTNmYzRlNDQwYWNhMjM1MmI2NTE1ZmMxZGI1ZmRlNjgsdHlwZT0xLHVybD1odHRwczovL3VzLWFwaS5hc20uc2t5cGUuY29tL3YxL29iamVjdHMvMC1jdXMtZDEyLTEzZmM0ZTQ0MGFjYTIzNTJiNjUxNWZjMWRiNWZkZTY4L3ZpZXdzL2ltZ28=/$value\" itemid=\"0-cus-d12-13fc4e440aca2352b6515fc1db5fde68\"><span style=\"font-size:0\">OH YEAH!</span>\n<p><span title=\"Party popper\" type=\"(1f389_partypopper)\" class=\"animated-emoticon-20-1f389_partypopper\"><img itemscope=\"\" itemtype=\"http://schema.skype.com/Emoji\" itemid=\"1f389_partypopper\" src=\"https://statics.teams.cdn.office.net/evergreen-assets/personal-expressions/v2/assets/emoticons/1f389_partypopper/default/20_f.png\" title=\"Party popper\" alt=\"🎉\" style=\"width:20px; height:20px\"></span></p>\n<img alt=\"Oh My God Omg GIF by Instacart (GIF Image)\" height=\"250\" width=\"250\" src=\"https://media3.giphy.com/media/8KFTElqwyGHIV4qLqo/giphy.gif?cid=de9bf95exzvk1yspjkyg8vh903jyi5yuio8cm7hx4bm1062q&amp;ep=v1_gifs_trending&amp;rid=giphy.gif&amp;ct=g\" style=\"padding-top:5px\">\n<p>&nbsp;</p><attachment id=\"c285cc11-a1ba-473f-ba49-433696e93ef5\"></attachment>",
			attachments: realAttachments,
			expect:      "one\n\nattach\n\nOH YEAH!\n\n[attachment:thankyoublue.jpg]",
			expectErr:   assert.NoError,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			msg := models.NewChatMessage()
			body := models.NewItemBody()
			body.SetContent(ptr.To(test.content))
			msg.SetBody(body)
			msg.SetAttachments(test.attachments)

			// not testing len; it's effectively covered by the content assertion
			result, _, err := stripChatMessageHTML(msg)
			assert.Equal(t, test.expect, result)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
