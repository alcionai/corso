package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/tester/tsetup"
)

type MailPagerIntgSuite struct {
	tester.Suite
	its tsetup.M365
}

func TestMailPagerIntgSuite(t *testing.T) {
	suite.Run(t, &MailPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *MailPagerIntgSuite) SetupSuite() {
	suite.its = tsetup.NewM365IntegrationTester(suite.T())
}

func (suite *MailPagerIntgSuite) TestMail_GetItemsInContainerByCollisionKey() {
	t := suite.T()
	ac := suite.its.AC.Mail()

	ctx, flush := tester.NewContext(t)
	defer flush()

	container, err := ac.GetContainerByID(ctx, suite.its.User.ID, MailInbox)
	require.NoError(t, err, clues.ToCore(err))

	msgs, err := ac.Stable.
		Client().
		Users().
		ByUserId(suite.its.User.ID).
		MailFolders().
		ByMailFolderId(ptr.Val(container.GetId())).
		Messages().
		Get(ctx, nil)
	require.NoError(t, err, clues.ToCore(err))

	ms := msgs.GetValue()
	expectM := map[string]struct{}{}

	for _, m := range ms {
		expectM[MailCollisionKey(m)] = struct{}{}
	}

	expect := maps.Keys(expectM)

	results, err := suite.its.AC.Mail().GetItemsInContainerByCollisionKey(ctx, suite.its.User.ID, MailInbox)
	require.NoError(t, err, clues.ToCore(err))
	require.Less(t, 0, len(results), "requires at least one result")

	for _, k := range expect {
		t.Log("expects key", k)
	}

	for k := range results {
		t.Log("results key", k)
	}

	for k, v := range results {
		assert.NotEmpty(t, k, "all keys should be populated")
		assert.NotEmpty(t, v, "all values should be populated")
	}

	for _, e := range expect {
		_, ok := results[e]
		assert.Truef(t, ok, "expected results to contain collision key: %s", e)
	}
}

func (suite *MailPagerIntgSuite) TestMail_GetItemsIDsInContainer() {
	t := suite.T()
	ac := suite.its.AC.Mail()

	ctx, flush := tester.NewContext(t)
	defer flush()

	config := &users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersItemMessagesRequestBuilderGetQueryParameters{
			Top: ptr.To[int32](1000),
		},
	}

	msgs, err := ac.Stable.
		Client().
		Users().
		ByUserId(suite.its.User.ID).
		MailFolders().
		ByMailFolderId(MailInbox).
		Messages().
		Get(ctx, config)
	require.NoError(t, err, clues.ToCore(err))

	ms := msgs.GetValue()
	expect := map[string]struct{}{}

	for _, m := range ms {
		expect[ptr.Val(m.GetId())] = struct{}{}
	}

	results, err := suite.its.AC.Mail().
		GetItemIDsInContainer(ctx, suite.its.User.ID, MailInbox)
	require.NoError(t, err, clues.ToCore(err))
	require.Less(t, 0, len(results), "requires at least one result")
	require.Equal(t, len(expect), len(results), "must have same count of items")

	for k := range expect {
		t.Log("expects key", k)
	}

	for k := range results {
		t.Log("results key", k)
	}

	assert.Equal(t, expect, results)
}
