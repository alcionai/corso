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

type EventsPagerIntgSuite struct {
	tester.Suite
	cts clientTesterSetup
}

func TestEventsPagerIntgSuite(t *testing.T) {
	suite.Run(t, &EventsPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *EventsPagerIntgSuite) SetupSuite() {
	suite.cts = newClientTesterSetup(suite.T())
}

func (suite *EventsPagerIntgSuite) TestGetItemsInContainerByCollisionKey() {
	t := suite.T()
	ac := suite.cts.ac.Events()

	ctx, flush := tester.NewContext(t)
	defer flush()

	container, err := ac.GetContainerByID(ctx, suite.cts.userID, "calendar")
	require.NoError(t, err, clues.ToCore(err))

	evts, err := ac.Stable.
		Client().
		Users().
		ByUserId(suite.cts.userID).
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

	results, err := ac.GetItemsInContainerByCollisionKey(ctx, suite.cts.userID, "calendar")
	require.NoError(t, err, clues.ToCore(err))
	require.Less(t, 0, len(results), "requires at least one result")

	for k, v := range results {
		assert.NotEmpty(t, k, "all keys should be populated")
		assert.NotEmpty(t, v, "all values should be populated")
	}
}
