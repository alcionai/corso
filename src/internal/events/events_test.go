package events_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/tester"
)

type EventsIntegrationSuite struct {
	suite.Suite
}

func TestMetricsIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(tester.CorsoCITests, "floob"); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(EventsIntegrationSuite))
}

func (suite *EventsIntegrationSuite) TestNewBus() {
	t := suite.T()

	b := events.NewBus("s3", "bckt", "prfx", "tenid")
	require.NotEmpty(t, b)

	require.NoError(t, b.Close())
}
