package api

import (
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

type ConversationsPagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestConversationsPagerIntgSuite(t *testing.T) {
	suite.Run(t, &ConversationsPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ConversationsPagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *ConversationsPagerIntgSuite) TestEnumerateConversations_withThreadsAndPosts() {
	var (
		t  = suite.T()
		ac = suite.its.ac.Conversations()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	convs, err := ac.GetConversations(ctx, suite.its.group.id, CallConfig{})
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, convs)

	for _, conv := range convs {
		threads := testEnumerateConvThreads(suite, conv)

		for _, thread := range threads {
			posts := testEnumerateConvPosts(suite, conv, thread)

			aar, err := ac.GetConversationThreadPostIDs(
				ctx,
				suite.its.group.id,
				ptr.Val(conv.GetId()),
				ptr.Val(thread.GetId()),
				CallConfig{})
			require.NoError(t, err, clues.ToCore(err))
			require.Equal(t, len(posts), len(aar.Added), "added the same number of ids and posts")
			assert.True(t, aar.ValidModTimes, "mod times should be valid")
			assert.Empty(t, aar.Removed, "no items should get removed")
			assert.Empty(t, aar.DU.URL, "no delta update token should be provided")

			for _, post := range posts {
				testGetPostByID(suite, conv, thread, post)
			}
		}
	}
}

func testEnumerateConvThreads(
	suite *ConversationsPagerIntgSuite,
	conv models.Conversationable,
) []models.ConversationThreadable {
	var threads []models.ConversationThreadable

	suite.Run("threads", func() {
		var (
			t   = suite.T()
			ac  = suite.its.ac.Conversations()
			err error
		)

		ctx, flush := tester.NewContext(t)
		defer flush()

		threads, err = ac.GetConversationThreads(
			ctx,
			suite.its.group.id,
			ptr.Val(conv.GetId()),
			CallConfig{})
		require.NoError(t, err, clues.ToCore(err))
		// to the best of our knowledge, there's only ever one
		// thread per conversation.  Even trying to create a new
		// thread within a conversation will create an entirely
		// new conversation.  We want this test to fail as a potential
		// identifier if that changes on us.
		require.Equal(t, 1, len(threads))
	})

	return threads
}

func testEnumerateConvPosts(
	suite *ConversationsPagerIntgSuite,
	conv models.Conversationable,
	thread models.ConversationThreadable,
) []models.Postable {
	var posts []models.Postable

	suite.Run("posts", func() {
		var (
			t   = suite.T()
			ac  = suite.its.ac.Conversations()
			err error
		)

		ctx, flush := tester.NewContext(t)
		defer flush()

		posts, err = ac.GetConversationThreadPosts(
			ctx,
			suite.its.group.id,
			ptr.Val(conv.GetId()),
			ptr.Val(thread.GetId()),
			CallConfig{})
		require.NoError(t, err, clues.ToCore(err))
		require.NotEmpty(t, posts)
	})

	return posts
}
