package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type MailPagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestMailPagerIntgSuite(t *testing.T) {
	suite.Run(t, &MailPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *MailPagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *MailPagerIntgSuite) TestGetItemsInContainerByCollisionKey() {
	t := suite.T()
	ac := suite.its.ac.Mail()

	ctx, flush := tester.NewContext(t)
	defer flush()

	container, err := ac.GetContainerByID(ctx, suite.its.userID, "inbox")
	require.NoError(t, err, clues.ToCore(err))

	msgs, err := ac.Stable.
		Client().
		Users().
		ByUserId(suite.its.userID).
		MailFolders().
		ByMailFolderId(ptr.Val(container.GetId())).
		Messages().
		Get(ctx, nil)
	require.NoError(t, err, clues.ToCore(err))

	ms := msgs.GetValue()
	expect := make([]string, 0, len(ms))

	for _, m := range ms {
		expect = append(expect, api.MailCollisionKey(m))
	}

	results, err := suite.its.ac.Mail().GetItemsInContainerByCollisionKey(ctx, suite.its.userID, "inbox")
	require.NoError(t, err, clues.ToCore(err))
	require.Less(t, 0, len(results), "requires at least one result")

	for k, v := range results {
		assert.NotEmpty(t, k, "all keys should be populated")
		assert.NotEmpty(t, v, "all values should be populated")
	}

	for _, e := range expect {
		_, ok := results[e]
		assert.Truef(t, ok, "expected results to contain collision key: %s", e)
	}
}
