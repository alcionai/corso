package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ContactsPagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestContactsPagerIntgSuite(t *testing.T) {
	suite.Run(t, &ContactsPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ContactsPagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *ContactsPagerIntgSuite) TestContacts_GetItemsInContainerByCollisionKey() {
	t := suite.T()
	ac := suite.its.ac.Contacts()

	ctx, flush := tester.NewContext(t)
	defer flush()

	container, err := ac.GetContainerByID(ctx, suite.its.userID, "contacts")
	require.NoError(t, err, clues.ToCore(err))

	conts, err := ac.Stable.
		Client().
		Users().
		ByUserId(suite.its.userID).
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

	results, err := suite.its.ac.Contacts().GetItemsInContainerByCollisionKey(ctx, suite.its.userID, "contacts")
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
