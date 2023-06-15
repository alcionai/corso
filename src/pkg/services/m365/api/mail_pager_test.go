package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type MailPagerIntgSuite struct {
	tester.Suite
	cts clientTesterSetup
}

func TestMailPagerIntgSuite(t *testing.T) {
	suite.Run(t, &MailPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *MailPagerIntgSuite) SetupSuite() {
	suite.cts = newClientTesterSetup(suite.T())
}

func (suite *MailPagerIntgSuite) TestGetItemsInContainerByCollisionKey() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	results, err := suite.cts.ac.Mail().GetItemsInContainerByCollisionKey(ctx, suite.cts.userID, "inbox")
	require.NoError(t, err, clues.ToCore(err))
	require.Less(t, 0, len(results), "requires at least one result")

	for k, v := range results {
		assert.NotEmpty(t, k, "all keys should be populated")
		assert.NotEmpty(t, v, "all values should be populated")
	}
}
