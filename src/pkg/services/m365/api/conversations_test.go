package api

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// called by the pager test, since it is already enumerating
// posts.
func testGetPostByID(
	suite *ConversationsPagerIntgSuite,
	conv models.Conversationable,
	thread models.ConversationThreadable,
	post models.Postable,
) {
	suite.Run("post_by_id", func() {
		var (
			t  = suite.T()
			ac = suite.its.ac.Conversations()
		)

		ctx, flush := tester.NewContext(t)
		defer flush()

		resp, _, err := ac.GetConversationPost(
			ctx,
			suite.its.group.id,
			ptr.Val(conv.GetId()),
			ptr.Val(thread.GetId()),
			ptr.Val(post.GetId()),
			CallConfig{})
		require.NoError(t, err, clues.ToCore(err))
		require.Equal(t, ptr.Val(post.GetId()), ptr.Val(resp.GetId()))
	})
}

type ConversationsAPIUnitSuite struct {
	tester.Suite
}

func TestConversationsAPIUnitSuite(t *testing.T) {
	suite.Run(t, &ConversationsAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ConversationsAPIUnitSuite) TestConversationPostInfo() {
	var (
		now     = time.Now()
		content = "content"
		body    = models.NewItemBody()
	)

	body.SetContent(ptr.To(content))

	tests := []struct {
		name        string
		postAndInfo func() (models.Postable, *details.GroupsInfo)
	}{
		{
			name: "No body",
			postAndInfo: func() (models.Postable, *details.GroupsInfo) {
				post := models.NewPost()
				post.SetCreatedDateTime(&now)
				post.SetLastModifiedDateTime(&now)

				sender := "foo@bar.com"
				sea := models.NewEmailAddress()
				sea.SetAddress(&sender)

				recipient := models.NewRecipient()
				recipient.SetEmailAddress(sea)

				post.SetSender(recipient)

				iden := models.NewIdentity()
				iden.SetDisplayName(ptr.To("user"))

				i := &details.GroupsInfo{
					ItemType: details.GroupsConversationPost,
					Modified: now,
					Post: details.ConversationPostInfo{
						CreatedAt: now,
						Creator:   "foo@bar.com",
						Preview:   "",
						Size:      0,
						// TODO: feed the subject in from the conversation
						Subject: "",
					},
				}

				return post, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			chMsg, expected := test.postAndInfo()
			result := conversationPostInfo(chMsg)

			assert.Equal(t, expected, result)
		})
	}
}

type ConversationAPIIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

// We do end up mocking the actual request, but creating the rest
// similar to full integration tests.
func TestConversationAPIIntgSuite(t *testing.T) {
	suite.Run(t, &ConversationAPIIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ConversationAPIIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *ConversationAPIIntgSuite) TestConversations_attachmentListDownload() {
	pid := "fake-post-id"
	aid := "fake-attachment-id"

	tests := []struct {
		name            string
		setupf          func()
		attachmentCount int
		size            int64
		expect          assert.ErrorAssertionFunc
	}{
		{
			name: "no attachments",
			setupf: func() {
				itm := models.NewPost()
				itm.SetId(&pid)

				interceptV1Path(
					"groups",
					"group",
					"conversations",
					"conv",
					"threads",
					"thread",
					"posts",
					pid).
					Reply(200).
					JSON(requireParseableToMap(suite.T(), itm))
			},
			expect: assert.NoError,
		},
		{
			name: "fetch with attachment",
			setupf: func() {
				itm := models.NewPost()
				itm.SetId(&pid)
				itm.SetHasAttachments(ptr.To(true))

				attch := models.NewAttachment()
				attch.SetSize(ptr.To[int32](50))

				itm.SetAttachments([]models.Attachmentable{attch})

				interceptV1Path(
					"groups",
					"group",
					"conversations",
					"conv",
					"threads",
					"thread",
					"posts",
					pid).
					Reply(200).
					JSON(requireParseableToMap(suite.T(), itm))

			},
			attachmentCount: 1,
			size:            50,
			expect:          assert.NoError,
		},
		{
			name: "fetch multiple individual attachments",
			setupf: func() {
				truthy := true
				itm := models.NewPost()
				itm.SetId(&pid)
				itm.SetHasAttachments(&truthy)

				attch := models.NewAttachment()
				attch.SetId(&aid)
				attch.SetSize(ptr.To[int32](200))

				itm.SetAttachments([]models.Attachmentable{attch, attch, attch, attch, attch})

				interceptV1Path(
					"groups",
					"group",
					"conversations",
					"conv",
					"threads",
					"thread",
					"posts",
					pid).
					Reply(200).
					JSON(requireParseableToMap(suite.T(), itm))
			},
			attachmentCount: 5,
			size:            1000,
			expect:          assert.NoError,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			defer gock.Off()
			test.setupf()

			item, _, err := suite.its.gockAC.
				Conversations().
				GetConversationPost(
					ctx,
					"group",
					"conv",
					"thread",
					pid,
					CallConfig{})
			test.expect(t, err)

			var size int64

			if item.GetBody() != nil {
				content := ptr.Val(item.GetBody().GetContent())
				size = int64(len(content))
			}

			attachments := item.GetAttachments()
			for _, attachment := range attachments {
				size += int64(*attachment.GetSize())
			}

			assert.Equal(t, *item.GetId(), pid)
			assert.Equal(t, test.attachmentCount, len(attachments), "attachment count")
			assert.Equal(t, test.size, size, "item size")
			assert.True(t, gock.IsDone(), "made all requests")
		})
	}
}
