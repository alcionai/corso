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

type ContactsPagerIntgSuite struct {
	tester.Suite
	cts clientTesterSetup
}

func TestContactsPagerIntgSuite(t *testing.T) {
	suite.Run(t, &ContactsPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *ContactsPagerIntgSuite) SetupSuite() {
	suite.cts = newClientTesterSetup(suite.T())
}

func (suite *ContactsPagerIntgSuite) TestGetItemsInContainerByCollisionKey() {
	t := suite.T()
	ac := suite.cts.ac.Contacts()

	ctx, flush := tester.NewContext(t)
	defer flush()

	container, err := ac.GetContainerByID(ctx, suite.cts.userID, "contacts")
	require.NoError(t, err, clues.ToCore(err))

	conts, err := ac.Stable.
		Client().
		Users().
		ByUserId(suite.cts.userID).
		ContactFolders().
		ByContactFolderId(ptr.Val(container.GetId())).
		Contacts().
		Get(ctx, nil)
	require.NoError(t, err, clues.ToCore(err))

	cs := conts.GetValue()
	expect := make([]string, 0, len(cs))

	for _, c := range cs {
		expect = append(expect, api.ContactCollisionKey(c))
	}

	results, err := ac.GetItemsInContainerByCollisionKey(ctx, suite.cts.userID, "contacts")
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
