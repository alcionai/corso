package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

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
	expectM := map[string]struct{}{}

	for _, c := range cs {
		expectM[api.ContactCollisionKey(c)] = struct{}{}
	}

	expect := maps.Keys(expectM)

	results, err := suite.its.ac.Contacts().GetItemsInContainerByCollisionKey(ctx, suite.its.userID, "contacts")
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

func (suite *ContactsPagerIntgSuite) TestContacts_GetItemsIDsInContainer() {
	t := suite.T()
	ac := suite.its.ac.Contacts()

	ctx, flush := tester.NewContext(t)
	defer flush()

	container, err := ac.GetContainerByID(ctx, suite.its.userID, api.DefaultContacts)
	require.NoError(t, err, clues.ToCore(err))

	msgs, err := ac.Stable.
		Client().
		Users().
		ByUserId(suite.its.userID).
		ContactFolders().
		ByContactFolderId(ptr.Val(container.GetId())).
		Contacts().
		Get(ctx, nil)
	require.NoError(t, err, clues.ToCore(err))

	ms := msgs.GetValue()
	expect := map[string]struct{}{}

	for _, m := range ms {
		expect[ptr.Val(m.GetId())] = struct{}{}
	}

	results, err := suite.its.ac.Contacts().
		GetItemIDsInContainer(ctx, suite.its.userID, api.DefaultContacts)
	require.NoError(t, err, clues.ToCore(err))
	require.Less(t, 0, len(results), "requires at least one result")
	require.Equal(t, len(expect), len(results), "must have same count of items")

	for _, k := range expect {
		t.Log("expects key", k)
	}

	for k := range results {
		t.Log("results key", k)
	}

	assert.Equal(t, expect, results)
}
