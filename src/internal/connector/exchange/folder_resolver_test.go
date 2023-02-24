package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
)

type CacheResolverSuite struct {
	tester.Suite
	credentials account.M365Config
}

func TestCacheResolverIntegrationSuite(t *testing.T) {
	suite.Run(t, &CacheResolverSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *CacheResolverSuite) SetupSuite() {
	t := suite.T()

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.credentials = m365
}

func (suite *CacheResolverSuite) TestPopulate() {
	ctx, flush := tester.NewContext()
	defer flush()

	ac, err := api.NewClient(suite.credentials)
	require.NoError(suite.T(), err)

	cal, err := ac.Events().GetContainerByID(ctx, tester.M365UserID(suite.T()), DefaultCalendar)
	require.NoError(suite.T(), err)

	eventFunc := func(t *testing.T) graph.ContainerResolver {
		return &eventCalendarCache{
			userID: tester.M365UserID(t),
			enumer: ac.Events(),
			getter: ac.Events(),
		}
	}

	contactFunc := func(t *testing.T) graph.ContainerResolver {
		return &contactFolderCache{
			userID: tester.M365UserID(t),
			enumer: ac.Contacts(),
			getter: ac.Contacts(),
		}
	}

	tests := []struct {
		name, folderInCache, root, basePath string
		resolverFunc                        func(t *testing.T) graph.ContainerResolver
		canFind                             assert.BoolAssertionFunc
	}{
		{
			name:          "Default Event Cache",
			folderInCache: *cal.GetId(),
			root:          DefaultCalendar,
			basePath:      DefaultCalendar,
			resolverFunc:  eventFunc,
			canFind:       assert.True,
		},
		{
			name:          "Default Event Folder Hidden",
			folderInCache: DefaultContactFolder,
			root:          DefaultCalendar,
			canFind:       assert.False,
			resolverFunc:  eventFunc,
		},
		{
			name:          "Name Not in Cache",
			folderInCache: "testFooBarWhoBar",
			root:          DefaultCalendar,
			canFind:       assert.False,
			resolverFunc:  eventFunc,
		},
		{
			name:          "Default Contact Cache",
			folderInCache: DefaultContactFolder,
			root:          DefaultContactFolder,
			basePath:      DefaultContactFolder,
			canFind:       assert.True,
			resolverFunc:  contactFunc,
		},
		{
			name:          "Default Contact Hidden",
			folderInCache: DefaultContactFolder,
			root:          DefaultContactFolder,
			canFind:       assert.False,
			resolverFunc:  contactFunc,
		},
		{
			name:          "Name Not in Cache",
			folderInCache: "testFooBarWhoBar",
			root:          DefaultContactFolder,
			canFind:       assert.False,
			resolverFunc:  contactFunc,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			resolver := test.resolverFunc(t)
			require.NoError(t, resolver.Populate(ctx, fault.New(true), test.root, test.basePath))

			_, isFound := resolver.PathInCache(test.folderInCache)
			test.canFind(t, isFound)
		})
	}
}
