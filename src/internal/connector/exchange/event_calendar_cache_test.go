package exchange

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

type EventCalendarCacheSuite struct {
	suite.Suite
	gs graph.Service
}

func TestEventCalendarCacheIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(EventCalendarCacheSuite))
}

func (suite *EventCalendarCacheSuite) SetupSuite() {
	t := suite.T()

	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(t, err)

	a := tester.NewM365Account(t)
	require.NoError(t, err)

	m365, err := a.M365Config()
	require.NoError(t, err)

	service, err := createService(m365, false)
	require.NoError(t, err)

	suite.gs = service
}

func (suite *EventCalendarCacheSuite) TestPopulate() {
	ctx := context.Background()
	ecc := eventCalendarCache{
		userID: tester.M365UserID(suite.T()),
		gs:     suite.gs,
	}

	tests := []struct {
		name       string
		folderName string
		basePath   string
		canFind    assert.BoolAssertionFunc
	}{
		{
			name:       "Default Event Cache",
			folderName: DefaultCalendar,
			basePath:   DefaultCalendar,
			canFind:    assert.True,
		},
		{
			name:       "Default Event Folder Hidden",
			folderName: DefaultCalendar,
			canFind:    assert.False,
		},
		{
			name:       "Name Not in Cache",
			folderName: "testFooBarWhoBar",
			canFind:    assert.False,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			require.NoError(t, ecc.Populate(ctx, DefaultCalendar, test.basePath))
			_, isFound := ecc.PathInCache(test.folderName)
			test.canFind(t, isFound)
		})
	}
}
