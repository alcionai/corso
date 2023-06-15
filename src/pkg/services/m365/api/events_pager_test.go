package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
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

	ctx, flush := tester.NewContext(t)
	defer flush()

	results, err := suite.cts.ac.Events().GetItemsInContainerByCollisionKey(ctx, suite.cts.userID, "calendar")
	require.NoError(t, err, clues.ToCore(err))
	require.Less(t, 0, len(results), "requires at least one result")

	for k, v := range results {
		assert.NotEmpty(t, k, "all keys should be populated")
		assert.NotEmpty(t, v, "all values should be populated")
	}
}
