package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/config"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type EventsPagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestEventsPagerIntgSuite(t *testing.T) {
	suite.Run(t, &EventsPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{config.M365AcctCredEnvs}),
	})
}

func (suite *EventsPagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *EventsPagerIntgSuite) TestEvents_GetItemsInContainerByCollisionKey() {
	t := suite.T()
	ac := suite.its.ac.Events()

	ctx, flush := tester.NewContext(t)
	defer flush()

	container, err := ac.GetContainerByID(ctx, suite.its.userID, "calendar")
	require.NoError(t, err, clues.ToCore(err))

	evts, err := ac.Stable.
		Client().
		Users().
		ByUserId(suite.its.userID).
		Calendars().
		ByCalendarId(ptr.Val(container.GetId())).
		Events().
		Get(ctx, nil)
	require.NoError(t, err, clues.ToCore(err))

	es := evts.GetValue()
	expect := make([]string, 0, len(es))

	for _, e := range es {
		expect = append(expect, api.EventCollisionKey(e))
	}

	results, err := suite.its.ac.Events().GetItemsInContainerByCollisionKey(ctx, suite.its.userID, "calendar")
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
