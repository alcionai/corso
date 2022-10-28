package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

type CacheResolverSuite struct {
	suite.Suite
	gs graph.Service
}

func TestCacheResolverIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(CacheResolverSuite))
}

func (suite *CacheResolverSuite) SetupSuite() {
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

func (suite *CacheResolverSuite) TestPopulate() {
	ctx, flush := tester.NewContext()
	defer flush()

	ecc := eventCalendarCache{
		userID: tester.M365UserID(suite.T()),
		gs:     suite.gs,
	}

	cfc := contactFolderCache{
		userID: tester.M365UserID(suite.T()),
		gs:     suite.gs,
	}

	tests := []struct {
		name       string
		folderName string
		basePath   string
		resolver   graph.ContainerResolver
		canFind    assert.BoolAssertionFunc
	}{
		{
			name:       "Default Event Cache",
			folderName: DefaultCalendar,
			basePath:   DefaultCalendar,
			resolver:   &ecc,
			canFind:    assert.True,
		},
		{
			name:       "Default Event Folder Hidden",
			folderName: DefaultContactFolder,
			canFind:    assert.False,
			resolver:   &ecc,
		},
		{
			name:       "Name Not in Cache",
			folderName: "testFooBarWhoBar",
			canFind:    assert.False,
			resolver:   &ecc,
		},
		{
			name:       "Default Contact Cache",
			folderName: DefaultContactFolder,
			basePath:   DefaultContactFolder,
			canFind:    assert.True,
			resolver:   &cfc,
		},
		{
			name:       "Default Contact Hidden",
			folderName: DefaultContactFolder,
			canFind:    assert.False,
			resolver:   &cfc,
		},
		{
			name:       "Name Not in Cache",
			folderName: "testFooBarWhoBar",
			canFind:    assert.False,
			resolver:   &cfc,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			require.NoError(t, test.resolver.Populate(ctx, DefaultCalendar, test.basePath))
			_, isFound := test.resolver.PathInCache(test.folderName)
			test.canFind(t, isFound)
			assert.Greater(t, len(ecc.cache), 0)
		})
	}
}
