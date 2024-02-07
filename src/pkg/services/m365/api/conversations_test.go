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

	"github.com/alcionai/canario/src/internal/common/ptr"
	exchMock "github.com/alcionai/canario/src/internal/m365/service/exchange/mock"
	stub "github.com/alcionai/canario/src/internal/m365/service/groups/mock"
	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/tester/tconfig"
	"github.com/alcionai/canario/src/pkg/backup/details"
	graphTD "github.com/alcionai/canario/src/pkg/services/m365/api/graph/testdata"
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
			ptr.Val(post.GetId()))
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
						Topic: "",
					},
				}

				return post, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			post, expected := test.postAndInfo()
			result := conversationPostInfo(post, 0, "")

			assert.Equal(t, expected, result)
		})
	}
}

// TestBytesToPostable_InvalidError tests that the error message kiota returns
// for invalid JSON matches what we check for. This helps keep things in sync
// when kiota is updated.
func (suite *MailAPIUnitSuite) TestBytesToPostable_InvalidError() {
	t := suite.T()
	input := exchMock.MessageWithSpecialCharacters("m365 mail support test")

	_, err := CreateFromBytes(input, models.CreatePostFromDiscriminatorValue)
	require.Error(t, err, clues.ToCore(err))

	assert.Contains(t, err.Error(), invalidJSON)
}

func (suite *ConversationsAPIUnitSuite) TestBytesToPostable() {
	table := []struct {
		name        string
		byteArray   []byte
		checkError  assert.ErrorAssertionFunc
		checkObject assert.ValueAssertionFunc
	}{
		{
			name:        "Empty Bytes",
			byteArray:   make([]byte, 0),
			checkError:  assert.Error,
			checkObject: assert.Nil,
		},
		{
			name: "post bytes",
			// Note: inReplyTo is not serialized or deserialized by kiota so we can't
			// test that aspect. The payload does contain inReplyTo data for future use.
			byteArray:   []byte(stub.PostWithAttachments),
			checkError:  assert.NoError,
			checkObject: assert.NotNil,
		},
		// Using test data from exchMock package for these tests because posts are
		// essentially email messages.
		{
			name:        "malformed JSON bytes passes sanitization",
			byteArray:   exchMock.MessageWithSpecialCharacters("m365 mail support test"),
			checkError:  assert.NoError,
			checkObject: assert.NotNil,
		},
		{
			name: "invalid JSON bytes",
			byteArray: append(
				exchMock.MessageWithSpecialCharacters("m365 mail support test"),
				[]byte("}")...),
			checkError:  assert.Error,
			checkObject: assert.Nil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := BytesToPostable(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.checkObject(t, result)
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

func (suite *ConversationAPIIntgSuite) TestGetConversationPost() {
	pid := "fake-post-id"
	replyToID := "fake-reply-to-id"
	aid := "fake-attachment-id"
	contentWithAttachment := "<html><body><img src=\"cid:fake-attach-id\"/></body></html>"

	interceptPathWithExpand := func(
		item *models.Post,
		expandProperty string,
		path ...string,
	) *gock.Response {
		return interceptV1Path(
			"groups",
			"group",
			"conversations",
			"conv",
			"threads",
			"thread",
			"posts",
			pid).
			MatchParam("$expand", expandProperty).
			Reply(200).
			JSON(graphTD.ParseableToMap(suite.T(), item))
	}

	tests := []struct {
		name             string
		setupf           func()
		attachmentCount  int
		inReplyToPresent bool
		size             int64
		expect           assert.ErrorAssertionFunc
	}{
		{
			name: "no inReplyTo, no attachment",
			setupf: func() {
				item := models.NewPost()
				item.SetId(&pid)

				interceptPathWithExpand(item, "inReplyTo")
			},
			expect: assert.NoError,
		},
		{
			name: "with inReplyTo, no attachment",
			setupf: func() {
				item := models.NewPost()
				parentPost := models.NewPost()

				item.SetId(&pid)
				parentPost.SetId(&replyToID)
				item.SetInReplyTo(parentPost)

				interceptPathWithExpand(item, "inReplyTo")
			},
			inReplyToPresent: true,
			expect:           assert.NoError,
		},
		{
			name: "no inreplyTo, with attachment",
			setupf: func() {
				item := models.NewPost()
				item.SetId(&pid)
				item.SetHasAttachments(ptr.To(true))

				interceptPathWithExpand(item, "inReplyTo")

				attch := models.NewAttachment()
				attch.SetSize(ptr.To[int32](50))
				item.SetAttachments([]models.Attachmentable{attch})

				interceptPathWithExpand(item, "attachments")
			},
			attachmentCount: 1,
			size:            50,
			expect:          assert.NoError,
		},
		{
			name: "with inreplyTo, with attachment",
			setupf: func() {
				item := models.NewPost()
				parentPost := models.NewPost()

				item.SetId(&pid)
				parentPost.SetId(&replyToID)
				item.SetInReplyTo(parentPost)
				item.SetHasAttachments(ptr.To(true))

				interceptPathWithExpand(item, "inReplyTo")

				attch := models.NewAttachment()
				attch.SetSize(ptr.To[int32](50))
				item.SetAttachments([]models.Attachmentable{attch})

				interceptPathWithExpand(item, "attachments")
			},
			inReplyToPresent: true,
			attachmentCount:  1,
			size:             50,
			expect:           assert.NoError,
		},
		// At this point we have tested inReplyTo behavior thoroughly.
		// Skip for remaining tests.
		{
			name: "fetch multiple attachments",
			setupf: func() {
				item := models.NewPost()

				item.SetId(&pid)
				item.SetHasAttachments(ptr.To(true))

				interceptPathWithExpand(item, "inReplyTo")

				attch := models.NewAttachment()
				attch.SetId(&aid)
				attch.SetSize(ptr.To[int32](200))

				item.SetAttachments([]models.Attachmentable{attch, attch, attch, attch, attch})

				interceptPathWithExpand(item, "attachments")
			},
			attachmentCount: 5,
			size:            1000,
			expect:          assert.NoError,
		},
		{
			name: "embedded attachment",
			setupf: func() {
				item := models.NewPost()
				item.SetId(&pid)

				body := models.NewItemBody()
				body.SetContentType(ptr.To(models.HTML_BODYTYPE))
				body.SetContent(ptr.To(contentWithAttachment))
				item.SetBody(body)

				interceptPathWithExpand(item, "inReplyTo")

				attch := models.NewAttachment()
				attch.SetSize(ptr.To[int32](50))
				item.SetAttachments([]models.Attachmentable{attch})

				interceptPathWithExpand(item, "attachments")
			},
			attachmentCount: 1,
			size:            50 + int64(len(contentWithAttachment)),
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
					pid)
			test.expect(t, err)

			// inReplyTo checks
			if test.inReplyToPresent {
				require.NotNil(t, item.GetInReplyTo())
				assert.Equal(t, replyToID, ptr.Val(item.GetInReplyTo().GetId()))
			}

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
